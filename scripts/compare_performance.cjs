const fs = require('node:fs');
const path = require('node:path');
const os = require('node:os');
const crypto = require('node:crypto');
const { execFileSync } = require('node:child_process');
const { renderComparison } = require('./render_performance.cjs');

// The shared CRLF correction precedes the performance changes.
const PERFORMANCE_BASE = '8e3ac76991f5ffe810bd9a18597cdb017c909020';

const root = path.resolve(__dirname, '..');
const extraFixtures = [
  {
    path: 'pkg/gocui/text_area_navigation_test.go',
    symbols: ['BenchmarkTextAreaVerticalMovement', 'BenchmarkTextAreaVerticalKeypress', 'BenchmarkTextAreaLineStart'],
    imports: ['fmt', 'strings', 'testing'],
  },
  {
    path: 'pkg/gocui/search_stream_test.go',
    symbols: ['BenchmarkIncrementalSearchedDiff'],
    imports: ['strings', 'testing'],
  },
  {
    path: 'pkg/gui/filetree/selected_nodes_test.go',
    symbols: ['BenchmarkSmallTreeRangeSelection'],
    imports: ['fmt', 'testing', 'github.com/jesseduffield/lazygit/pkg/commands/models', 'github.com/jesseduffield/lazygit/pkg/common'],
  },
  {
    path: 'pkg/gui/filetree/flat_tree_parity_test.go',
    symbols: ['BenchmarkBuildFlatFileTrees', 'BenchmarkFlatTreeRefresh'],
    imports: ['fmt', 'testing', 'github.com/jesseduffield/lazygit/pkg/commands/models', 'github.com/jesseduffield/lazygit/pkg/common'],
  },
  {
    path: 'pkg/gui/presentation/graph/graph_layout_test.go',
    symbols: ['continuingGraphRow', 'BenchmarkGraphColumnAllocation'],
    imports: ['fmt', 'testing', 'github.com/jesseduffield/lazygit/pkg/commands/models', 'github.com/jesseduffield/lazygit/pkg/gui/style', 'github.com/jesseduffield/lazygit/pkg/utils'],
  },
  {
    path: 'pkg/gui/presentation/graph/graph_render_test.go',
    symbols: ['graphViewportFixture', 'BenchmarkGraphViewport'],
    imports: ['fmt', 'testing', 'github.com/gookit/color', 'github.com/jesseduffield/lazygit/pkg/commands/models', 'github.com/jesseduffield/lazygit/pkg/gui/style', 'github.com/jesseduffield/lazygit/pkg/utils', 'github.com/xo/terminfo'],
  },
];

function fixtureSource(file, text) {
  const fixture = extraFixtures.find(fixture => fixture.path === file);
  if (!fixture) return text;
  const declarations = fixture.symbols.map(symbol => {
    const declaration = text.match(new RegExp(`^func ${symbol}\\([\\s\\S]*?\\n}`, 'm'))?.[0];
    if (!declaration) throw new Error(`Missing fixture symbol: ${file}: ${symbol}`);
    return declaration;
  });
  return `package ${path.basename(path.dirname(file))}\n\nimport (\n${fixture.imports.map(name => `\t"${name}"`).join('\n')}\n)\n\n${declarations.join('\n\n')}\n`;
}

function benchmarkSource(file, text, rendererSource = '') {
  // Seeding this existing-text fixture is outside the timed edit pair.
  // Direct initialization avoids benchmarking TypeString during setup.
  if (file === 'pkg/gocui/edit_bench_test.go') {
    const seed = 'v.TextArea.TypeString(strings.Repeat("a", size))';
    if (!text.includes(seed)) throw new Error('The existing-text benchmark setup changed');
    text = text.replace(seed, 'v.TextArea.content = strings.Repeat("a", size)\n\t\t\tv.TextArea.cursor = len(v.TextArea.content)\n\t\t\tv.TextArea.updateCells()');
  }
  if (file === 'pkg/gocui/text_area_navigation_test.go') {
    const seed = 'v.TextArea.TypeString(strings.Repeat("some text for a line\\n", count))';
    if (!text.includes(seed)) throw new Error('The editor-navigation setup changed');
    text = text.replace(seed, 'v.TextArea.content = strings.Repeat("some text for a line\\n", count)\n v.TextArea.updateCells()');
  }
  for (const [target, variable] of [
    ['pkg/gui/context/filtered_list_bench_test.go', 'list'],
    ['pkg/gui/filetree/range_bench_test.go', 'view'],
  ]) {
    if (file !== target) continue;
    const step = `for b.Loop() {\n\t\t\t\t${variable}.MoveSelectedLine(1)`;
    const resetStep = `for b.Loop() {\n\t\t\t\t${variable}.SetSelectedLineIdx(count / 2)\n\t\t\t\t${variable}.MoveSelectedLine(1)`;
    if (text.includes(resetStep)) continue;
    if (!text.includes(step)) throw new Error(`The navigation benchmark changed: ${file}`);
    text = text.replace(step, resetStep);
  }
  if (file === 'pkg/gui/filetree/build_tree_bench_test.go') {
    const order = /cmp := (?:NewNodeSortOrder|NodeSortComparator\[models\.CommitFile\])\("mixed", false\)/;
    if (!order.test(text)) throw new Error('The tree-construction fixture changed');
    text = text.replace('"testing"', '"testing"\n\t"github.com/jesseduffield/lazygit/pkg/common"');
    text = text.replace(order, `c := common.NewDummyCommon()
      c.UserConfig().Gui.ShowRootItemInFileTree = true
      c.UserConfig().Gui.FileTreeSortOrder = "mixed"
      c.UserConfig().Gui.FileTreeSortCaseSensitive = false
      tree := NewCommitFileTree(func() []*models.CommitFile { return files }, c, true)`);
    text = text.replace('BuildTreeFromCommitFiles(files, true, cmp)', 'tree.SetTree()');
  }
  if (file === 'pkg/gui/filetree/flat_tree_parity_test.go') {
    const orders = /fileOrder := NewNodeSortOrder\("mixed", false\)\s+commitOrder := NewNodeSortOrder\("mixed", false\)/;
    if (!orders.test(text)) throw new Error('The flat-construction fixture changed');
    text = text.replace(orders, `c := common.NewDummyCommon()
      c.UserConfig().Gui.ShowRootItemInFileTree = true
      c.UserConfig().Gui.FileTreeSortOrder = "mixed"
      c.UserConfig().Gui.FileTreeSortCaseSensitive = false
      working := NewFileTree(func() []*models.File { return files }, c, false)
      committed := NewCommitFileTree(func() []*models.CommitFile { return commits }, c, false)`);
    text = text.replace('BuildFlatTreeFromCommitFiles(commits, true, commitOrder)', 'committed.SetTree()');
    text = text.replace('BuildFlatTreeFromFiles(files, true, fileOrder)', 'working.SetTree()');
  }
  if (file === 'pkg/gui/filetree/selected_nodes_test.go') {
    text = text.replace('c := common.NewDummyCommon()', 'c := common.NewDummyCommon()\n c.UserConfig().Gui.ShowRootItemInFileTree = false');
    for (const [variable, tree] of [['working', 'IFileTree'], ['committed', 'ICommitFileTree']]) {
      const cast = tree === 'IFileTree' ? 'FileTree' : 'CommitFileTree';
      const invalidation = `${variable}.${tree}.(*${cast}).indexedTree = nil`;
      if (!text.includes(invalidation)) throw new Error('The cold range fixture changed');
      text = text.replace(invalidation, `${variable}.ToggleCollapsed("dir")\n ${variable}.ToggleCollapsed("dir")`);
    }
  }
  if (file === 'pkg/gocui/search_stream_test.go') {
    if (!text.includes('v.RefreshSearch()')) throw new Error('The streaming-search fixture changed');
    text = text.replaceAll('v.RefreshSearch()', 'v.GetSearchStatus()');
  }
  if (file === 'pkg/gui/presentation/commits_bench_test.go' && rendererSource.includes('mergedIntoLocalMainBranch *set.Set[string]')) {
    const parameters = rendererSource.match(/func GetCommitListDisplayStrings\(([\s\S]*?)\) \[\]\[\]string/)?.[1];
    const loop = /for b\.Loop\(\) \{\s*GetCommitListDisplayStrings\(([^\n]+)\)\s*\}/;
    const argumentsText = text.match(loop)?.[1];
    const setup = 'mergedLocally := branchstack.CommitsMergedIntoLocalMainBranch(commits, branches, c.UserConfig().Git.MainBranches)';
    if (!parameters || !argumentsText?.includes('branches, mergedLocally,') || !text.includes(setup)) throw new Error('The commit-viewport fixture changed');
    // This fixture's unmerged main tip has only a merged parent. Prepare its
    // known ancestry outside timing, without importing a newer package.
    text = text.replace(/\s*"github.com\/jesseduffield\/lazygit\/pkg\/gui\/branchstack"/, '');
    text = text.replace(setup, 'mergedLocally := set.New[string]()\n if main == "unmerged" { mergedLocally.Add(tip) }');
    const legacyParameters = parameters.split('\n').filter(line => !line.includes('mergedIntoLocalMainBranch ')).join('\n');
    const legacyArguments = argumentsText.replace('branches, mergedLocally,', 'branches,');
    // Select the typed call before timing. Both revisions run the same fixture
    // and call a function value once per operation.
    text = text.replace(loop, `switch render := any(GetCommitListDisplayStrings).(type) {
      case func(${parameters}) [][]string:
        for b.Loop() { render(${argumentsText}) }
      case func(${legacyParameters}) [][]string:
        for b.Loop() { render(${legacyArguments}) }
      default:
        b.Fatal("Unsupported commit renderer interface")
      }`);
  }
  return text;
}

function benchmarkCheckpoint(directory, pkg, args, settings) {
  const digest = text => crypto.createHash('sha256').update(text).digest('hex');
  const packageDirectory = path.join(directory, pkg);
  const sources = fs.readdirSync(packageDirectory).filter(file => file.endsWith('.go')).sort().map(file => ({
    path: file, sha256: digest(fs.readFileSync(path.join(packageDirectory, file))),
  }));
  const inputs = {
    revision: execFileSync('git', ['rev-parse', 'HEAD'], { cwd: directory, encoding: 'utf8' }).trim(),
    pkg, args, settings, sources,
    goVersion: execFileSync('go', ['version'], { cwd: directory, encoding: 'utf8' }).trim(),
    cpu: os.cpus()[0].model,
  };
  return { key: digest(JSON.stringify(inputs)), inputs };
}

function measure() {
  const git = (...args) => execFileSync('git', args, { cwd: root, encoding: 'utf8' }).trim();
  const options = process.argv.slice(2);
  for (const option of options.filter(arg => arg.startsWith('--'))) {
    if (!['--prepare-only', '--resume'].includes(option) && !option.startsWith('--package=')) throw new Error(`Unknown comparison option: ${option}`);
  }
  const prepareOnly = options.includes('--prepare-only');
  const resume = options.includes('--resume');
  const onlyPackage = options.find(arg => arg.startsWith('--package='))?.slice('--package='.length);
  if (onlyPackage === '') throw new Error('A benchmark package is required');
  const args = options.filter(arg => !arg.startsWith('--'));
  const refs = { before: args[0] || PERFORMANCE_BASE, after: args[1] || 'HEAD' };
  const revisions = Object.fromEntries(Object.entries(refs).map(([side, ref]) => [side, git('rev-parse', `${ref}^{commit}`)]));
  const workspace = path.join(git('rev-parse', '--absolute-git-dir'), 'perf-comparison');
  const settings = { samples: 5, benchtime: '200ms', GOMAXPROCS: '1', GOGC: '100', GOMEMLIMIT: 'off' };
  const env = { ...process.env, GOMAXPROCS: settings.GOMAXPROCS, GOGC: settings.GOGC, GOMEMLIMIT: settings.GOMEMLIMIT, GOFLAGS: '-mod=vendor', GOTOOLCHAIN: 'local', TERM: 'xterm-256color', FORCE_COLOR: '1' };
  delete env.NO_COLOR;
  const go = (cwd, ...args) => execFileSync('go', args, { cwd, env, encoding: 'utf8', maxBuffer: 16 * 1024 * 1024, stdio: ['ignore', 'pipe', 'pipe'] });
  const sha256 = value => crypto.createHash('sha256').update(value).digest('hex');
  const source = (revision, file) => execFileSync('git', ['show', `${revision}:${file}`], { cwd: root, encoding: 'utf8' });
  const treeFiles = git('ls-tree', '-r', '--name-only', revisions.after).split('\n');
  const files = treeFiles.filter(file => /bench.*_test\.go$/.test(file) || extraFixtures.some(fixture => fixture.path === file));
  const harness = [];
  const families = new Set();
  fs.mkdirSync(workspace, { recursive: true });
  const checkouts = {};

  for (const [side, revision] of Object.entries(revisions)) {
    const directory = path.join(workspace, `${side === 'before' ? 'upstream' : 'head'}-${revision.slice(0, 12)}`);
    if (!fs.existsSync(directory)) git('worktree', 'add', '--detach', directory, revision);
    const actual = execFileSync('git', ['rev-parse', 'HEAD'], { cwd: directory, encoding: 'utf8' }).trim();
    if (actual !== revision) throw new Error(`Unexpected revision in ${directory}`);
    const changed = execFileSync('git', ['diff', '--name-only', 'HEAD'], { cwd: directory, encoding: 'utf8' }).trim().split('\n').filter(Boolean);
    const untracked = execFileSync('git', ['ls-files', '--others', '--exclude-standard'], { cwd: directory, encoding: 'utf8' }).trim().split('\n').filter(Boolean);
    if ([...changed, ...untracked].some(file => !files.includes(file))) throw new Error(`Unexpected changes in ${directory}`);
    checkouts[side] = directory;
  }

  for (const file of files) {
    const rendererSource = file === 'pkg/gui/presentation/commits_bench_test.go' ? source(revisions.after, 'pkg/gui/presentation/commits.go') : '';
    const text = benchmarkSource(file, fixtureSource(file, source(revisions.after, file)), rendererSource);
    for (const match of text.matchAll(/^func (Benchmark\w+)\(/gm)) families.add(match[1]);
    const digest = sha256(text);
    for (const directory of Object.values(checkouts)) {
      const target = path.join(directory, file);
      fs.mkdirSync(path.dirname(target), { recursive: true });
      fs.writeFileSync(target, text);
      if (sha256(fs.readFileSync(target)) !== digest) throw new Error(`Harness mismatch: ${file}`);
    }
    const fixture = extraFixtures.find(fixture => fixture.path === file);
    harness.push({ path: file, sha256: digest, ...(fixture ? { symbols: fixture.symbols } : {}) });
  }

  // These existing cases also have identical fixtures in both revisions.
  for (const [file, name, helpers] of [
    ['pkg/gocui/search_test.go', 'BenchmarkWriteToSearchedView', []],
    ['pkg/gui/presentation/graph/graph_test.go', 'BenchmarkRenderCommitGraph', ['generateCommits']],
    ['pkg/gocui/view_test.go', null, ['osc1717']],
  ]) {
    for (const symbol of [name, ...helpers].filter(Boolean)) {
      const expression = new RegExp(`func ${symbol}\\([\\s\\S]*?\\n}`);
      const before = source(revisions.before, file).match(expression)?.[0];
      const after = source(revisions.after, file).match(expression)?.[0];
      if (!before || before !== after) throw new Error(`Benchmark fixture differs: ${symbol}`);
      harness.push({ path: file, symbol, sha256: sha256(after) });
    }
    if (name) families.add(name);
  }

  const descriptions = {
    BenchmarkViewDraw: p => `Redraw a ${p[0]} view with ${p[1]} content.`,
    BenchmarkGuiFlush: p => `Flush two panes through a mock terminal at ${p[0]}, changing selection and scroll position.`,
    BenchmarkDrawSearchedView: p => `Redraw an 80 × 24 searched viewport within ${value(p, 'lines')} lines.`,
    BenchmarkViewWrite: p => `Write 2,000 ${p[0]} lines into a new view.`,
    BenchmarkViewWriteLongLine: () => 'Write one 13,001-byte line into a new view.',
    BenchmarkViewWriteGitColoredDiff: () => 'Stream 50,000 mixed Unicode lines with separately colored Git prefixes into a new view.',
    BenchmarkEscapeSpecialChars: p => `Escape control characters in a ${p[0] === 'ordinary' ? 'plain filename' : 'filename containing tabs and newlines'}.`,
    BenchmarkWriteToSearchedView: () => 'Write 2,000 lines into a searched view and read search status.',
    BenchmarkEditorInput: p => `Type ${Math.floor(value(p, 'bytes') / 10) * 10} bytes with hard newlines, updating layout and view per key.`,
    BenchmarkEditorParagraphInput: p => `Type ${Math.floor(value(p, 'bytes') / 10) * 10} bytes of a spaced paragraph, wrapping at 72 columns.`,
    BenchmarkEditorUnbrokenInput: p => `Type ${value(p, 'bytes')} unbroken ASCII characters, updating layout and view per key.`,
    BenchmarkEditorEditUnbrokenText: p => `Insert and backspace one character in ${value(p, 'bytes')} bytes of existing unbroken text; update layout and view after each key.`,
    BenchmarkCommitDescriptionInput: p => `Type ${Math.floor(value(p, 'bytes') / 10) * 10} bytes through the commit-description editor.`,
    BenchmarkCommitViewport: p => `Render 40 commit rows from ${value(p, 'commits')} loaded commits without a graph; local main tip ${value(p, 'main')}, other branch marker ${value(p, 'markers') === 'true' ? 'visible' : 'absent'}.`,
    BenchmarkPasteText: p => `Insert ${Math.floor(value(p, 'bytes') / 17) * 17} bytes into prefilled text with TypeString, including layout.`,
    BenchmarkFilteredListNavigation: p => `Reset selection to the midpoint, move one row, retrieve it, and count ${value(p, 'items')} filtered branches.`,
    BenchmarkBuildCommitFileTree: p => `Build and sort ${value(p, 'files')} reverse-ordered commit files through the file-tree construction entry point.`,
    BenchmarkBuildFlatFileTrees: p => `Build a flat ${value(p, 'commits') === 'true' ? 'commit-file' : 'working-file'} list of ${value(p, 'files')} files with ${value(p, 'deep') === 'true' ? 'a distinct nested directory per file' : 'one shared directory'}; construction entry point included.`,
    BenchmarkFlatTreeRefresh: p => `Refresh a flat ${value(p, 'commits') === 'true' ? 'commit-file' : 'working-file'} view model with ${value(p, 'files')} files, including selection recovery.`,
    BenchmarkSmallTreeRangeSelection: p => `Select one row near the end of a ${value(p, 'files')}-file ${value(p, 'commits') === 'true' ? 'commit-file' : 'working-file'} tree; ${value(p, 'cold') === 'true' ? 'toggle a directory twice before each selection to invalidate visibility' : 'visibility unchanged'}.`,
    BenchmarkTextAreaVerticalMovement: p => `Move Up and Down in ${value(p, 'lines')} prepared editor lines, ${value(p, 'bottom') === 'true' ? 'near the end' : 'near the beginning'}; rendering excluded.`,
    BenchmarkTextAreaVerticalKeypress: p => `Handle Up and Down through the default editor near the end of ${value(p, 'lines')} lines, including rendering.`,
    BenchmarkTextAreaLineStart: p => `Move to the start of a line near the end of ${value(p, 'lines')} prepared editor lines.`,
    BenchmarkIncrementalSearchedDiff: p => `Stream 8,192 searched lines in 64 batches, reading search status after each batch; ${p[0]} view.`,
    BenchmarkGraphColumnAllocation: p => `Lay out a graph row with ${value(p, 'columns')} incoming columns and its current commit in the final column.`,
    BenchmarkGraphViewport: p => `Render ${value(p, 'rows')} cached ${value(p, 'branches') === 'true' ? 'branching' : 'linear'} graph rows with GOMAXPROCS=${value(p, 'cpu') || '1'}; layout excluded.`,
    BenchmarkFileTreeRangeSelection: p => `Retrieve a range spanning a ${value(p, 'files')}-file tree.`,
    BenchmarkFileTreeNavigation: p => `Reset selection to the midpoint, move one row, retrieve it, and count rows in a ${value(p, 'files')}-file tree.`,
    BenchmarkRestoreFileSelection: p => `Restore selection among ${value(p, 'files')} files; selected file ${value(p, 'removed') === 'true' ? 'removed' : 'retained'}.`,
    BenchmarkRestoreAfterSingleRemoval: p => `Restore selection after removing one file from ${value(p, 'files')} remaining files.`,
    BenchmarkRefreshAfterDirectoryRemoval: p => `Refresh a tree after removing a ${value(p, 'files')}-file directory; fixture setup excluded.`,
    BenchmarkSearchRepeatedPrefix: p => `Search 2,000 lines of 200 a characters for ${value(p, 'query') === '1' ? 'a' : `a ${value(p, 'query')}-character query ending in b`}; rendered layout warmed.`,
    BenchmarkSearchOrdinaryText: p => `Search 2,000 ordinary diff lines for "${p[0]}"; rendered layout warmed.`,
    BenchmarkSearchLargeDiff: p => `Search ${value(p, 'lines')} diff lines with wide and combining characters for "${value(p, 'query')}"; rendered layout and initial query warmed.`,
    BenchmarkViewWriteLargeDiff: p => `Stream 50,000 ${p[0]} diff lines into a new view, including parsing and allocation.`,
    BenchmarkSearchQueryLifetime: p => `Alternate aa, aaa, aaaa, and a queries across 2,000 lines of 200 a characters; ${value(p, 'clear') === 'true' ? 'clear search after each query' : 'keep search active'}; rendered layout and initial a query warmed.`,
    BenchmarkNearestSearchPosition: p => `Find the nearest result near the end of ${value(p, 'matches')} search matches.`,
    BenchmarkScrollSearchPosition: p => `Select a visible result after a distant scroll ${value(p, 'up') === 'true' ? 'up' : 'down'} across ${value(p, 'matches')} matches; reset each iteration.`,
    BenchmarkRenderCommitGraph: () => 'Render a graph of 50 commits generated with seed 1234.',
    BenchmarkRenderCommitGraphVariants: p => `Render ${value(p, 'commits')} commits generated with seed 1234, using ${value(p, 'symbols')} symbols; ${value(p, 'selected') === 'true' ? 'select the middle commit' : 'no selected commit'}.`,
  };
  function value(parts, name) { return parts.find(part => part.startsWith(`${name}=`))?.slice(name.length + 1); }
  for (const name of families) if (!descriptions[name]) throw new Error(`Missing description: ${name}`);
  let packages = [...new Set([...files.map(file => `./${path.dirname(file)}`), './pkg/gui/presentation/graph'])];
  if (onlyPackage) {
    if (!packages.includes(onlyPackage)) throw new Error(`Unknown benchmark package: ${onlyPackage}`);
    packages = [onlyPackage];
  }
  const pattern = `^(${[...families].join('|')})$`;
  const samples = { before: new Map(), after: new Map() };
  const checkpoints = [];

  function runBenchmark(pkg, side, args, filename) {
    const file = path.join(workspace, filename);
    const metadataFile = file + '.json';
    const checkpoint = benchmarkCheckpoint(checkouts[side], pkg, args, settings);
    let output;
    let metadata;
    if (resume && fs.existsSync(file) && fs.existsSync(metadataFile)) {
      metadata = JSON.parse(fs.readFileSync(metadataFile, 'utf8'));
      const saved = fs.readFileSync(file, 'utf8');
      if (metadata.key === checkpoint.key && metadata.outputHash === sha256(saved)) {
        output = saved;
        console.log(`${side}: ${pkg} reused ${filename}`);
      }
    }
    if (output === undefined) {
      output = go(checkouts[side], ...args);
      metadata = { ...checkpoint, outputHash: sha256(output) };
      fs.writeFileSync(file, output);
      fs.writeFileSync(metadataFile, JSON.stringify(metadata, null, 2) + '\n');
    }
    checkpoints.push({ side, pkg, file: filename, key: metadata.key, outputHash: metadata.outputHash, recoveredFromInterruptedRun: metadata.recoveredFromInterruptedRun || undefined });
    return output;
  }

  function recordSamples(output, pkg, side) {
    for (const line of output.split('\n')) {
      const match = line.match(/^(Benchmark\S+?)(?:-(\d+))?\s+\d+\s+([\d.]+) ns\/op/);
      if (!match) continue;
      const cpu = Number(match[2] || 1);
      const id = `${pkg.slice(6)}/${match[1]}${cpu === 1 ? '' : `/cpu=${cpu}`}`;
      const row = { ns: Number(match[3]), bytes: Number(line.match(/(\d+) B\/op/)?.[1]), allocations: Number(line.match(/(\d+) allocs\/op/)?.[1]) };
      if (!Number.isFinite(row.bytes) || !Number.isFinite(row.allocations)) throw new Error(`Incomplete sample: ${id}`);
      if (!samples[side].has(id)) samples[side].set(id, []);
      samples[side].get(id).push(row);
    }
  }

  for (const pkg of packages) {
    for (const side of ['before', 'after']) {
      console.log(`${side}: ${pkg}`);
      if (prepareOnly) {
        go(checkouts[side], 'test', pkg, '-run=^$', '-bench=^$');
        console.log(`${side}: ${pkg} fixtures compile`);
        continue;
      }
      const filename = `${side}-${pkg.slice(2).replaceAll('/', '-')}.txt`;
      const output = runBenchmark(pkg, side, ['test', pkg, '-run=^$', `-bench=${pattern}`, '-benchmem', `-benchtime=${settings.benchtime}`, `-count=${settings.samples}`, '-cpu=1'], filename);
      recordSamples(output, pkg, side);
      console.log(`${side}: ${pkg} complete`);
      if (pkg === './pkg/gui/presentation/graph' && families.has('BenchmarkGraphViewport')) {
        const parallelOutput = runBenchmark(pkg, side, ['test', pkg, '-run=^$', '-bench=^BenchmarkGraphViewport$', '-benchmem', `-benchtime=${settings.benchtime}`, `-count=${settings.samples}`, '-cpu=4,16'], `${side}-graph-viewport-parallel.txt`);
        recordSamples(parallelOutput, pkg, side);
      }
    }
  }
  if (prepareOnly || onlyPackage) return null;

  const ids = [...samples.before.keys()].sort((a, b) => a.localeCompare(b, 'en', { numeric: true }));
  if (ids.length !== samples.after.size || [...samples.after.keys()].some(id => !samples.before.has(id))) throw new Error('Benchmark case sets differ');
  for (const family of families) {
    if (!ids.some(id => id.endsWith('/' + family) || id.includes('/' + family + '/'))) throw new Error('Missing benchmark family: ' + family);
  }
  const median = numbers => [...numbers].sort((a, b) => a - b)[Math.floor(numbers.length / 2)];
  const rows = ids.map(id => {
    const before = samples.before.get(id), after = samples.after.get(id);
    if (before.length !== settings.samples || after.length !== settings.samples) throw new Error(`Incomplete repetitions: ${id}`);
    const benchmark = id.slice(id.indexOf('/Benchmark') + 1).split('/');
    const describe = descriptions[benchmark.shift()];
    if (!describe) throw new Error(`Missing description: ${id}`);
    return { id, description: describe(benchmark), before: median(before.map(row => row.ns)), after: median(after.map(row => row.ns)), samples: { before, after } };
  });
  return { refs, revisions, settings, graphViewportCPUs: families.has('BenchmarkGraphViewport') ? [1, 4, 16] : undefined, goVersion: go(root, 'version').trim(), platform: `${os.platform()}/${os.arch()}`, cpu: os.cpus()[0].model, harness, checkpoints, rows };
}

if (require.main === module) {
  const dataPath = path.join(root, 'docs-master/dev/performance-comparison.json');
  const data = process.argv.slice(2).includes('--render-only')
    ? JSON.parse(fs.readFileSync(dataPath, 'utf8'))
    : measure();
  if (data === null) process.exit(0);
  const { revisions, settings } = data;
  const imageDirectory = 'docs-master/dev/performance-comparisons';
  const { table, images } = renderComparison(data, imageDirectory);
  const summary = [
    '## Performance',
    '',
    `Before (baseline): \`${revisions.before.slice(0, 12)}\` · After (measured revision): \`${revisions.after.slice(0, 12)}\`.`,
    '',
  ];
  if (data.equivalentAfter) {
    summary.push(`Equivalent branch revision: \`${data.equivalentAfter.revision.slice(0, 12)}\`. Its Git tree is identical to the measured after revision.`, '');
  }
  summary.push(
    `Same benchmark sources and settings. Median of ${settings.samples} samples, ${settings.benchtime} per sample, GOMAXPROCS=${settings.GOMAXPROCS}, GOGC=${settings.GOGC}, GOMEMLIMIT=${settings.GOMEMLIMIT}.`,
    '',
    ...(data.graphViewportCPUs ? [`Cached graph viewports also run with GOMAXPROCS=${data.graphViewportCPUs.filter(cpu => cpu !== Number(settings.GOMAXPROCS)).join(' and ')}; those rows state their CPU setting.`, ''] : []),
    `\`${data.goVersion}\` · ${data.cpu}.`,
    '',
    'Each pair of bars uses its own scale. Negative change means faster; positive means slower.',
    '',
    '[Raw samples and fixture metadata](docs-master/dev/performance-comparison.json) · [Investigation notes](docs-master/dev/Performance-follow-up.md) · [Independent follow-up results](docs-master/dev/Independent-performance-follow-ups.md).',
    '',
    table,
  );

  const readmePath = path.join(root, 'README.md');
  const readme = fs.readFileSync(readmePath, 'utf8');
  const startMarker = '<!-- performance-summary:start -->';
  const endMarker = '<!-- performance-summary:end -->';
  const start = readme.indexOf(startMarker);
  const end = readme.indexOf(endMarker, start + startMarker.length);
  if (start < 0 || end < 0) throw new Error('README.md is missing performance summary placeholders');
  const updated = readme.slice(0, start + startMarker.length) + '\n\n' + summary.join('\n') + '\n\n' + readme.slice(end);

  fs.writeFileSync(dataPath, JSON.stringify(data, null, 2) + '\n');
  const directory = path.join(root, imageDirectory);
  fs.mkdirSync(directory, { recursive: true });
  const generated = new Set(images.map(image => image.fileName));
  for (const image of images) fs.writeFileSync(path.join(directory, image.fileName), image.svg);
  for (const fileName of fs.readdirSync(directory)) {
    if (/^[a-zA-Z0-9_-]+-[0-9a-f]{12}\.svg$/.test(fileName) && !generated.has(fileName)) {
      fs.unlinkSync(path.join(directory, fileName));
    }
  }
  fs.writeFileSync(readmePath, updated);
  console.log(`Wrote ${images.length} timing SVGs to ${imageDirectory} and updated README.md`);
}

module.exports = { benchmarkSource, fixtureSource, benchmarkCheckpoint };
