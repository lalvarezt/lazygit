## Personal fork disclaimer

> [!WARNING]
> **This is a personal project, provided as-is.**
>
> I maintain this fork for my own use. Expect rough edges and changes driven by
> my needs, with no guarantees of stability, maintenance, or support.
>
> Most of the changes in this fork were developed with AI assistance.

This fork builds on [Lazygit](https://github.com/jesseduffield/lazygit) by
[Jesse Duffield](https://github.com/jesseduffield) and the
[Lazygit contributors](https://github.com/jesseduffield/lazygit/graphs/contributors).

<!-- performance-summary:start -->

## Performance

Before (baseline): `8e3ac76991f5` · After (measured revision): `b920143e765f`.

Same benchmark sources and settings. Median of 5 samples, 200ms per sample, GOMAXPROCS=1, GOGC=100, GOMEMLIMIT=off.

Cached graph viewports also run with GOMAXPROCS=4 and 16; those rows state their CPU setting.

`go version go1.27.3 linux/amd64` · 13th Gen Intel(R) Core(TM) i7-1370P.

Each pair of bars uses its own scale. Negative change means faster; positive means slower.

[Raw samples and fixture metadata](docs-master/dev/performance-comparison.json) · [Investigation notes](docs-master/dev/Performance-follow-up.md) · [Independent follow-up results](docs-master/dev/Independent-performance-follow-ups.md).

<table>
<thead><tr><th align="left">Benchmark / workload</th><th align="left">Before / after and change</th></tr></thead>
<tbody>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkDrawSearchedView/lines=100">Draw searched view</strong><br>
Redraw an 80 × 24 searched viewport within 100 lines.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkDrawSearchedView-lines-100-b3cacbadd434.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkDrawSearchedView/lines=100: 655.3 µs before, 53.9 µs after, -91.77% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkDrawSearchedView/lines=1000">Draw searched view</strong><br>
Redraw an 80 × 24 searched viewport within 1000 lines.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkDrawSearchedView-lines-1000-d66a7e5b1fe8.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkDrawSearchedView/lines=1000: 5.36 ms before, 55.25 µs after, -98.97% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkDrawSearchedView/lines=10000">Draw searched view</strong><br>
Redraw an 80 × 24 searched viewport within 10000 lines.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkDrawSearchedView-lines-10000-eafc1afee103.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkDrawSearchedView/lines=10000: 51.43 ms before, 56.82 µs after, -99.89% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorEditUnbrokenText/bytes=4096">Editor edit unbroken text</strong><br>
Insert and backspace one character in 4096 bytes of existing unbroken text; update layout and view after each key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorEditUnbrokenText-bytes-4096-87ce01a97fe1.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorEditUnbrokenText/bytes=4096: 1.297 ms before, 10.08 µs after, -99.22% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorEditUnbrokenText/bytes=100000">Editor edit unbroken text</strong><br>
Insert and backspace one character in 100000 bytes of existing unbroken text; update layout and view after each key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorEditUnbrokenText-bytes-100000-e62a15199e42.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorEditUnbrokenText/bytes=100000: 56.67 ms before, 18.11 µs after, -99.97% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorEditUnbrokenText/bytes=1000000">Editor edit unbroken text</strong><br>
Insert and backspace one character in 1000000 bytes of existing unbroken text; update layout and view after each key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorEditUnbrokenText-bytes-1000000-e42a0fa5f66d.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorEditUnbrokenText/bytes=1000000: 482.9 ms before, 161.8 µs after, -99.97% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorInput/bytes=256">Editor input</strong><br>
Type 250 bytes with hard newlines, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorInput-bytes-256-b6528ee6e145.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorInput/bytes=256: 5.271 ms before, 150.1 µs after, -97.15% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorInput/bytes=1024">Editor input</strong><br>
Type 1020 bytes with hard newlines, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorInput-bytes-1024-48ddabbd61d7.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorInput/bytes=1024: 92.08 ms before, 712.4 µs after, -99.23% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorInput/bytes=4096">Editor input</strong><br>
Type 4090 bytes with hard newlines, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorInput-bytes-4096-3ec3b006afc6.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorInput/bytes=4096: 1.56 s before, 4.04 ms after, -99.74% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorParagraphInput/bytes=256">Editor paragraph input</strong><br>
Type 250 bytes of a spaced paragraph, wrapping at 72 columns.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorParagraphInput-bytes-256-8236d0ba98ff.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorParagraphInput/bytes=256: 4.439 ms before, 395.2 µs after, -91.10% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorParagraphInput/bytes=1024">Editor paragraph input</strong><br>
Type 1020 bytes of a spaced paragraph, wrapping at 72 columns.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorParagraphInput-bytes-1024-6a6ccb836113.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorParagraphInput/bytes=1024: 75.56 ms before, 1.8 ms after, -97.62% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorParagraphInput/bytes=4096">Editor paragraph input</strong><br>
Type 4090 bytes of a spaced paragraph, wrapping at 72 columns.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorParagraphInput-bytes-4096-b320b87eaf3f.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorParagraphInput/bytes=4096: 1.37 s before, 8.541 ms after, -99.38% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorUnbrokenInput/bytes=256">Editor unbroken input</strong><br>
Type 256 unbroken ASCII characters, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorUnbrokenInput-bytes-256-2420cfe2fe98.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorUnbrokenInput/bytes=256: 4.584 ms before, 625.1 µs after, -86.36% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorUnbrokenInput/bytes=1024">Editor unbroken input</strong><br>
Type 1024 unbroken ASCII characters, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorUnbrokenInput-bytes-1024-14480fa12f5e.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorUnbrokenInput/bytes=1024: 72.71 ms before, 2.703 ms after, -96.28% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorUnbrokenInput/bytes=4096">Editor unbroken input</strong><br>
Type 4096 unbroken ASCII characters, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorUnbrokenInput-bytes-4096-277fee7a932b.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorUnbrokenInput/bytes=4096: 1.304 s before, 12.6 ms after, -99.03% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkGuiFlush/80x24">Gui flush</strong><br>
Flush two panes through a mock terminal at 80x24, changing selection and scroll position.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkGuiFlush-80x24-1e32dc0240a0.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkGuiFlush/80x24: 290.5 µs before, 207.8 µs after, -28.48% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkGuiFlush/160x48">Gui flush</strong><br>
Flush two panes through a mock terminal at 160x48, changing selection and scroll position.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkGuiFlush-160x48-a7299703a799.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkGuiFlush/160x48: 930.4 µs before, 553.6 µs after, -40.51% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkIncrementalSearchedDiff/plain">Incremental searched diff</strong><br>
Stream 8,192 searched lines in 64 batches, reading search status after each batch; plain view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkIncrementalSearchedDiff-plain-0faf60c398eb.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkIncrementalSearchedDiff/plain: 182.5 ms before, 15.49 ms after, -91.51% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkIncrementalSearchedDiff/wrapped">Incremental searched diff</strong><br>
Stream 8,192 searched lines in 64 batches, reading search status after each batch; wrapped view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkIncrementalSearchedDiff-wrapped-478204c5ff0e.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkIncrementalSearchedDiff/wrapped: 205.8 ms before, 18.94 ms after, -90.80% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkNearestSearchPosition/matches=100">Nearest search position</strong><br>
Find the nearest result near the end of 100 search matches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkNearestSearchPosition-matches-100-deecbafa7f0a.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkNearestSearchPosition/matches=100: 44.64 ns before, 4.17 ns after, -90.66% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkNearestSearchPosition/matches=1000">Nearest search position</strong><br>
Find the nearest result near the end of 1000 search matches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkNearestSearchPosition-matches-1000-9d66696ac3b1.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkNearestSearchPosition/matches=1000: 331.9 ns before, 7.017 ns after, -97.89% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkNearestSearchPosition/matches=10000">Nearest search position</strong><br>
Find the nearest result near the end of 10000 search matches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkNearestSearchPosition-matches-10000-6e376f106dbf.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkNearestSearchPosition/matches=10000: 3.242 µs before, 8.668 ns after, -99.73% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkNearestSearchPosition/matches=100000">Nearest search position</strong><br>
Find the nearest result near the end of 100000 search matches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkNearestSearchPosition-matches-100000-b3d9ccb08d3d.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkNearestSearchPosition/matches=100000: 38.52 µs before, 11.47 ns after, -99.97% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkPasteText/bytes=1000">Paste text</strong><br>
Insert 986 bytes into prefilled text with TypeString, including layout.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkPasteText-bytes-1000-d48054bd2a07.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkPasteText/bytes=1000: 131.9 µs before, 30.25 µs after, -77.06% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkPasteText/bytes=10000">Paste text</strong><br>
Insert 9996 bytes into prefilled text with TypeString, including layout.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkPasteText-bytes-10000-9620c0f37395.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkPasteText/bytes=10000: 6.433 ms before, 305.8 µs after, -95.25% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkPasteText/bytes=100000">Paste text</strong><br>
Insert 99994 bytes into prefilled text with TypeString, including layout.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkPasteText-bytes-100000-ce9f87d16d3e.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkPasteText/bytes=100000: 419.6 ms before, 2.912 ms after, -99.31% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=100/up=false">Scroll search position</strong><br>
Select a visible result after a distant scroll down across 100 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-100-up-false-38dc04d5d5f9.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=100/up=false: 56.04 ns before, 16.62 ns after, -70.34% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=100/up=true">Scroll search position</strong><br>
Select a visible result after a distant scroll up across 100 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-100-up-true-5fdc83f5a17f.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=100/up=true: 37.09 ns before, 16.07 ns after, -56.67% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=1000/up=false">Scroll search position</strong><br>
Select a visible result after a distant scroll down across 1000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-1000-up-false-a52956fd955b.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=1000/up=false: 431.1 ns before, 18.92 ns after, -95.61% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=1000/up=true">Scroll search position</strong><br>
Select a visible result after a distant scroll up across 1000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-1000-up-true-c1abe24c1406.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=1000/up=true: 315.6 ns before, 17.76 ns after, -94.37% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=10000/up=false">Scroll search position</strong><br>
Select a visible result after a distant scroll down across 10000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-10000-up-false-29e4a4854751.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=10000/up=false: 4.289 µs before, 22.09 ns after, -99.48% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=10000/up=true">Scroll search position</strong><br>
Select a visible result after a distant scroll up across 10000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-10000-up-true-f8b05b03fd52.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=10000/up=true: 3.099 µs before, 20.02 ns after, -99.35% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=100000/up=false">Scroll search position</strong><br>
Select a visible result after a distant scroll down across 100000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-100000-up-false-cea1f49e2a95.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=100000/up=false: 43.27 µs before, 25.25 ns after, -99.94% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=100000/up=true">Scroll search position</strong><br>
Select a visible result after a distant scroll up across 100000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-100000-up-true-e3f2a565cd05.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=100000/up=true: 36.81 µs before, 23.14 ns after, -99.94% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=a">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;a&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-mixed-case-query-a-c9e2f9c6ed26.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=a: 43.24 ms before, 12.12 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=MATCH">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;MATCH&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-mixed-case-query-MATCH-bdc3c481df01.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=MATCH: 20.31 ms before, 12.33 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=missing">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;missing&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-mixed-case-query-missing-19ddd862b43a.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=missing: 33.37 ms before, 12.22 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=performance-match">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;performance-match&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-mixed-case-query-performance-match-6c3b9aef3f96.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=performance-match: 46.48 ms before, 12.54 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=a">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;a&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-ordinary-query-a-3a97cbadacb7.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=a: 38.15 ms before, 12.02 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=MATCH">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;MATCH&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-ordinary-query-MATCH-d302bbec5080.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=MATCH: 20.68 ms before, 12.07 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=missing">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;missing&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-ordinary-query-missing-00d1a42effd7.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=missing: 28.75 ms before, 12.11 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=performance-match">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;performance-match&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-ordinary-query-performance-match-673b717a9ff6.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=performance-match: 41.93 ms before, 12.97 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchOrdinaryText/match">Search ordinary text</strong><br>
Search 2,000 ordinary diff lines for &quot;match&quot;; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchOrdinaryText-match-d6a687ba458f.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchOrdinaryText/match: 779.9 µs before, 12.21 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchOrdinaryText/MATCH">Search ordinary text</strong><br>
Search 2,000 ordinary diff lines for &quot;MATCH&quot;; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchOrdinaryText-MATCH-b940bf313969.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchOrdinaryText/MATCH: 537.1 µs before, 12.21 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchOrdinaryText/missing">Search ordinary text</strong><br>
Search 2,000 ordinary diff lines for &quot;missing&quot;; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchOrdinaryText-missing-c9f77e277cc1.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchOrdinaryText/missing: 706.9 µs before, 12.82 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchQueryLifetime/clear=false">Search query lifetime</strong><br>
Alternate aa, aaa, aaaa, and a queries across 2,000 lines of 200 a characters; keep search active; rendered layout and initial a query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchQueryLifetime-clear-false-082d6aa0b8b1.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchQueryLifetime/clear=false: 18.06 ms before, 2.974 ms after, -83.53% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchQueryLifetime/clear=true">Search query lifetime</strong><br>
Alternate aa, aaa, aaaa, and a queries across 2,000 lines of 200 a characters; clear search after each query; rendered layout and initial a query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchQueryLifetime-clear-true-7721ac7e3570.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchQueryLifetime/clear=true: 19.35 ms before, 11.56 ms after, -40.29% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchRepeatedPrefix/query=1">Search repeated prefix</strong><br>
Search 2,000 lines of 200 a characters for a; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchRepeatedPrefix-query-1-86e50afcc00a.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchRepeatedPrefix/query=1: 15.39 ms before, 12.09 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchRepeatedPrefix/query=32">Search repeated prefix</strong><br>
Search 2,000 lines of 200 a characters for a 32-character query ending in b; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchRepeatedPrefix-query-32-1774ace96361.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchRepeatedPrefix/query=32: 56.07 ms before, 11.98 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchRepeatedPrefix/query=128">Search repeated prefix</strong><br>
Search 2,000 lines of 200 a characters for a 128-character query ending in b; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchRepeatedPrefix-query-128-90a8ca3888d9.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchRepeatedPrefix/query=128: 157.8 ms before, 12.01 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaLineStart/lines=100">Text area line start</strong><br>
Move to the start of a line near the end of 100 prepared editor lines.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaLineStart-lines-100-420a9008d7d0.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaLineStart/lines=100: 1.239 µs before, 55.42 ns after, -95.53% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaLineStart/lines=1000">Text area line start</strong><br>
Move to the start of a line near the end of 1000 prepared editor lines.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaLineStart-lines-1000-1ebf05460fbb.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaLineStart/lines=1000: 12.12 µs before, 68.07 ns after, -99.44% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaLineStart/lines=10000">Text area line start</strong><br>
Move to the start of a line near the end of 10000 prepared editor lines.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaLineStart-lines-10000-f60249d9b1a4.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaLineStart/lines=10000: 163.7 µs before, 72.21 ns after, -99.96% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaVerticalKeypress/lines=100">Text area vertical keypress</strong><br>
Handle Up and Down through the default editor near the end of 100 lines, including rendering.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaVerticalKeypress-lines-100-e5df6524de16.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaVerticalKeypress/lines=100: 567.5 µs before, 189.4 ns after, -99.97% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaVerticalKeypress/lines=1000">Text area vertical keypress</strong><br>
Handle Up and Down through the default editor near the end of 1000 lines, including rendering.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaVerticalKeypress-lines-1000-1593b14dfa99.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaVerticalKeypress/lines=1000: 6.929 ms before, 209.7 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaVerticalKeypress/lines=10000">Text area vertical keypress</strong><br>
Handle Up and Down through the default editor near the end of 10000 lines, including rendering.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaVerticalKeypress-lines-10000-cc30cd026039.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaVerticalKeypress/lines=10000: 77.54 ms before, 236.2 ns after, -100.00% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaVerticalMovement/lines=100/bottom=false">Text area vertical movement</strong><br>
Move Up and Down in 100 prepared editor lines, near the beginning; rendering excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaVerticalMovement-lines-100-bottom-false-dd8d1b4ac796.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaVerticalMovement/lines=100/bottom=false: 122.5 ns before, 86.1 ns after, -29.71% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaVerticalMovement/lines=100/bottom=true">Text area vertical movement</strong><br>
Move Up and Down in 100 prepared editor lines, near the end; rendering excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaVerticalMovement-lines-100-bottom-true-76073ef179b9.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaVerticalMovement/lines=100/bottom=true: 4.677 µs before, 87.1 ns after, -98.14% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaVerticalMovement/lines=1000/bottom=false">Text area vertical movement</strong><br>
Move Up and Down in 1000 prepared editor lines, near the beginning; rendering excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaVerticalMovement-lines-1000-bottom-false-fd62ad794d50.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaVerticalMovement/lines=1000/bottom=false: 132.8 ns before, 104.5 ns after, -21.31% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaVerticalMovement/lines=1000/bottom=true">Text area vertical movement</strong><br>
Move Up and Down in 1000 prepared editor lines, near the end; rendering excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaVerticalMovement-lines-1000-bottom-true-c64cf5c10154.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaVerticalMovement/lines=1000/bottom=true: 48.19 µs before, 95.68 ns after, -99.80% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaVerticalMovement/lines=10000/bottom=false">Text area vertical movement</strong><br>
Move Up and Down in 10000 prepared editor lines, near the beginning; rendering excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaVerticalMovement-lines-10000-bottom-false-7a7289719d7c.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaVerticalMovement/lines=10000/bottom=false: 145.6 ns before, 120.9 ns after, -16.96% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkTextAreaVerticalMovement/lines=10000/bottom=true">Text area vertical movement</strong><br>
Move Up and Down in 10000 prepared editor lines, near the end; rendering excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkTextAreaVerticalMovement-lines-10000-bottom-true-9b30edcdb93b.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkTextAreaVerticalMovement/lines=10000/bottom=true: 409.6 µs before, 111.5 ns after, -99.97% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/80x24/empty">View draw</strong><br>
Redraw a 80x24 view with empty content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-80x24-empty-9dbba7c18dcf.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/80x24/empty: 55.35 µs before, 4.871 µs after, -91.20% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/80x24/full">View draw</strong><br>
Redraw a 80x24 view with full content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-80x24-full-694f9f37dcb1.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/80x24/full: 124.1 µs before, 40.47 µs after, -67.38% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/80x24/short">View draw</strong><br>
Redraw a 80x24 view with short content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-80x24-short-e515e1456e1f.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/80x24/short: 55.31 µs before, 4.824 µs after, -91.28% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/80x24/unicode">View draw</strong><br>
Redraw a 80x24 view with unicode content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-80x24-unicode-d63ab1e908fc.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/80x24/unicode: 145.2 µs before, 32.01 µs after, -77.94% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/160x48/empty">View draw</strong><br>
Redraw a 160x48 view with empty content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-160x48-empty-d5ce01334541.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/160x48/empty: 231 µs before, 16.4 µs after, -92.90% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/160x48/full">View draw</strong><br>
Redraw a 160x48 view with full content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-160x48-full-a5954f539bef.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/160x48/full: 520.2 µs before, 175.7 µs after, -66.23% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/160x48/short">View draw</strong><br>
Redraw a 160x48 view with short content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-160x48-short-953de0120bb0.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/160x48/short: 230.2 µs before, 16.45 µs after, -92.85% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/160x48/unicode">View draw</strong><br>
Redraw a 160x48 view with unicode content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-160x48-unicode-d880f113dcc7.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/160x48/unicode: 634 µs before, 129.9 µs after, -79.52% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWrite/colored">View write</strong><br>
Write 2,000 colored lines into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWrite-colored-317a17285f62.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWrite/colored: 9.348 ms before, 3.905 ms after, -58.22% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWrite/metadata">View write</strong><br>
Write 2,000 metadata lines into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWrite-metadata-597c3fcb95c8.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWrite/metadata: 5.894 ms before, 2.668 ms after, -54.73% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWrite/plain">View write</strong><br>
Write 2,000 plain lines into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWrite-plain-3b2f90c3be4d.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWrite/plain: 13.66 ms before, 2.865 ms after, -79.03% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWrite/unicode">View write</strong><br>
Write 2,000 unicode lines into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWrite-unicode-bf14bcb7023f.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWrite/unicode: 8.044 ms before, 3.964 ms after, -50.72% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWriteGitColoredDiff">View write git colored diff</strong><br>
Stream 50,000 mixed Unicode lines with separately colored Git prefixes into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWriteGitColoredDiff-bcacb328385e.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWriteGitColoredDiff: 496.9 ms before, 141.6 ms after, -71.51% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWriteLargeDiff/colored">View write large diff</strong><br>
Stream 50,000 colored diff lines into a new view, including parsing and allocation.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWriteLargeDiff-colored-41d27aa65e97.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWriteLargeDiff/colored: 257.2 ms before, 114.6 ms after, -55.43% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWriteLargeDiff/metadata">View write large diff</strong><br>
Stream 50,000 metadata diff lines into a new view, including parsing and allocation.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWriteLargeDiff-metadata-f87a83d5651f.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWriteLargeDiff/metadata: 176.1 ms before, 77.36 ms after, -56.06% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWriteLongLine">View write long line</strong><br>
Write one 13,001-byte line into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWriteLongLine-b0623613cc10.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWriteLongLine: 1.897 ms before, 225.5 µs after, -88.11% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkWriteToSearchedView">Write to searched view</strong><br>
Write 2,000 lines into a searched view and read search status.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkWriteToSearchedView-fb9fd7bdb721.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkWriteToSearchedView: 15.48 ms before, 3.907 ms after, -74.76% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/BenchmarkCommitDescriptionInput/bytes=256">Commit description input</strong><br>
Type 250 bytes through the commit-description editor.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-BenchmarkCommitDescriptionInput-bytes-256-12e08211e049.svg" width="420" height="60" align="middle" alt="gui/BenchmarkCommitDescriptionInput/bytes=256: 10.32 ms before, 160.1 µs after, -98.45% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/BenchmarkCommitDescriptionInput/bytes=1024">Commit description input</strong><br>
Type 1020 bytes through the commit-description editor.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-BenchmarkCommitDescriptionInput-bytes-1024-79a891150daf.svg" width="420" height="60" align="middle" alt="gui/BenchmarkCommitDescriptionInput/bytes=1024: 172.5 ms before, 753 µs after, -99.56% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/BenchmarkCommitDescriptionInput/bytes=4096">Commit description input</strong><br>
Type 4090 bytes through the commit-description editor.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-BenchmarkCommitDescriptionInput-bytes-4096-daee709fba0b.svg" width="420" height="60" align="middle" alt="gui/BenchmarkCommitDescriptionInput/bytes=4096: 3.047 s before, 4.438 ms after, -99.85% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/context/BenchmarkFilteredListNavigation/items=100">Filtered list navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count 100 filtered branches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-context-BenchmarkFilteredListNavigation-items-100-931d27cdab9d.svg" width="420" height="60" align="middle" alt="gui/context/BenchmarkFilteredListNavigation/items=100: 2.684 µs before, 37.48 ns after, -98.60% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/context/BenchmarkFilteredListNavigation/items=1000">Filtered list navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count 1000 filtered branches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-context-BenchmarkFilteredListNavigation-items-1000-d8d3b935c672.svg" width="420" height="60" align="middle" alt="gui/context/BenchmarkFilteredListNavigation/items=1000: 26.09 µs before, 37.64 ns after, -99.86% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/context/BenchmarkFilteredListNavigation/items=10000">Filtered list navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count 10000 filtered branches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-context-BenchmarkFilteredListNavigation-items-10000-1ac590a5eed4.svg" width="420" height="60" align="middle" alt="gui/context/BenchmarkFilteredListNavigation/items=10000: 282.7 µs before, 38.38 ns after, -99.99% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildCommitFileTree/files=100">Build commit file tree</strong><br>
Build and sort 100 reverse-ordered commit files through the file-tree construction entry point.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildCommitFileTree-files-100-94492e381348.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildCommitFileTree/files=100: 35.35 µs before, 31.82 µs after, -9.99% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildCommitFileTree/files=1000">Build commit file tree</strong><br>
Build and sort 1000 reverse-ordered commit files through the file-tree construction entry point.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildCommitFileTree-files-1000-48d1c38294f4.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildCommitFileTree/files=1000: 1.265 ms before, 330.9 µs after, -73.83% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildCommitFileTree/files=10000">Build commit file tree</strong><br>
Build and sort 10000 reverse-ordered commit files through the file-tree construction entry point.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildCommitFileTree-files-10000-7e4526f158ab.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildCommitFileTree/files=10000: 108.9 ms before, 3.627 ms after, -96.67% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=100/deep=false/commits=false">Build flat file trees</strong><br>
Build a flat working-file list of 100 files with one shared directory; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-100-deep-false-commits-false-7206e320bab3.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=100/deep=false/commits=false: 36.98 µs before, 24.45 µs after, -33.88% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=100/deep=false/commits=true">Build flat file trees</strong><br>
Build a flat commit-file list of 100 files with one shared directory; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-100-deep-false-commits-true-821cb7db900d.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=100/deep=false/commits=true: 35.14 µs before, 20.55 µs after, -41.50% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=100/deep=true/commits=false">Build flat file trees</strong><br>
Build a flat working-file list of 100 files with a distinct nested directory per file; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-100-deep-true-commits-false-ba0371e8179d.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=100/deep=true/commits=false: 97.17 µs before, 42.26 µs after, -56.51% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=100/deep=true/commits=true">Build flat file trees</strong><br>
Build a flat commit-file list of 100 files with a distinct nested directory per file; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-100-deep-true-commits-true-ff43ec78dd8b.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=100/deep=true/commits=true: 63.33 µs before, 38.81 µs after, -38.72% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=1000/deep=false/commits=false">Build flat file trees</strong><br>
Build a flat working-file list of 1000 files with one shared directory; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-1000-deep-false-commits-false-9c0c60cb1d8e.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=1000/deep=false/commits=false: 382.9 µs before, 247.2 µs after, -35.45% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=1000/deep=false/commits=true">Build flat file trees</strong><br>
Build a flat commit-file list of 1000 files with one shared directory; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-1000-deep-false-commits-true-ea05449e4282.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=1000/deep=false/commits=true: 1.167 ms before, 204.9 µs after, -82.44% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=1000/deep=true/commits=false">Build flat file trees</strong><br>
Build a flat working-file list of 1000 files with a distinct nested directory per file; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-1000-deep-true-commits-false-762ebd94be3b.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=1000/deep=true/commits=false: 1.175 ms before, 390.7 µs after, -66.75% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=1000/deep=true/commits=true">Build flat file trees</strong><br>
Build a flat commit-file list of 1000 files with a distinct nested directory per file; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-1000-deep-true-commits-true-62f4fd29490d.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=1000/deep=true/commits=true: 1.508 ms before, 349.9 µs after, -76.80% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=10000/deep=false/commits=false">Build flat file trees</strong><br>
Build a flat working-file list of 10000 files with one shared directory; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-10000-deep-false-commits-false-34324236e6d7.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=10000/deep=false/commits=false: 4.221 ms before, 2.753 ms after, -34.79% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=10000/deep=false/commits=true">Build flat file trees</strong><br>
Build a flat commit-file list of 10000 files with one shared directory; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-10000-deep-false-commits-true-72d31ef2ece8.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=10000/deep=false/commits=true: 95.91 ms before, 2.267 ms after, -97.64% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=10000/deep=true/commits=false">Build flat file trees</strong><br>
Build a flat working-file list of 10000 files with a distinct nested directory per file; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-10000-deep-true-commits-false-8aa7bf87f7c1.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=10000/deep=true/commits=false: 13.04 ms before, 5.493 ms after, -57.88% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildFlatFileTrees/files=10000/deep=true/commits=true">Build flat file trees</strong><br>
Build a flat commit-file list of 10000 files with a distinct nested directory per file; construction entry point included.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildFlatFileTrees-files-10000-deep-true-commits-true-ce20a1bbb6ff.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildFlatFileTrees/files=10000/deep=true/commits=true: 122.1 ms before, 4.996 ms after, -95.91% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeNavigation/files=100">File tree navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count rows in a 100-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeNavigation-files-100-38c1cc7e23dc.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeNavigation/files=100: 1.444 µs before, 28.76 ns after, -98.01% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeNavigation/files=1000">File tree navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count rows in a 1000-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeNavigation-files-1000-4bbb80b23742.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeNavigation/files=1000: 13.61 µs before, 29.05 ns after, -99.79% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeNavigation/files=10000">File tree navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count rows in a 10000-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeNavigation-files-10000-4afbd9af84e3.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeNavigation/files=10000: 135 µs before, 29.6 ns after, -99.98% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeRangeSelection/files=100">File tree range selection</strong><br>
Retrieve a range spanning a 100-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeRangeSelection-files-100-2fb911f791ed.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeRangeSelection/files=100: 19.6 µs before, 1.35 µs after, -93.11% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeRangeSelection/files=1000">File tree range selection</strong><br>
Retrieve a range spanning a 1000-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeRangeSelection-files-1000-f30b72fdee2e.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeRangeSelection/files=1000: 1.655 ms before, 13.54 µs after, -99.18% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeRangeSelection/files=10000">File tree range selection</strong><br>
Retrieve a range spanning a 10000-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeRangeSelection-files-10000-cd1741432ddd.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeRangeSelection/files=10000: 164.4 ms before, 150.5 µs after, -99.91% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFlatTreeRefresh/files=100/commits=false">Flat tree refresh</strong><br>
Refresh a flat working-file view model with 100 files, including selection recovery.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFlatTreeRefresh-files-100-commits-false-0ba71dced7ba.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFlatTreeRefresh/files=100/commits=false: 94.11 µs before, 39.48 µs after, -58.05% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFlatTreeRefresh/files=100/commits=true">Flat tree refresh</strong><br>
Refresh a flat commit-file view model with 100 files, including selection recovery.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFlatTreeRefresh-files-100-commits-true-45734e52dcdb.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFlatTreeRefresh/files=100/commits=true: 53.03 µs before, 33.4 µs after, -37.01% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFlatTreeRefresh/files=10000/commits=false">Flat tree refresh</strong><br>
Refresh a flat working-file view model with 10000 files, including selection recovery.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFlatTreeRefresh-files-10000-commits-false-7e2b628199c5.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFlatTreeRefresh/files=10000/commits=false: 12.6 ms before, 5.37 ms after, -57.38% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFlatTreeRefresh/files=10000/commits=true">Flat tree refresh</strong><br>
Refresh a flat commit-file view model with 10000 files, including selection recovery.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFlatTreeRefresh-files-10000-commits-true-4c2db369f6be.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFlatTreeRefresh/files=10000/commits=true: 151.5 ms before, 4.627 ms after, -96.95% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=100">Refresh after directory removal</strong><br>
Refresh a tree after removing a 100-file directory; fixture setup excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRefreshAfterDirectoryRemoval-files-100-90069f748cab.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=100: 985.8 µs before, 42.03 µs after, -95.74% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=1000">Refresh after directory removal</strong><br>
Refresh a tree after removing a 1000-file directory; fixture setup excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRefreshAfterDirectoryRemoval-files-1000-9b9906ac9c86.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=1000: 101.8 ms before, 382.4 µs after, -99.62% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=5000">Refresh after directory removal</strong><br>
Refresh a tree after removing a 5000-file directory; fixture setup excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRefreshAfterDirectoryRemoval-files-5000-0b9d3a6459a7.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=5000: 2.718 s before, 2.188 ms after, -99.92% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=100">Restore after single removal</strong><br>
Restore selection after removing one file from 100 remaining files.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreAfterSingleRemoval-files-100-85d9dfe6be51.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=100: 7.786 µs before, 1.963 µs after, -74.79% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=1000">Restore after single removal</strong><br>
Restore selection after removing one file from 1000 remaining files.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreAfterSingleRemoval-files-1000-99f5712ef27d.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=1000: 78.34 µs before, 19.56 µs after, -75.03% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=5000">Restore after single removal</strong><br>
Restore selection after removing one file from 5000 remaining files.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreAfterSingleRemoval-files-5000-b796c4b2f690.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=5000: 435.7 µs before, 109.6 µs after, -74.85% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=100/removed=false">Restore file selection</strong><br>
Restore selection among 100 files; selected file retained.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-100-removed-false-f21bbf9cfe1b.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=100/removed=false: 163.9 ns before, 37.98 ns after, -76.83% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=100/removed=true">Restore file selection</strong><br>
Restore selection among 100 files; selected file removed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-100-removed-true-d75a1de15d91.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=100/removed=true: 789.7 µs before, 12.37 µs after, -98.43% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=1000/removed=false">Restore file selection</strong><br>
Restore selection among 1000 files; selected file retained.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-1000-removed-false-fab28116babd.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=1000/removed=false: 163.5 ns before, 39.95 ns after, -75.57% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=1000/removed=true">Restore file selection</strong><br>
Restore selection among 1000 files; selected file removed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-1000-removed-true-80bd973c8dd6.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=1000/removed=true: 80.31 ms before, 124.1 µs after, -99.85% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=5000/removed=false">Restore file selection</strong><br>
Restore selection among 5000 files; selected file retained.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-5000-removed-false-173d0156293d.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=5000/removed=false: 176.4 ns before, 43.94 ns after, -75.09% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=5000/removed=true">Restore file selection</strong><br>
Restore selection among 5000 files; selected file removed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-5000-removed-true-0248f58474a5.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=5000/removed=true: 2.163 s before, 720.9 µs after, -99.97% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkSmallTreeRangeSelection/files=100/cold=false/commits=false">Small tree range selection</strong><br>
Select one row near the end of a 100-file working-file tree; visibility unchanged.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkSmallTreeRangeSelection-files-100-cold-false-commits-false-1bb3ba10a52b.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkSmallTreeRangeSelection/files=100/cold=false/commits=false: 693.1 ns before, 31.99 ns after, -95.38% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkSmallTreeRangeSelection/files=100/cold=false/commits=true">Small tree range selection</strong><br>
Select one row near the end of a 100-file commit-file tree; visibility unchanged.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkSmallTreeRangeSelection-files-100-cold-false-commits-true-db5c0ee6364f.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkSmallTreeRangeSelection/files=100/cold=false/commits=true: 694.2 ns before, 31.79 ns after, -95.42% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkSmallTreeRangeSelection/files=100/cold=true/commits=false">Small tree range selection</strong><br>
Select one row near the end of a 100-file working-file tree; toggle a directory twice before each selection to invalidate visibility.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkSmallTreeRangeSelection-files-100-cold-true-commits-false-f981c5cb0789.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkSmallTreeRangeSelection/files=100/cold=true/commits=false: 710.3 ns before, 389.8 ns after, -45.12% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkSmallTreeRangeSelection/files=100/cold=true/commits=true">Small tree range selection</strong><br>
Select one row near the end of a 100-file commit-file tree; toggle a directory twice before each selection to invalidate visibility.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkSmallTreeRangeSelection-files-100-cold-true-commits-true-4e21d1debce7.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkSmallTreeRangeSelection/files=100/cold=true/commits=true: 733.6 ns before, 368.8 ns after, -49.73% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkSmallTreeRangeSelection/files=10000/cold=false/commits=false">Small tree range selection</strong><br>
Select one row near the end of a 10000-file working-file tree; visibility unchanged.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkSmallTreeRangeSelection-files-10000-cold-false-commits-false-8976ea9ca8d3.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkSmallTreeRangeSelection/files=10000/cold=false/commits=false: 61.91 µs before, 34.11 ns after, -99.94% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkSmallTreeRangeSelection/files=10000/cold=false/commits=true">Small tree range selection</strong><br>
Select one row near the end of a 10000-file commit-file tree; visibility unchanged.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkSmallTreeRangeSelection-files-10000-cold-false-commits-true-fd0e9c8d1cc1.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkSmallTreeRangeSelection/files=10000/cold=false/commits=true: 61.87 µs before, 34.2 ns after, -99.94% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkSmallTreeRangeSelection/files=10000/cold=true/commits=false">Small tree range selection</strong><br>
Select one row near the end of a 10000-file working-file tree; toggle a directory twice before each selection to invalidate visibility.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkSmallTreeRangeSelection-files-10000-cold-true-commits-false-31ef0aff52e0.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkSmallTreeRangeSelection/files=10000/cold=true/commits=false: 62.36 µs before, 29.92 µs after, -52.01% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkSmallTreeRangeSelection/files=10000/cold=true/commits=true">Small tree range selection</strong><br>
Select one row near the end of a 10000-file commit-file tree; toggle a directory twice before each selection to invalidate visibility.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkSmallTreeRangeSelection-files-10000-cold-true-commits-true-2b4b11b2bb98.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkSmallTreeRangeSelection/files=10000/cold=true/commits=true: 62.27 µs before, 29.67 µs after, -52.35% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=1000/main=absent/markers=false">Commit viewport</strong><br>
Render 40 commit rows from 1000 loaded commits without a graph; local main tip absent, other branch marker absent.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-1000-main-absent-markers-false-061cef0f8e32.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=1000/main=absent/markers=false: 15.55 µs before, 15.99 µs after, +2.80% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=1000/main=absent/markers=true">Commit viewport</strong><br>
Render 40 commit rows from 1000 loaded commits without a graph; local main tip absent, other branch marker visible.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-1000-main-absent-markers-true-d744a829355d.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=1000/main=absent/markers=true: 16.17 µs before, 18.2 µs after, +12.57% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=1000/main=merged/markers=false">Commit viewport</strong><br>
Render 40 commit rows from 1000 loaded commits without a graph; local main tip merged, other branch marker absent.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-1000-main-merged-markers-false-eab34025f752.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=1000/main=merged/markers=false: 15.8 µs before, 15.93 µs after, +0.78% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=1000/main=merged/markers=true">Commit viewport</strong><br>
Render 40 commit rows from 1000 loaded commits without a graph; local main tip merged, other branch marker visible.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-1000-main-merged-markers-true-9c39391fe5ed.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=1000/main=merged/markers=true: 16.55 µs before, 20.68 µs after, +24.94% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=1000/main=unmerged/markers=false">Commit viewport</strong><br>
Render 40 commit rows from 1000 loaded commits without a graph; local main tip unmerged, other branch marker absent.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-1000-main-unmerged-markers-false-1b61d22fbce9.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=1000/main=unmerged/markers=false: 15.65 µs before, 16.03 µs after, +2.40% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=1000/main=unmerged/markers=true">Commit viewport</strong><br>
Render 40 commit rows from 1000 loaded commits without a graph; local main tip unmerged, other branch marker visible.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-1000-main-unmerged-markers-true-ca7f903c4a32.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=1000/main=unmerged/markers=true: 16.53 µs before, 94.3 µs after, +470.35% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=10000/main=absent/markers=false">Commit viewport</strong><br>
Render 40 commit rows from 10000 loaded commits without a graph; local main tip absent, other branch marker absent.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-10000-main-absent-markers-false-87d6d22da058.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=10000/main=absent/markers=false: 67.43 µs before, 70.21 µs after, +4.12% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=10000/main=absent/markers=true">Commit viewport</strong><br>
Render 40 commit rows from 10000 loaded commits without a graph; local main tip absent, other branch marker visible.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-10000-main-absent-markers-true-f689d61de812.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=10000/main=absent/markers=true: 73.31 µs before, 111 µs after, +51.34% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=10000/main=merged/markers=false">Commit viewport</strong><br>
Render 40 commit rows from 10000 loaded commits without a graph; local main tip merged, other branch marker absent.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-10000-main-merged-markers-false-7db830a47337.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=10000/main=merged/markers=false: 69.52 µs before, 72.14 µs after, +3.78% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=10000/main=merged/markers=true">Commit viewport</strong><br>
Render 40 commit rows from 10000 loaded commits without a graph; local main tip merged, other branch marker visible.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-10000-main-merged-markers-true-1bbc6df6c572.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=10000/main=merged/markers=true: 70.39 µs before, 127.7 µs after, +81.47% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=10000/main=unmerged/markers=false">Commit viewport</strong><br>
Render 40 commit rows from 10000 loaded commits without a graph; local main tip unmerged, other branch marker absent.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-10000-main-unmerged-markers-false-1d3b80a7921e.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=10000/main=unmerged/markers=false: 68.33 µs before, 69.57 µs after, +1.81% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=10000/main=unmerged/markers=true">Commit viewport</strong><br>
Render 40 commit rows from 10000 loaded commits without a graph; local main tip unmerged, other branch marker visible.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-10000-main-unmerged-markers-true-f412a97d62f9.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=10000/main=unmerged/markers=true: 72.06 µs before, 894.3 µs after, +1,140.96% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=100000/main=absent/markers=false">Commit viewport</strong><br>
Render 40 commit rows from 100000 loaded commits without a graph; local main tip absent, other branch marker absent.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-100000-main-absent-markers-false-245e4b3f26cd.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=100000/main=absent/markers=false: 992.8 µs before, 1.053 ms after, +6.04% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=100000/main=absent/markers=true">Commit viewport</strong><br>
Render 40 commit rows from 100000 loaded commits without a graph; local main tip absent, other branch marker visible.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-100000-main-absent-markers-true-b0e67a2a24c6.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=100000/main=absent/markers=true: 996.4 µs before, 1.805 ms after, +81.14% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=100000/main=merged/markers=false">Commit viewport</strong><br>
Render 40 commit rows from 100000 loaded commits without a graph; local main tip merged, other branch marker absent.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-100000-main-merged-markers-false-09456e599312.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=100000/main=merged/markers=false: 1.002 ms before, 1.091 ms after, +8.84% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=100000/main=merged/markers=true">Commit viewport</strong><br>
Render 40 commit rows from 100000 loaded commits without a graph; local main tip merged, other branch marker visible.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-100000-main-merged-markers-true-aeb80d51791e.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=100000/main=merged/markers=true: 1.008 ms before, 2.065 ms after, +104.93% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=100000/main=unmerged/markers=false">Commit viewport</strong><br>
Render 40 commit rows from 100000 loaded commits without a graph; local main tip unmerged, other branch marker absent.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-100000-main-unmerged-markers-false-4cd08242aa10.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=100000/main=unmerged/markers=false: 998.1 µs before, 1.077 ms after, +7.93% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/BenchmarkCommitViewport/commits=100000/main=unmerged/markers=true">Commit viewport</strong><br>
Render 40 commit rows from 100000 loaded commits without a graph; local main tip unmerged, other branch marker visible.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-BenchmarkCommitViewport-commits-100000-main-unmerged-markers-true-dbd1d867c9ed.svg" width="420" height="60" align="middle" alt="gui/presentation/BenchmarkCommitViewport/commits=100000/main=unmerged/markers=true: 998.6 µs before, 11.03 ms after, +1,004.72% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphColumnAllocation/columns=4">Graph column allocation</strong><br>
Lay out a graph row with 4 incoming columns and its current commit in the final column.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphColumnAllocation-columns-4-11e773dc51c4.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphColumnAllocation/columns=4: 355 ns before, 260.2 ns after, -26.70% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphColumnAllocation/columns=16">Graph column allocation</strong><br>
Lay out a graph row with 16 incoming columns and its current commit in the final column.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphColumnAllocation-columns-16-c4afac1ed186.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphColumnAllocation/columns=16: 3.515 µs before, 996.9 ns after, -71.64% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphColumnAllocation/columns=64">Graph column allocation</strong><br>
Lay out a graph row with 64 incoming columns and its current commit in the final column.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphColumnAllocation-columns-64-766d0a9e52d1.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphColumnAllocation/columns=64: 22.5 µs before, 3.45 µs after, -84.67% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphColumnAllocation/columns=256">Graph column allocation</strong><br>
Lay out a graph row with 256 incoming columns and its current commit in the final column.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphColumnAllocation-columns-256-a58db7c1f45a.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphColumnAllocation/columns=256: 202.7 µs before, 12.9 µs after, -93.63% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=1">Graph viewport</strong><br>
Render 1 cached linear graph rows with GOMAXPROCS=1; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-1-31a8a4f2cf50.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=1: 1.436 µs before, 318.2 ns after, -77.84% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=1/cpu=4">Graph viewport</strong><br>
Render 1 cached linear graph rows with GOMAXPROCS=4; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-1-cpu-4-c9154de69360.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=1/cpu=4: 2.34 µs before, 237.7 ns after, -89.84% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=1/cpu=16">Graph viewport</strong><br>
Render 1 cached linear graph rows with GOMAXPROCS=16; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-1-cpu-16-4a5460c62e86.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=1/cpu=16: 5.627 µs before, 257.5 ns after, -95.42% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=8">Graph viewport</strong><br>
Render 8 cached linear graph rows with GOMAXPROCS=1; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-8-6e146f61f0fd.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=8: 3.137 µs before, 2.086 µs after, -33.50% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=8/cpu=4">Graph viewport</strong><br>
Render 8 cached linear graph rows with GOMAXPROCS=4; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-8-cpu-4-96c5a9c63b70.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=8/cpu=4: 6.576 µs before, 1.512 µs after, -77.01% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=8/cpu=16">Graph viewport</strong><br>
Render 8 cached linear graph rows with GOMAXPROCS=16; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-8-cpu-16-d219b90130fb.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=8/cpu=16: 8.121 µs before, 1.623 µs after, -80.01% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=40">Graph viewport</strong><br>
Render 40 cached linear graph rows with GOMAXPROCS=1; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-40-69e8e46ca6b5.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=40: 10.55 µs before, 10.09 µs after, -4.33% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=40/cpu=4">Graph viewport</strong><br>
Render 40 cached linear graph rows with GOMAXPROCS=4; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-40-cpu-4-634a2e865bc0.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=40/cpu=4: 12 µs before, 7.483 µs after, -37.66% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=40/cpu=16">Graph viewport</strong><br>
Render 40 cached linear graph rows with GOMAXPROCS=16; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-40-cpu-16-322936f0f61a.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=40/cpu=16: 24.86 µs before, 7.867 µs after, -68.36% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=120">Graph viewport</strong><br>
Render 120 cached linear graph rows with GOMAXPROCS=1; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-120-88aea746ad25.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=120: 28.68 µs before, 27.05 µs after, -5.66% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=120/cpu=4">Graph viewport</strong><br>
Render 120 cached linear graph rows with GOMAXPROCS=4; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-120-cpu-4-07cc0516bdfe.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=120/cpu=4: 19.51 µs before, 20.18 µs after, +3.44% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=120/cpu=16">Graph viewport</strong><br>
Render 120 cached linear graph rows with GOMAXPROCS=16; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-false-rows-120-cpu-16-a6deec92b3e9.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=false/rows=120/cpu=16: 37.16 µs before, 25.15 µs after, -32.31% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=1">Graph viewport</strong><br>
Render 1 cached branching graph rows with GOMAXPROCS=1; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-1-0894c94dd681.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=1: 4.4 µs before, 3.086 µs after, -29.86% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=1/cpu=4">Graph viewport</strong><br>
Render 1 cached branching graph rows with GOMAXPROCS=4; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-1-cpu-4-a41752fe2467.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=1/cpu=4: 5.206 µs before, 2.245 µs after, -56.88% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=1/cpu=16">Graph viewport</strong><br>
Render 1 cached branching graph rows with GOMAXPROCS=16; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-1-cpu-16-22b57db1fa7a.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=1/cpu=16: 9.955 µs before, 2.727 µs after, -72.61% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=8">Graph viewport</strong><br>
Render 8 cached branching graph rows with GOMAXPROCS=1; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-8-1fb13b64f058.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=8: 32.32 µs before, 31.29 µs after, -3.19% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=8/cpu=4">Graph viewport</strong><br>
Render 8 cached branching graph rows with GOMAXPROCS=4; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-8-cpu-4-4cdf5dadb7a7.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=8/cpu=4: 21.05 µs before, 21.67 µs after, +2.93% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=8/cpu=16">Graph viewport</strong><br>
Render 8 cached branching graph rows with GOMAXPROCS=16; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-8-cpu-16-39690868d3b6.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=8/cpu=16: 36.9 µs before, 27.82 µs after, -24.59% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=40">Graph viewport</strong><br>
Render 40 cached branching graph rows with GOMAXPROCS=1; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-40-4ac0a67ef512.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=40: 142.1 µs before, 133.2 µs after, -6.23% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=40/cpu=4">Graph viewport</strong><br>
Render 40 cached branching graph rows with GOMAXPROCS=4; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-40-cpu-4-7f650e2fae5a.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=40/cpu=4: 76.99 µs before, 68.71 µs after, -10.75% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=40/cpu=16">Graph viewport</strong><br>
Render 40 cached branching graph rows with GOMAXPROCS=16; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-40-cpu-16-853d7007e895.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=40/cpu=16: 176.3 µs before, 110.2 µs after, -37.48% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=120">Graph viewport</strong><br>
Render 120 cached branching graph rows with GOMAXPROCS=1; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-120-93866275c087.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=120: 547.1 µs before, 492.7 µs after, -9.94% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=120/cpu=4">Graph viewport</strong><br>
Render 120 cached branching graph rows with GOMAXPROCS=4; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-120-cpu-4-ea6c6ff52d39.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=120/cpu=4: 302.7 µs before, 262 µs after, -13.44% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=120/cpu=16">Graph viewport</strong><br>
Render 120 cached branching graph rows with GOMAXPROCS=16; layout excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkGraphViewport-branches-true-rows-120-cpu-16-686bddcfe924.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkGraphViewport/branches=true/rows=120/cpu=16: 493.2 µs before, 438.9 µs after, -11.02% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraph">Render commit graph</strong><br>
Render a graph of 50 commits generated with seed 1234.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraph-882b6dc87d2f.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraph: 90.18 µs before, 60.41 µs after, -33.01% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=box/selected=false">Render commit graph variants</strong><br>
Render 50 commits generated with seed 1234, using box symbols; no selected commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-50-symbols-box-selected-false-861282eb9fd0.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=box/selected=false: 88.51 µs before, 60.06 µs after, -32.14% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=box/selected=true">Render commit graph variants</strong><br>
Render 50 commits generated with seed 1234, using box symbols; select the middle commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-50-symbols-box-selected-true-0621cb2aae16.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=box/selected=true: 93.26 µs before, 65.39 µs after, -29.88% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=branch/selected=false">Render commit graph variants</strong><br>
Render 50 commits generated with seed 1234, using branch symbols; no selected commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-50-symbols-branch-selected-false-8481d564fc12.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=branch/selected=false: 94.05 µs before, 63.76 µs after, -32.21% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=branch/selected=true">Render commit graph variants</strong><br>
Render 50 commits generated with seed 1234, using branch symbols; select the middle commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-50-symbols-branch-selected-true-c1fa9bf50b45.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=branch/selected=true: 100.7 µs before, 68.72 µs after, -31.79% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=box/selected=false">Render commit graph variants</strong><br>
Render 500 commits generated with seed 1234, using box symbols; no selected commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-500-symbols-box-selected-false-45087f7a544e.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=box/selected=false: 6.3 ms before, 3.501 ms after, -44.44% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=box/selected=true">Render commit graph variants</strong><br>
Render 500 commits generated with seed 1234, using box symbols; select the middle commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-500-symbols-box-selected-true-fc79f91a2666.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=box/selected=true: 6.454 ms before, 3.544 ms after, -45.09% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=branch/selected=false">Render commit graph variants</strong><br>
Render 500 commits generated with seed 1234, using branch symbols; no selected commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-500-symbols-branch-selected-false-875d1a3719cf.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=branch/selected=false: 6.622 ms before, 3.771 ms after, -43.05% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=branch/selected=true">Render commit graph variants</strong><br>
Render 500 commits generated with seed 1234, using branch symbols; select the middle commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-500-symbols-branch-selected-true-0d67c4979a57.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=branch/selected=true: 6.655 ms before, 3.868 ms after, -41.87% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="utils/BenchmarkEscapeSpecialChars/controls">Escape special chars</strong><br>
Escape control characters in a filename containing tabs and newlines.</p></td>
<td><img src="docs-master/dev/performance-comparisons/utils-BenchmarkEscapeSpecialChars-controls-d333fb93e862.svg" width="420" height="60" align="middle" alt="utils/BenchmarkEscapeSpecialChars/controls: 1.118 µs before, 69.37 ns after, -93.80% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="utils/BenchmarkEscapeSpecialChars/ordinary">Escape special chars</strong><br>
Escape control characters in a plain filename.</p></td>
<td><img src="docs-master/dev/performance-comparisons/utils-BenchmarkEscapeSpecialChars-ordinary-df5431b11a68.svg" width="420" height="60" align="middle" alt="utils/BenchmarkEscapeSpecialChars/ordinary: 996.6 ns before, 14.39 ns after, -98.56% change"></td>
</tr>
</tbody>
</table>

<!-- performance-summary:end -->

---

<div align="center">
<sup>Special thanks to:</sup>
<br>
<br>
<a href="https://www.warp.dev/?utm_source=github&utm_medium=referral&utm_campaign=lazygit_20231023">
  <div>
    <img src="https://github.com/warpdotdev/brand-assets/blob/main/Github/Sponsor/Warp-Github-LG-02.png?raw=true" width="400" alt="Warp">
  </div>
  <b>Warp, the intelligent terminal</b>
  <br>
  <b>Available for macOS and Linux</b>
  <br>
  <div>
    <sup>Visit warp.dev to learn more.</sup>
  </div>
</a>
<br>
<hr>
<a href="https://tuple.app/lazygit">
  <div>
    <img src="../assets/tuple.png" width="400" alt="Tuple">
  </div>
  <b>Tuple, the premier screen sharing app for developers on macOS and Windows.</b>
</a>
<br>
<hr>
<br>
<a href="https://www.subble.com">
  <div>
    <img src="../assets/subble.webp" width="400" alt="Subble">
  </div>
  <b>I (Jesse) co-founded Subble to save your company time and money by finding unused and over-provisioned SaaS licences. Check it out!</b>
</a>
<br>

<hr>
</div>

<p align="center">
  <img width="536" src="https://user-images.githubusercontent.com/8456633/174470852-339b5011-5800-4bb9-a628-ff230aa8cd4e.png">
</p>

<div align="center">

A simple terminal UI for git commands
<br/>

[![GitHub Releases](https://img.shields.io/github/downloads/jesseduffield/lazygit/total)](https://github.com/jesseduffield/lazygit/releases) [![golangci-lint](https://img.shields.io/badge/linted%20by-golangci--lint-brightgreen)](https://golangci-lint.run/) [![GitHub tag](https://img.shields.io/github/v/tag/jesseduffield/lazygit?color=blue)](https://github.com/jesseduffield/lazygit/releases/latest) [![homebrew](https://img.shields.io/homebrew/v/lazygit?color=blue)](https://formulae.brew.sh/formula/lazygit)

<!-- demo: commit_and_push -->
<video src="https://github.com/user-attachments/assets/031799d1-740d-4cda-90cf-689513fa9860" controls></video>

</div>

## Sponsors

<p align="center">
 Maintenance of this project is made possible by all the <a href="https://github.com/jesseduffield/lazygit/graphs/contributors">contributors</a> and <a href="https://github.com/sponsors/jesseduffield">sponsors</a>. If you'd like to sponsor this project and have your avatar or company logo appear below <a href="https://github.com/sponsors/jesseduffield">click here</a>. 💙
</p>

<p align="center">
<!-- sponsors --><a href="https://github.com/intabulas"><img src="https:&#x2F;&#x2F;github.com&#x2F;intabulas.png" width="60px" alt="User avatar: Mark Lussier" /></a><a href="https://github.com/peppy"><img src="https:&#x2F;&#x2F;github.com&#x2F;peppy.png" width="60px" alt="User avatar: Dean Herbert" /></a><a href="https://github.com/piot"><img src="https:&#x2F;&#x2F;github.com&#x2F;piot.png" width="60px" alt="User avatar: Peter Bjorklund" /></a><a href="https://github.com/oliverguenther"><img src="https:&#x2F;&#x2F;github.com&#x2F;oliverguenther.png" width="60px" alt="User avatar: Oliver Günther" /></a><a href="https://github.com/pawanjay176"><img src="https:&#x2F;&#x2F;github.com&#x2F;pawanjay176.png" width="60px" alt="User avatar: Pawan Dhananjay" /></a><a href="https://github.com/naoey"><img src="https:&#x2F;&#x2F;github.com&#x2F;naoey.png" width="60px" alt="User avatar: " /></a><a href="https://github.com/carstengehling"><img src="https:&#x2F;&#x2F;github.com&#x2F;carstengehling.png" width="60px" alt="User avatar: Carsten Gehling" /></a><a href="https://github.com/Xetera"><img src="https:&#x2F;&#x2F;github.com&#x2F;Xetera.png" width="60px" alt="User avatar: Ali" /></a><a href="https://github.com/nartc"><img src="https:&#x2F;&#x2F;github.com&#x2F;nartc.png" width="60px" alt="User avatar: Chau Tran" /></a><a href="https://github.com/matejcik"><img src="https:&#x2F;&#x2F;github.com&#x2F;matejcik.png" width="60px" alt="User avatar: matejcik" /></a><a href="https://github.com/lucatume"><img src="https:&#x2F;&#x2F;github.com&#x2F;lucatume.png" width="60px" alt="User avatar: theAverageDev (Luca Tumedei)" /></a><a href="https://github.com/escrafford"><img src="https:&#x2F;&#x2F;github.com&#x2F;escrafford.png" width="60px" alt="User avatar: Erik Scrafford" /></a><a href="https://github.com/IvanZuy"><img src="https:&#x2F;&#x2F;github.com&#x2F;IvanZuy.png" width="60px" alt="User avatar: " /></a><a href="https://github.com/ava1ar"><img src="https:&#x2F;&#x2F;github.com&#x2F;ava1ar.png" width="60px" alt="User avatar: Alex Stelmachonak" /></a><a href="https://github.com/pedropombeiro"><img src="https:&#x2F;&#x2F;github.com&#x2F;pedropombeiro.png" width="60px" alt="User avatar: Pedro Pombeiro" /></a><a href="https://github.com/bburgy"><img src="https:&#x2F;&#x2F;github.com&#x2F;bburgy.png" width="60px" alt="User avatar: Burgy Benjamin" /></a><a href="https://github.com/JoeKlemmer"><img src="https:&#x2F;&#x2F;github.com&#x2F;JoeKlemmer.png" width="60px" alt="User avatar: Joe Klemmer" /></a><a href="https://github.com/tobi"><img src="https:&#x2F;&#x2F;github.com&#x2F;tobi.png" width="60px" alt="User avatar: Tobias Lütke" /></a><a href="https://github.com/benbfortis"><img src="https:&#x2F;&#x2F;github.com&#x2F;benbfortis.png" width="60px" alt="User avatar: Ben Beaumont" /></a><a href="https://github.com/jakewarren"><img src="https:&#x2F;&#x2F;github.com&#x2F;jakewarren.png" width="60px" alt="User avatar: " /></a><a href="https://github.com/tgpholly"><img src="https:&#x2F;&#x2F;github.com&#x2F;tgpholly.png" width="60px" alt="User avatar: Holly" /></a><a href="https://github.com/Tommylans"><img src="https:&#x2F;&#x2F;github.com&#x2F;Tommylans.png" width="60px" alt="User avatar: Tom Lanser" /></a><a href="https://github.com/socketbox"><img src="https:&#x2F;&#x2F;github.com&#x2F;socketbox.png" width="60px" alt="User avatar: Casey Boettcher" /></a><a href="https://github.com/bitprophet"><img src="https:&#x2F;&#x2F;github.com&#x2F;bitprophet.png" width="60px" alt="User avatar: Jeff Forcier" /></a><a href="https://github.com/Novakov"><img src="https:&#x2F;&#x2F;github.com&#x2F;Novakov.png" width="60px" alt="User avatar: Maciej T. Nowak" /></a><a href="https://github.com/jhillyerd"><img src="https:&#x2F;&#x2F;github.com&#x2F;jhillyerd.png" width="60px" alt="User avatar: James Hillyerd" /></a><a href="https://github.com/nekhaevskiy"><img src="https:&#x2F;&#x2F;github.com&#x2F;nekhaevskiy.png" width="60px" alt="User avatar: Yury" /></a><a href="https://github.com/reivilibre"><img src="https:&#x2F;&#x2F;github.com&#x2F;reivilibre.png" width="60px" alt="User avatar: Olivier reivilibre" /></a><a href="https://github.com/BSteffaniak"><img src="https:&#x2F;&#x2F;github.com&#x2F;BSteffaniak.png" width="60px" alt="User avatar: Braden Steffaniak" /></a><a href="https://github.com/smangels"><img src="https:&#x2F;&#x2F;github.com&#x2F;smangels.png" width="60px" alt="User avatar: Sebastian" /></a><a href="https://github.com/amslezak"><img src="https:&#x2F;&#x2F;github.com&#x2F;amslezak.png" width="60px" alt="User avatar: Andy Slezak" /></a><a href="https://github.com/mkock"><img src="https:&#x2F;&#x2F;github.com&#x2F;mkock.png" width="60px" alt="User avatar: Martin Kock" /></a><a href="https://github.com/danielkokott"><img src="https:&#x2F;&#x2F;github.com&#x2F;danielkokott.png" width="60px" alt="User avatar: Daniel Kokott" /></a><a href="https://github.com/heijmans"><img src="https:&#x2F;&#x2F;github.com&#x2F;heijmans.png" width="60px" alt="User avatar: Jan Heijmans" /></a><a href="https://github.com/knowald"><img src="https:&#x2F;&#x2F;github.com&#x2F;knowald.png" width="60px" alt="User avatar: Kevin Nowald" /></a><a href="https://github.com/ethanjli"><img src="https:&#x2F;&#x2F;github.com&#x2F;ethanjli.png" width="60px" alt="User avatar: Ethan Li" /></a><a href="https://github.com/phubaba"><img src="https:&#x2F;&#x2F;github.com&#x2F;phubaba.png" width="60px" alt="User avatar: Robert Forler" /></a><a href="https://github.com/neunkasulle"><img src="https:&#x2F;&#x2F;github.com&#x2F;neunkasulle.png" width="60px" alt="User avatar: Jan Zenkner" /></a><a href="https://github.com/RVxLab"><img src="https:&#x2F;&#x2F;github.com&#x2F;RVxLab.png" width="60px" alt="User avatar: " /></a><a href="https://github.com/FrederickGeek8"><img src="https:&#x2F;&#x2F;github.com&#x2F;FrederickGeek8.png" width="60px" alt="User avatar: Frederick Morlock" /></a><a href="https://github.com/ezdac"><img src="https:&#x2F;&#x2F;github.com&#x2F;ezdac.png" width="60px" alt="User avatar: Maximilian Langenfeld" /></a><a href="https://github.com/amalavet"><img src="https:&#x2F;&#x2F;github.com&#x2F;amalavet.png" width="60px" alt="User avatar: Alejandro" /></a><a href="https://github.com/lppassos"><img src="https:&#x2F;&#x2F;github.com&#x2F;lppassos.png" width="60px" alt="User avatar: " /></a><a href="https://github.com/neilcode"><img src="https:&#x2F;&#x2F;github.com&#x2F;neilcode.png" width="60px" alt="User avatar: Neil Lambert" /></a><a href="https://github.com/dhh"><img src="https:&#x2F;&#x2F;github.com&#x2F;dhh.png" width="60px" alt="User avatar: David Heinemeier Hansson" /></a><a href="https://github.com/ethanfischer"><img src="https:&#x2F;&#x2F;github.com&#x2F;ethanfischer.png" width="60px" alt="User avatar: Ethan Fischer" /></a><a href="https://github.com/poshboytl"><img src="https:&#x2F;&#x2F;github.com&#x2F;poshboytl.png" width="60px" alt="User avatar: Terry Tai" /></a><a href="https://github.com/roesnera"><img src="https:&#x2F;&#x2F;github.com&#x2F;roesnera.png" width="60px" alt="User avatar: Adam Roesner" /></a><a href="https://github.com/seven1m"><img src="https:&#x2F;&#x2F;github.com&#x2F;seven1m.png" width="60px" alt="User avatar: Tim Morgan" /></a><a href="https://github.com/sgoridotla1"><img src="https:&#x2F;&#x2F;github.com&#x2F;sgoridotla1.png" width="60px" alt="User avatar: Maksym Shypulniak" /></a><a href="https://github.com/ADIX7"><img src="https:&#x2F;&#x2F;github.com&#x2F;ADIX7.png" width="60px" alt="User avatar: Kovács Ádám" /></a><a href="https://github.com/serranomorante"><img src="https:&#x2F;&#x2F;github.com&#x2F;serranomorante.png" width="60px" alt="User avatar: Patricio Serrano" /></a><a href="https://github.com/kiriDevs"><img src="https:&#x2F;&#x2F;github.com&#x2F;kiriDevs.png" width="60px" alt="User avatar: Kiri" /></a><a href="https://github.com/StevenMasini"><img src="https:&#x2F;&#x2F;github.com&#x2F;StevenMasini.png" width="60px" alt="User avatar: Steven Masini" /></a><a href="https://github.com/bjornevik"><img src="https:&#x2F;&#x2F;github.com&#x2F;bjornevik.png" width="60px" alt="User avatar: John Even Bjørnevik" /></a><a href="https://github.com/moberst"><img src="https:&#x2F;&#x2F;github.com&#x2F;moberst.png" width="60px" alt="User avatar: Michael Oberst" /></a><a href="https://github.com/adam-e-trepanier"><img src="https:&#x2F;&#x2F;github.com&#x2F;adam-e-trepanier.png" width="60px" alt="User avatar: Adam Trepanier" /></a><a href="https://github.com/arkalon76"><img src="https:&#x2F;&#x2F;github.com&#x2F;arkalon76.png" width="60px" alt="User avatar: Kenth Fagerlund" /></a><a href="https://github.com/Djuuu"><img src="https:&#x2F;&#x2F;github.com&#x2F;Djuuu.png" width="60px" alt="User avatar: Julien Tardot" /></a><a href="https://github.com/ellord"><img src="https:&#x2F;&#x2F;github.com&#x2F;ellord.png" width="60px" alt="User avatar: Ellord Tayag" /></a><a href="https://github.com/EdgarPost"><img src="https:&#x2F;&#x2F;github.com&#x2F;EdgarPost.png" width="60px" alt="User avatar: Edgar Post-Buijs" /></a><a href="https://github.com/caillou"><img src="https:&#x2F;&#x2F;github.com&#x2F;caillou.png" width="60px" alt="User avatar: Pierre Spring" /></a><a href="https://github.com/mebezac"><img src="https:&#x2F;&#x2F;github.com&#x2F;mebezac.png" width="60px" alt="User avatar: Zac Clay" /></a><a href="https://github.com/Tom94"><img src="https:&#x2F;&#x2F;github.com&#x2F;Tom94.png" width="60px" alt="User avatar: Thomas Müller" /></a><a href="https://github.com/carlassmann"><img src="https:&#x2F;&#x2F;github.com&#x2F;carlassmann.png" width="60px" alt="User avatar: Carl Assmann" /></a><a href="https://github.com/moodyhunter"><img src="https:&#x2F;&#x2F;github.com&#x2F;moodyhunter.png" width="60px" alt="User avatar: Moody Liu" /></a><a href="https://github.com/LasseBloch"><img src="https:&#x2F;&#x2F;github.com&#x2F;LasseBloch.png" width="60px" alt="User avatar: Lasse Bloch Lauritsen" /></a><a href="https://github.com/slavshik"><img src="https:&#x2F;&#x2F;github.com&#x2F;slavshik.png" width="60px" alt="User avatar: Alexander Slavschik" /></a><a href="https://github.com/aidalgol"><img src="https:&#x2F;&#x2F;github.com&#x2F;aidalgol.png" width="60px" alt="User avatar: Aidan Gauland" /></a><a href="https://github.com/mbienkowsk"><img src="https:&#x2F;&#x2F;github.com&#x2F;mbienkowsk.png" width="60px" alt="User avatar: Maksym Bieńkowski" /></a><a href="https://github.com/joshuawootonn"><img src="https:&#x2F;&#x2F;github.com&#x2F;joshuawootonn.png" width="60px" alt="User avatar: Joshua Wootonn" /></a><a href="https://github.com/I4nJ"><img src="https:&#x2F;&#x2F;github.com&#x2F;I4nJ.png" width="60px" alt="User avatar: " /></a><a href="https://github.com/sandviklee"><img src="https:&#x2F;&#x2F;github.com&#x2F;sandviklee.png" width="60px" alt="User avatar: Simon Sandvik Lee" /></a><a href="https://github.com/glagnar"><img src="https:&#x2F;&#x2F;github.com&#x2F;glagnar.png" width="60px" alt="User avatar: Thomas Gilbert" /></a><a href="https://github.com/skrzepto"><img src="https:&#x2F;&#x2F;github.com&#x2F;skrzepto.png" width="60px" alt="User avatar: Szymon Mucha" /></a><a href="https://github.com/unnawut"><img src="https:&#x2F;&#x2F;github.com&#x2F;unnawut.png" width="60px" alt="User avatar: Unnawut Leepaisalsuwanna" /></a><a href="https://github.com/wortmanb"><img src="https:&#x2F;&#x2F;github.com&#x2F;wortmanb.png" width="60px" alt="User avatar: Bret Wortman" /></a><a href="https://github.com/simoncdna"><img src="https:&#x2F;&#x2F;github.com&#x2F;simoncdna.png" width="60px" alt="User avatar: Simon Cardona" /></a><a href="https://github.com/andre-lameirinhas"><img src="https:&#x2F;&#x2F;github.com&#x2F;andre-lameirinhas.png" width="60px" alt="User avatar: André Lameirinhas" /></a><a href="https://github.com/SVappsLAB"><img src="https:&#x2F;&#x2F;github.com&#x2F;SVappsLAB.png" width="60px" alt="User avatar: Scott Velez" /></a><a href="https://github.com/ooojustin"><img src="https:&#x2F;&#x2F;github.com&#x2F;ooojustin.png" width="60px" alt="User avatar: justin" /></a><a href="https://github.com/mayfieldiv"><img src="https:&#x2F;&#x2F;github.com&#x2F;mayfieldiv.png" width="60px" alt="User avatar: Mayfield" /></a><a href="https://github.com/somaholiday"><img src="https:&#x2F;&#x2F;github.com&#x2F;somaholiday.png" width="60px" alt="User avatar: Soma Holiday" /></a><a href="https://github.com/drew-council"><img src="https:&#x2F;&#x2F;github.com&#x2F;drew-council.png" width="60px" alt="User avatar: Drew Council" /></a><a href="https://github.com/dessalines"><img src="https:&#x2F;&#x2F;github.com&#x2F;dessalines.png" width="60px" alt="User avatar: Dessalines" /></a><a href="https://github.com/asukaminato0721"><img src="https:&#x2F;&#x2F;github.com&#x2F;asukaminato0721.png" width="60px" alt="User avatar: Asuka Minato" /></a><a href="https://github.com/felipeospina21"><img src="https:&#x2F;&#x2F;github.com&#x2F;felipeospina21.png" width="60px" alt="User avatar: Felipe Ospina" /></a><a href="https://github.com/riccardoNovaglia"><img src="https:&#x2F;&#x2F;github.com&#x2F;riccardoNovaglia.png" width="60px" alt="User avatar: Riccardo Novaglia" /></a><a href="https://github.com/rbuchberger"><img src="https:&#x2F;&#x2F;github.com&#x2F;rbuchberger.png" width="60px" alt="User avatar: Robert Buchberger" /></a><a href="https://github.com/tyamori"><img src="https:&#x2F;&#x2F;github.com&#x2F;tyamori.png" width="60px" alt="User avatar: housekeeper" /></a><a href="https://github.com/ShGKme"><img src="https:&#x2F;&#x2F;github.com&#x2F;ShGKme.png" width="60px" alt="User avatar: Grigorii K. Shartsev" /></a><a href="https://github.com/jccint"><img src="https:&#x2F;&#x2F;github.com&#x2F;jccint.png" width="60px" alt="User avatar: " /></a><a href="https://github.com/odysseyalive"><img src="https:&#x2F;&#x2F;github.com&#x2F;odysseyalive.png" width="60px" alt="User avatar: Francis Meetze" /></a><a href="https://github.com/CarlosElizalde98"><img src="https:&#x2F;&#x2F;github.com&#x2F;CarlosElizalde98.png" width="60px" alt="User avatar: Carlos E" /></a><a href="https://github.com/darrenspurgeon"><img src="https:&#x2F;&#x2F;github.com&#x2F;darrenspurgeon.png" width="60px" alt="User avatar: Darren Spurgeon" /></a><a href="https://github.com/nestor-custodio"><img src="https:&#x2F;&#x2F;github.com&#x2F;nestor-custodio.png" width="60px" alt="User avatar: Nestor Custodio" /></a><a href="https://github.com/cb341"><img src="https:&#x2F;&#x2F;github.com&#x2F;cb341.png" width="60px" alt="User avatar: Dani Bengl" /></a><a href="https://github.com/zakrywilson"><img src="https:&#x2F;&#x2F;github.com&#x2F;zakrywilson.png" width="60px" alt="User avatar: Zach" /></a><a href="https://github.com/JusteNav"><img src="https:&#x2F;&#x2F;github.com&#x2F;JusteNav.png" width="60px" alt="User avatar: " /></a><a href="https://github.com/ivanfed0t0v"><img src="https:&#x2F;&#x2F;github.com&#x2F;ivanfed0t0v.png" width="60px" alt="User avatar: Ivan Fedotov" /></a><!-- sponsors -->
</p>

## Elevator Pitch

Rant time: You've heard it before, git is _powerful_, but what good is that power when everything is so damn hard to do? Interactive rebasing requires you to edit a goddamn TODO file in your editor? _Are you kidding me?_ To stage part of a file you need to use a command line program to step through each hunk and if a hunk can't be split down any further but contains code you don't want to stage, you have to edit an arcane patch file _by hand_? _Are you KIDDING me?!_ Sometimes you get asked to stash your changes when switching branches only to realise that after you switch and unstash that there weren't even any conflicts and it would have been fine to just checkout the branch directly? _YOU HAVE GOT TO BE KIDDING ME!_

If you're a mere mortal like me and you're tired of hearing how powerful git is when in your daily life it's a powerful pain in your ass, lazygit might be for you.

## Table of contents

- [Sponsors](#sponsors)
- [Elevator Pitch](#elevator-pitch)
- [Table of contents](#table-of-contents)
- [Features](#features)
  - [Stage hunks or individual lines](#stage-hunks-or-individual-lines)
  - [Interactive Rebase](#interactive-rebase)
  - [Cherry-pick](#cherry-pick)
  - [Bisect](#bisect)
  - [Nuke the working tree](#nuke-the-working-tree)
  - [Amend an old commit](#amend-an-old-commit)
  - [Filter](#filter)
  - [Invoke a custom command](#invoke-a-custom-command)
  - [Worktrees](#worktrees)
  - [Rebase magic (custom patches)](#rebase-magic-custom-patches)
  - [Rebase from marked base commit](#rebase-from-marked-base-commit)
  - [Undo](#undo)
  - [Commit graph](#commit-graph)
  - [Compare two commits](#compare-two-commits)
- [Tutorials](#tutorials)
- [Installation](#installation)
  - [Binary Releases](#binary-releases)
  - [Dev container](#dev-container-feature)
  - [Homebrew](#homebrew)
  - [MacPorts](#macports)
  - [Void Linux](#void-linux)
  - [Scoop (Windows)](#scoop-windows)
  - [gah (Linux and Mac OS)](#gah-linux-and-mac-os)
  - [Arch Linux](#arch-linux)
  - [Fedora / Amazon Linux 2023 / CentOS Stream](#fedora--amazon-linux-2023--centos-stream)
    - [Fedora / RHEL Derivatives (Terra)](#fedora--rhel-derivatives-terra)
  - [Solus Linux](#solus-linux)
  - [Debian and Ubuntu](#debian-and-ubuntu)
  - [Funtoo Linux](#funtoo-linux)
  - [Gentoo Linux](#gentoo-linux)
  - [openSUSE](#opensuse)
  - [NixOS](#nixos)
  - [Flox](#flox)
  - [FreeBSD](#freebsd)
  - [Termux](#termux)
  - [Conda](#conda)
  - [Go](#go)
  - [Chocolatey (Windows)](#chocolatey-windows)
  - [Winget (Windows 10 1709 or later)](#winget-windows-10-1709-or-later)
  - [Manual](#manual)
- [Usage](#usage)
  - [Keybindings](#keybindings)
  - [Changing Directory On Exit](#changing-directory-on-exit)
  - [Undo/Redo](#undoredo)
- [Configuration](#configuration)
  - [Custom Diff Renderers](#custom-diff-renderers)
  - [Custom Commands](#custom-commands)
  - [Git flow support](#git-flow-support)
- [Contributing](#contributing)
  - [Debugging Locally](#debugging-locally)
- [Donate](#donate)
- [FAQ](#faq)
  - [What do the commit colors represent?](#what-do-the-commit-colors-represent)
- [Shameless Plug](#shameless-plug)
- [Alternatives](#alternatives)

Lazygit is not my fulltime job but it is a hefty part time job so if you want to support the project please consider [sponsoring me](https://github.com/sponsors/jesseduffield)

## Features

### Stage hunks or individual lines

Press `0` on a changed file to focus its diff in the main view. The selection covers a whole hunk to begin with, so `<space>` stages that hunk and moves on to the next one. When you want only part of a hunk, press `a` for line-by-line selection and `v` to select a range of lines. What you staged shows up in the pane below. Press `<tab>` to move between the two panes; `<space>` unstages in the lower one. Press `c` to commit without leaving the diff.

<!-- demo: stage_hunks_or_lines -->
<video src="https://github.com/user-attachments/assets/f02dfaca-c64b-4836-953a-9c856c00f7c2" controls></video>

### Interactive Rebase

Press `i` to start an interactive rebase. Then squash (`s`), fixup (`f`), drop (`d`), edit (`e`), move up (`ctrl+k`) or move down (`ctrl+j`) any of the TODO commits, before continuing the rebase by bringing up the rebase options menu with `m` and then selecting `continue`.

You can also perform any of these actions as a once-off (e.g. pressing `s` on a commit to squash it) without explicitly starting a rebase.

This demo also uses shift+down to select a range of commits to move and fixup.

<!-- demo: interactive_rebase -->
<video src="https://github.com/user-attachments/assets/fc546a17-7bde-41d4-b58d-2ec95134d8ae" controls></video>

### Cherry-pick

Press `shift+c` on a commit to copy it and press `shift+v` to paste (cherry-pick) it.

<!-- demo: cherry_pick -->
<video src="https://github.com/user-attachments/assets/740d3af2-e548-47ad-8480-6b117dca939e" controls></video>

### Bisect

Press `b` in the commits view to mark a commit as good/bad in order to begin a git bisect.

<!-- demo: bisect -->
<video src="https://github.com/user-attachments/assets/b05f3d63-7c78-43ef-b4ed-dc335065cdef" controls></video>

### Nuke the working tree

For when you really want to just get rid of anything that shows up when you run `git status` (and yes that includes dirty submodules) [kidpix style](https://www.youtube.com/watch?v=N4E2B_k2Bss), press `shift+d` to bring up the reset options menu and then select the 'nuke' option.

<!-- demo: nuke_working_tree -->
<video src="https://github.com/user-attachments/assets/81f3d396-54f7-4b92-92d7-4711f56b9c17" controls></video>

### Amend an old commit

Pressing `shift+a` on any commit will amend that commit with the currently staged changes (running an interactive rebase in the background).

<!-- demo: amend_old_commit -->
<video src="https://github.com/user-attachments/assets/41e34cac-1698-446e-bae8-9eb794b33b3d" controls></video>

### Filter

You can filter a view with `/`. Here we filter down our branches view and then hit `enter` to view its commits.

<!-- demo: filter -->
<video src="https://github.com/user-attachments/assets/94bf28e8-2309-45b6-9e3d-de7f43200d22" controls></video>

### Invoke a custom command

Lazygit has a very flexible [custom command system](docs/Custom_Command_Keybindings.md). In this example a custom command is defined which emulates the built-in branch checkout action.

<!-- demo: custom_command -->
<video src="https://github.com/user-attachments/assets/f7c68622-2c58-43b9-bd47-a1117ff11fd3" controls></video>

### Worktrees

You can create worktrees to have multiple branches going at once without the need for stashing or creating WIP commits when switching between them. Press `w` in the branches view to create a worktree from the selected branch and switch to it.

<!-- demo: worktree_create_from_branches -->
<video src="https://github.com/user-attachments/assets/a8cafbb6-9fc0-45bb-860e-cbd88f8b65bb" controls></video>

### Rebase magic (custom patches)

You can build a custom patch from an old commit and then remove the patch from the commit, split out a new commit, apply the patch in reverse to the index, and more.

In this example an old commit contains a change that belongs in a commit of its own. We hit `0` on the commit to focus its diff. `<space>` on the hunk we want to move adds it to the custom patch, and `ctrl+p` opens the custom patch options, where we choose to move the patch into a new commit.

Learn more in the [Rebase magic Youtube tutorial](https://youtu.be/4XaToVut_hs).

<!-- demo: custom_patch -->
<video src="https://github.com/user-attachments/assets/9d36b387-56e6-4c72-9c44-83fcf5ad6228" controls></video>

If you only want to remove a hunk from an old commit, you don't need a custom patch for that. Select the hunk in the commit's diff and press `d`. Lazygit rewrites the commit without it, running an interactive rebase in the background.

<!-- demo: remove_hunk_from_commit -->
<video src="https://github.com/user-attachments/assets/86d0b3ec-126d-4c80-a9f2-d2419ab1e578" controls></video>

### Rebase from marked base commit

Say you're on a feature branch that was itself branched off of the develop branch, and you've decided you'd rather be branching off the master branch. You need a way to rebase only the commits from your feature branch. In this demo we check to see which was the last commit on the develop branch, then press `shift+b` to mark that commit as our base commit, then press `r` on the master branch to rebase onto it, only bringing across the commits from our feature branch. Then we push our changes with `shift+p`.

<!-- demo: rebase_onto -->
<video src="https://github.com/user-attachments/assets/7cd0c1d2-d737-4ef9-934e-54a01bd7d270" controls></video>

### Undo

You can undo the last action by pressing `z` and redo with `shift+z`. Here we drop a couple of commits and then undo the actions.
Undo uses the reflog which is specific to commits and branches so we can't undo changes to the working tree or stash.

[More info](/docs/Undoing.md)

<!-- demo: undo -->
<video src="https://github.com/user-attachments/assets/8a870289-2ab7-4a7e-8152-a04fa52310c1" controls></video>

### Commit graph

When viewing the commit graph in an enlarged window (use `+` and `_` to cycle screen modes), the commit graph is shown. Colours correspond to the commit authors, and as you navigate down the graph, the parent commits of the selected commit are highlighted.

<!-- demo: commit_graph -->
<video src="https://github.com/user-attachments/assets/83b8103c-dafc-4540-acd2-0fd9f8227aff" controls></video>

### Compare two commits

If you press `shift+w` on a commit (or branch/ref) a menu will open that allows you to mark that commit so that any other commit you select will be diffed against it. Once you've selected the second commit, you'll see the diff in the main view and if you press `<enter>` you'll see the files of the diff. You can press `shift+w` to view the diff menu again to see options like reversing the diff direction or exiting diff mode. You can also exit diff mode by pressing `<escape>`.

<!-- demo: diff_commits -->
<video src="https://github.com/user-attachments/assets/7094d9dd-d167-4caf-b79d-6da35a3f0279" controls></video>

### Show GitHub pull requests

In the branches panel, lazygit can show which of your branches have an associated GitHub pull request by showing a GitHub icon next to the branch name; its color shows the state of the PR (open, merged, etc.). For those that have one, you can press `shift-G` to open the PR in the browser. There is no configuration needed to enable this for github.com, but it requires the [`gh`](https://cli.github.com/) tool to be installed, and you need to do `gh auth login` once to allow lazygit to access GitHub. For GitHub Enterprise, also run `gh auth login --hostname <webDomain>` and add a [`services` entry](docs/Config.md#custom-pull-request-urls) for the host with the `github` provider.

## Tutorials

[<img src="https://i.imgur.com/sVEktDn.png">](https://youtu.be/CPLdltN7wgE)

- [15 Lazygit Features in 15 Minutes](https://youtu.be/CPLdltN7wgE)
- [Basics Tutorial](https://youtu.be/VDXvbHZYeKY)
- [Rebase Magic Tutorial](https://youtu.be/4XaToVut_hs)

## Installation

[![Packaging status](https://repology.org/badge/vertical-allrepos/lazygit.svg?columns=3)](https://repology.org/project/lazygit/versions)

_Most of the above packages are maintained by third parties so be sure to vet them yourself and confirm that the maintainer is a trustworthy looking person who attends local sports games and gives back to their communities with barbeque fundraisers etc_

### Binary Releases

For Windows, Mac OS(10.12+) or Linux, you can download a binary release [here](../../releases).

### Dev container feature

If you want to use lazygit in e.g. one of your GitHub Codespaces, there is a third-party [dev container feature](https://github.com/GeorgOfenbeck/features/tree/main/src/lazygit-linuxbinary) based on the binary releases mentioned above.

### Homebrew

It works with Linux, too.

```sh
brew install lazygit
```

### MacPorts

Latest version built from github releases.
Tap:

```
sudo port install lazygit
```

### Void Linux

Packages for Void Linux are available in the distro repo

They follow upstream latest releases

```sh
sudo xbps-install -S lazygit
```

### Scoop (Windows)

You can install `lazygit` using [scoop](https://scoop.sh/). It's in the `extras` bucket:

```sh
# Add the extras bucket
scoop bucket add extras

# Install lazygit
scoop install lazygit
```

### gah (Linux and Mac OS)

You can install `lazygit` using [gah](https://github.com/marverix/gah/):

```sh
gah install lazygit
```

### Arch Linux

Packages for Arch Linux are available via pacman and AUR (Arch User Repository).

There are two packages. The stable one which is built with the latest release
and the git version which builds from the most recent commit.

- Stable: `sudo pacman -S lazygit`
- Development: <https://aur.archlinux.org/packages/lazygit-git/>

Instruction of how to install AUR content can be found here:
<https://wiki.archlinux.org/index.php/Arch_User_Repository>

### Fedora / Amazon Linux 2023 / CentOS Stream

Packages for Fedora, Amazon Linux 2023 and CentOS Stream are available via
[Copr](https://copr.fedorainfracloud.org/coprs/dejan/lazygit/) (Cool Other Package Repo).

```sh
sudo dnf copr enable dejan/lazygit
sudo dnf install lazygit
```

These packages are built using the RPM spec file located here: https://codeberg.org/dejan/rpm-lazygit

You should be able to build RPMs for Fedora 41 or older, and other Fedora derivatives using the
SRPM (Source RPM) file that you can grab from the latest COPR build.

#### Fedora / RHEL Derivatives (Terra)

Packages for Fedora and RHEL derivatives are also available from the [Terra Repository](https://terra.fyralabs.com/).

```sh
sudo dnf install --nogpgcheck --repofrompath 'terra,https://repos.fyralabs.com/terra$releasever' terra-release
sudo dnf install lazygit
```

(Install guide for Atomic/EL can be found on their [README](https://github.com/terrapkg/packages/pull/9747))

Terra also has `lazygit-doc`, which contains the contents of the [docs](https://github.com/jesseduffield/lazygit/tree/master/docs) folder.

### Solus Linux

```sh
sudo eopkg install lazygit
```

### Debian and Ubuntu

For **Debian 13 "Trixie", Sid**, and later, or **Ubuntu 25.10 "Questing Quokka"** and later:

```sh
sudo apt install lazygit
```

For **Debian 12 "Bookworm", Ubuntu 25.04 "Plucky Puffin"** and earlier:

```sh
LAZYGIT_VERSION=$(curl -s "https://api.github.com/repos/jesseduffield/lazygit/releases/latest" | \grep -Po '"tag_name": *"v\K[^"]*')
LAZYGIT_ARCH=$(uname -m | sed -e 's/aarch64/arm64/')
curl -Lo lazygit.tar.gz "https://github.com/jesseduffield/lazygit/releases/download/v${LAZYGIT_VERSION}/lazygit_${LAZYGIT_VERSION}_Linux_${LAZYGIT_ARCH}.tar.gz"
tar xf lazygit.tar.gz lazygit
sudo install lazygit -D -t /usr/local/bin/
```

Verify the correct installation of lazygit:

```sh
lazygit --version
```

### Funtoo Linux

Funtoo Linux has an autogenerated lazygit package in [dev-kit](https://github.com/funtoo/dev-kit/tree/1.4-release/dev-vcs/lazygit):

```sh
sudo emerge dev-vcs/lazygit
```

### Gentoo Linux

Lazygit is not (yet) in main Gentoo portage, however an ebuild is available in [GURU overlay](https://github.com/gentoo-mirror/guru/tree/master/dev-vcs/lazygit)

You can either add the overlay to your system and install lazygit as usual:

```sh
sudo eselect repository enable guru
sudo emaint sync -r guru
sudo emerge dev-vcs/lazygit
```

### openSUSE

The lazygit package is currently built in [devel:languages:go/lazygit](https://build.opensuse.org/package/show/devel:languages:go/lazygit).

To install lazygit on openSUSE Tumbleweed run:

```sh
sudo zypper ar https://download.opensuse.org/repositories/devel:/languages:/go/openSUSE_Factory/devel:languages:go.repo
sudo zypper ref && sudo zypper in lazygit
```

To install lazygit on openSUSE Leap run:

```sh
source /etc/os-release
sudo zypper ar https://download.opensuse.org/repositories/devel:/languages:/go/$VERSION_ID/devel:languages:go.repo
sudo zypper ref && sudo zypper in lazygit
```

### NixOS

#### Using lazygit from nixpkgs

On NixOS, lazygit is packaged with nix and distributed via nixpkgs.
You can try lazygit without installing it with:

```sh
nix-shell -p lazygit
# or with flakes enabled
nix run nixpkgs#lazygit
```

Or you can add lazygit to your `configuration.nix` using the `environment.systemPackages` option.
More details can be found via NixOS search [page](https://search.nixos.org/).

#### Using the official lazygit flake

This repository includes a nix flake that provides the latest development version and additional development tools:

**Run lazygit directly from the repository:**

```sh
nix run github:jesseduffield/lazygit
# or from a local clone
nix run .
```

**Build lazygit from source:**

```sh
nix build github:jesseduffield/lazygit
# or from a local clone
nix build .
```

**Development environment:**
For contributors, the flake provides a development shell with Go toolchain, development tools, and dependencies:

```sh
nix develop github:jesseduffield/lazygit
# or from a local clone
nix develop
```

The development shell includes:

- Go toolchain
- git and make
- Proper environment variables for development

**Using in other flakes:**
The flake also provides an overlay for easy integration into other flake-based projects:

```nix
{
  inputs.lazygit.url = "github:jesseduffield/lazygit";

  outputs = { self, nixpkgs, lazygit }: {
    # Use the overlay
    nixpkgs.overlays = [ lazygit.overlays.default ];
  };
}
```

### Flox

Lazygit can be installed into a Flox environment as follows.

```sh
flox install lazygit
```

More details about Flox can be found on [their website](https://flox.dev/).

### FreeBSD

```sh
pkg install lazygit
```

### Termux

```sh
apt install lazygit
```

### Conda

Released versions are available for different platforms, see <https://anaconda.org/conda-forge/lazygit>

```sh
conda install -c conda-forge lazygit
```

### Go

```sh
go install github.com/jesseduffield/lazygit@latest
```

Please note:
If you get an error claiming that lazygit cannot be found or is not defined, you
may need to add `~/go/bin` to your $PATH (macOS/Linux), or `%HOME%\go\bin`
(Windows). Not to be mistaken for `C:\Go\bin` (which is for Go's own binaries,
not apps like lazygit).

### Chocolatey (Windows)

You can install `lazygit` using [Chocolatey](https://chocolatey.org/):

```sh
choco install lazygit
```

### Winget (Windows 10 1709 or later)

You can install `lazygit` using the `winget` command in the Windows Terminal with the following command:

```powershell
winget install -e --id=JesseDuffield.lazygit
```

### Manual

You'll need to [install Go](https://golang.org/doc/install)

```
git clone https://github.com/jesseduffield/lazygit.git
cd lazygit
go install
```

You can also use `go run main.go` to compile and run in one go (pun definitely intended)

## Usage

Call `lazygit` in your terminal inside a git repository.

```sh
$ lazygit
```

If you want, you can
also add an alias for this with `echo "alias lg='lazygit'" >> ~/.zshrc` (or
whichever rc file you're using).

### Keybindings

You can check out the list of keybindings [here](/docs/keybindings).

### Changing Directory On Exit

If you change repos in lazygit and want your shell to change directory into that repo on exiting lazygit, add this to your `~/.zshrc` (or other rc file):

```
lg()
{
    export LAZYGIT_NEW_DIR_FILE=~/.lazygit/newdir

    lazygit "$@"

    if [ -f $LAZYGIT_NEW_DIR_FILE ]; then
            cd "$(cat $LAZYGIT_NEW_DIR_FILE)"
            rm -f $LAZYGIT_NEW_DIR_FILE > /dev/null
    fi
}
```

Then `source ~/.zshrc` and from now on when you call `lg` and exit you'll switch directories to whatever you were in inside lazygit. To override this behaviour you can exit using `shift+Q` rather than just `q`.

### Undo/Redo

See the [docs](/docs/Undoing.md)

## Configuration

Check out the [configuration docs](docs/Config.md).

### Custom Diff Renderers

See the [docs](docs/Custom_DiffRenderers.md)

### Custom Commands

If lazygit is missing a feature, there's a good chance you can implement it yourself with a custom command!

See the [docs](docs/Custom_Command_Keybindings.md)

### Git flow support

Lazygit supports [Gitflow](https://github.com/nvie/gitflow) (or [git-flow-next](https://github.com/gittower/git-flow-next)) if you have it installed. To understand how the Gitflow model works check out Vincent Driessen's original [post](https://nvie.com/posts/a-successful-git-branching-model/) explaining it. To view Gitflow options from within Lazygit, press `i` from within the branches view.

## Contributing

We love your input! Please check out the [contributing guide](CONTRIBUTING.md).
For contributor discussion about things not better discussed here in the repo, join the [discord channel](https://discord.gg/ehwFt2t4wt)

<a href="https://discord.gg/ehwFt2t4wt"><img src='../assets/discord.png' width='75'></a>

Check out this [video](https://www.youtube.com/watch?v=kNavnhzZHtk) walking through the creation of a small feature in lazygit if you want an idea of where to get started.

### Debugging Locally

Run `lazygit --debug` in one terminal tab and `lazygit --logs` in another to view the program and its log output side by side

## Donate

If you would like to support the development of lazygit, consider [sponsoring me](https://github.com/sponsors/jesseduffield) (github is matching all donations dollar-for-dollar for 12 months)

## FAQ

### What do the commit colors represent?

- Green: the commit is included in the master branch
- Yellow: the commit is not included in the master branch
- Red: the commit has not been pushed to the upstream branch

## Shameless Plug

If you want to see what I (Jesse) am up to in terms of development, follow me on
[twitter](https://twitter.com/DuffieldJesse) or check out my [blog](https://jesseduffield.com/)

## Alternatives

If you find that lazygit doesn't quite satisfy your requirements, these may be a better fit:

- [GitUI](https://github.com/Extrawurst/gitui)
- [tig](https://github.com/jonas/tig)
- [GitArbor TUI](https://github.com/cadamsdev/gitarbor-tui)
