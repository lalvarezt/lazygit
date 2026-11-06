const fs = require('node:fs');
const path = require('node:path');
const os = require('node:os');
const crypto = require('node:crypto');
const { execFileSync } = require('node:child_process');
const { renderComparison } = require('./render_performance.cjs');

// The shared CRLF correction precedes the performance changes.
const PERFORMANCE_BASE = '8e3ac76991f5ffe810bd9a18597cdb017c909020';

const root = path.resolve(__dirname, '..');
function benchmarkSource(file, text) {
  // Seeding this existing-text fixture is outside the timed edit pair.
  // Direct initialization avoids benchmarking TypeString during setup.
  if (file === 'pkg/gocui/edit_bench_test.go') {
    const seed = 'v.TextArea.TypeString(strings.Repeat("a", size))';
    if (!text.includes(seed)) throw new Error('The existing-text benchmark setup changed');
    text = text.replace(seed, 'v.TextArea.content = strings.Repeat("a", size)\n\t\t\tv.TextArea.cursor = len(v.TextArea.content)\n\t\t\tv.TextArea.updateCells()');
  }
  for (const [target, variable] of [
    ['pkg/gui/context/filtered_list_bench_test.go', 'list'],
    ['pkg/gui/filetree/range_bench_test.go', 'view'],
  ]) {
    if (file !== target) continue;
    const step = `for b.Loop() {\n\t\t\t\t${variable}.MoveSelectedLine(1)`;
    if (!text.includes(step)) throw new Error(`The navigation benchmark changed: ${file}`);
    text = text.replace(step, `for b.Loop() {\n\t\t\t\t${variable}.SetSelectedLineIdx(count / 2)\n\t\t\t\t${variable}.MoveSelectedLine(1)`);
  }
  return text;
}

function measure() {
  const git = (...args) => execFileSync('git', args, { cwd: root, encoding: 'utf8' }).trim();
  const refs = { before: process.argv[2] || PERFORMANCE_BASE, after: process.argv[3] || 'HEAD' };
  const revisions = Object.fromEntries(Object.entries(refs).map(([side, ref]) => [side, git('rev-parse', `${ref}^{commit}`)]));
  const workspace = path.join(git('rev-parse', '--absolute-git-dir'), 'perf-comparison');
  const settings = { samples: 5, benchtime: '200ms', GOMAXPROCS: '1', GOGC: '100', GOMEMLIMIT: 'off' };
  const env = { ...process.env, GOMAXPROCS: settings.GOMAXPROCS, GOGC: settings.GOGC, GOMEMLIMIT: settings.GOMEMLIMIT, GOFLAGS: '-mod=vendor', GOTOOLCHAIN: 'local', TERM: 'xterm-256color', FORCE_COLOR: '1' };
  delete env.NO_COLOR;
  const go = (cwd, ...args) => execFileSync('go', args, { cwd, env, encoding: 'utf8', maxBuffer: 16 * 1024 * 1024, stdio: ['ignore', 'pipe', 'pipe'] });
  const sha256 = value => crypto.createHash('sha256').update(value).digest('hex');
  const source = (revision, file) => execFileSync('git', ['show', `${revision}:${file}`], { cwd: root, encoding: 'utf8' });
  const files = git('ls-tree', '-r', '--name-only', revisions.after).split('\n').filter(file => /bench.*_test\.go$/.test(file));
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
    const text = benchmarkSource(file, source(revisions.after, file));
    for (const match of text.matchAll(/^func (Benchmark\w+)\(/gm)) families.add(match[1]);
    const digest = sha256(text);
    for (const directory of Object.values(checkouts)) {
      const target = path.join(directory, file);
      fs.mkdirSync(path.dirname(target), { recursive: true });
      fs.writeFileSync(target, text);
      if (sha256(fs.readFileSync(target)) !== digest) throw new Error(`Harness mismatch: ${file}`);
    }
    harness.push({ path: file, sha256: digest });
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
    BenchmarkPasteText: p => `Insert ${Math.floor(value(p, 'bytes') / 17) * 17} bytes into prefilled text with TypeString, including layout.`,
    BenchmarkFilteredListNavigation: p => `Reset selection to the midpoint, move one row, retrieve it, and count ${value(p, 'items')} filtered branches.`,
    BenchmarkBuildCommitFileTree: p => `Build and sort a tree of ${value(p, 'files')} commit files supplied in reverse order.`,
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
  const packages = [...new Set([...files.map(file => `./${path.dirname(file)}`), './pkg/gui/presentation/graph'])];
  const pattern = `^(${[...families].join('|')})$`;
  const samples = { before: new Map(), after: new Map() };

  for (const pkg of packages) {
    for (const side of ['before', 'after']) {
      console.log(`${side}: ${pkg}`);
      const output = go(checkouts[side], 'test', pkg, '-run=^$', `-bench=${pattern}`, '-benchmem', `-benchtime=${settings.benchtime}`, `-count=${settings.samples}`, '-cpu=1');
      const filename = `${side}-${pkg.slice(2).replaceAll('/', '-')}.txt`;
      fs.writeFileSync(path.join(workspace, filename), output);
      for (const line of output.split('\n')) {
        const match = line.match(/^(Benchmark\S+?)(?:-\d+)?\s+\d+\s+([\d.]+) ns\/op/);
        if (!match) continue;
        const id = `${pkg.slice(6)}/${match[1]}`;
        const row = { ns: Number(match[2]), bytes: Number(line.match(/(\d+) B\/op/)?.[1]), allocations: Number(line.match(/(\d+) allocs\/op/)?.[1]) };
        if (!Number.isFinite(row.bytes) || !Number.isFinite(row.allocations)) throw new Error(`Incomplete sample: ${id}`);
        if (!samples[side].has(id)) samples[side].set(id, []);
        samples[side].get(id).push(row);
      }
      console.log(`${side}: ${pkg} complete`);
    }
  }

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
  return { refs, revisions, settings, goVersion: go(root, 'version').trim(), platform: `${os.platform()}/${os.arch()}`, cpu: os.cpus()[0].model, harness, rows };
}

if (require.main === module) {
  const dataPath = path.join(root, 'docs-master/dev/performance-comparison.json');
  const data = process.argv[2] === '--render-only'
    ? JSON.parse(fs.readFileSync(dataPath, 'utf8'))
    : measure();
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
    `\`${data.goVersion}\` · ${data.cpu}.`,
    '',
    'Each pair of bars uses its own scale. Negative change means faster; positive means slower.',
    '',
    '[Raw samples and fixture metadata](docs-master/dev/performance-comparison.json) · [Investigation notes](docs-master/dev/Performance-follow-up.md).',
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

module.exports = { benchmarkSource };
