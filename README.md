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

Before (baseline): `8e3ac76991f5` · After (measured revision): `818a17b1d21a`.

Same benchmark sources and settings. Median of 5 samples, 200ms per sample, GOMAXPROCS=1, GOGC=100, GOMEMLIMIT=off.

`go version go1.27.3 linux/amd64` · 13th Gen Intel(R) Core(TM) i7-1370P.

Each pair of bars uses its own scale. Negative change means faster; positive means slower.

[Raw samples and fixture metadata](docs-master/dev/performance-comparison.json) · [Investigation notes](docs-master/dev/Performance-follow-up.md).

<table>
<thead><tr><th align="left">Benchmark / workload</th><th align="left">Before / after and change</th></tr></thead>
<tbody>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkDrawSearchedView/lines=100">Draw searched view</strong><br>
Redraw an 80 × 24 searched viewport within 100 lines.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkDrawSearchedView-lines-100-5cafa93dd7c0.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkDrawSearchedView/lines=100: 858.2 µs before, 91.17 µs after, -89.38% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkDrawSearchedView/lines=1000">Draw searched view</strong><br>
Redraw an 80 × 24 searched viewport within 1000 lines.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkDrawSearchedView-lines-1000-6791c36508a9.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkDrawSearchedView/lines=1000: 6.955 ms before, 97.02 µs after, -98.61% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkDrawSearchedView/lines=10000">Draw searched view</strong><br>
Redraw an 80 × 24 searched viewport within 10000 lines.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkDrawSearchedView-lines-10000-6a2559ca5acb.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkDrawSearchedView/lines=10000: 73.78 ms before, 94.5 µs after, -99.87% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorEditUnbrokenText/bytes=4096">Editor edit unbroken text</strong><br>
Insert and backspace one character in 4096 bytes of existing unbroken text; update layout and view after each key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorEditUnbrokenText-bytes-4096-feb9a2c3e5b9.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorEditUnbrokenText/bytes=4096: 1.761 ms before, 18.62 µs after, -98.94% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorEditUnbrokenText/bytes=100000">Editor edit unbroken text</strong><br>
Insert and backspace one character in 100000 bytes of existing unbroken text; update layout and view after each key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorEditUnbrokenText-bytes-100000-eda65efcb6a7.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorEditUnbrokenText/bytes=100000: 77.18 ms before, 29.64 µs after, -99.96% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorEditUnbrokenText/bytes=1000000">Editor edit unbroken text</strong><br>
Insert and backspace one character in 1000000 bytes of existing unbroken text; update layout and view after each key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorEditUnbrokenText-bytes-1000000-03d9b331bd31.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorEditUnbrokenText/bytes=1000000: 622.5 ms before, 289.5 µs after, -99.95% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorInput/bytes=256">Editor input</strong><br>
Type 250 bytes with hard newlines, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorInput-bytes-256-09b9a90cdf09.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorInput/bytes=256: 5.645 ms before, 256.3 µs after, -95.46% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorInput/bytes=1024">Editor input</strong><br>
Type 1020 bytes with hard newlines, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorInput-bytes-1024-577aca8d08c8.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorInput/bytes=1024: 97.28 ms before, 1.627 ms after, -98.33% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorInput/bytes=4096">Editor input</strong><br>
Type 4090 bytes with hard newlines, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorInput-bytes-4096-33e86a7d08e9.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorInput/bytes=4096: 1.811 s before, 7.103 ms after, -99.61% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorParagraphInput/bytes=256">Editor paragraph input</strong><br>
Type 250 bytes of a spaced paragraph, wrapping at 72 columns.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorParagraphInput-bytes-256-17ada10d6ff3.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorParagraphInput/bytes=256: 5.093 ms before, 683.6 µs after, -86.58% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorParagraphInput/bytes=1024">Editor paragraph input</strong><br>
Type 1020 bytes of a spaced paragraph, wrapping at 72 columns.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorParagraphInput-bytes-1024-c5152d2cff38.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorParagraphInput/bytes=1024: 86.66 ms before, 3.149 ms after, -96.37% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorParagraphInput/bytes=4096">Editor paragraph input</strong><br>
Type 4090 bytes of a spaced paragraph, wrapping at 72 columns.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorParagraphInput-bytes-4096-3acf8f82a5b0.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorParagraphInput/bytes=4096: 1.822 s before, 12.96 ms after, -99.29% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorUnbrokenInput/bytes=256">Editor unbroken input</strong><br>
Type 256 unbroken ASCII characters, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorUnbrokenInput-bytes-256-b70d50913a0b.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorUnbrokenInput/bytes=256: 6.654 ms before, 881.5 µs after, -86.75% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorUnbrokenInput/bytes=1024">Editor unbroken input</strong><br>
Type 1024 unbroken ASCII characters, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorUnbrokenInput-bytes-1024-57f46c5d5b92.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorUnbrokenInput/bytes=1024: 101.3 ms before, 4.314 ms after, -95.74% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkEditorUnbrokenInput/bytes=4096">Editor unbroken input</strong><br>
Type 4096 unbroken ASCII characters, updating layout and view per key.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkEditorUnbrokenInput-bytes-4096-e97edfb1951f.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkEditorUnbrokenInput/bytes=4096: 1.721 s before, 21.3 ms after, -98.76% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkGuiFlush/80x24">Gui flush</strong><br>
Flush two panes through a mock terminal at 80x24, changing selection and scroll position.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkGuiFlush-80x24-3e1eea137cef.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkGuiFlush/80x24: 447.3 µs before, 202.9 µs after, -54.64% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkGuiFlush/160x48">Gui flush</strong><br>
Flush two panes through a mock terminal at 160x48, changing selection and scroll position.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkGuiFlush-160x48-14ecf55d6d47.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkGuiFlush/160x48: 1.509 ms before, 544 µs after, -63.94% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkNearestSearchPosition/matches=100">Nearest search position</strong><br>
Find the nearest result near the end of 100 search matches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkNearestSearchPosition-matches-100-4c321f724001.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkNearestSearchPosition/matches=100: 75.42 ns before, 3.971 ns after, -94.73% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkNearestSearchPosition/matches=1000">Nearest search position</strong><br>
Find the nearest result near the end of 1000 search matches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkNearestSearchPosition-matches-1000-1530f7777617.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkNearestSearchPosition/matches=1000: 735.9 ns before, 5.522 ns after, -99.25% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkNearestSearchPosition/matches=10000">Nearest search position</strong><br>
Find the nearest result near the end of 10000 search matches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkNearestSearchPosition-matches-10000-a853c867a9d2.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkNearestSearchPosition/matches=10000: 6.794 µs before, 7.962 ns after, -99.88% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkNearestSearchPosition/matches=100000">Nearest search position</strong><br>
Find the nearest result near the end of 100000 search matches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkNearestSearchPosition-matches-100000-8170d4b19fdc.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkNearestSearchPosition/matches=100000: 76.62 µs before, 10.19 ns after, -99.99% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkPasteText/bytes=1000">Paste text</strong><br>
Insert 986 bytes into prefilled text with TypeString, including layout.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkPasteText-bytes-1000-176f072ac98d.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkPasteText/bytes=1000: 240.7 µs before, 29.83 µs after, -87.60% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkPasteText/bytes=10000">Paste text</strong><br>
Insert 9996 bytes into prefilled text with TypeString, including layout.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkPasteText-bytes-10000-867dc80de1df.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkPasteText/bytes=10000: 14.06 ms before, 302.4 µs after, -97.85% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkPasteText/bytes=100000">Paste text</strong><br>
Insert 99994 bytes into prefilled text with TypeString, including layout.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkPasteText-bytes-100000-e9e57ef8def6.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkPasteText/bytes=100000: 686.6 ms before, 2.658 ms after, -99.61% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=100/up=false">Scroll search position</strong><br>
Select a visible result after a distant scroll down across 100 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-100-up-false-a860fde48d0f.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=100/up=false: 102 ns before, 15.97 ns after, -84.34% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=100/up=true">Scroll search position</strong><br>
Select a visible result after a distant scroll up across 100 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-100-up-true-0553b419f53a.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=100/up=true: 75.42 ns before, 15.5 ns after, -79.45% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=1000/up=false">Scroll search position</strong><br>
Select a visible result after a distant scroll down across 1000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-1000-up-false-5bccc6ef4de6.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=1000/up=false: 776.9 ns before, 17.95 ns after, -97.69% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=1000/up=true">Scroll search position</strong><br>
Select a visible result after a distant scroll up across 1000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-1000-up-true-f3521677c3ff.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=1000/up=true: 566.1 ns before, 17.08 ns after, -96.98% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=10000/up=false">Scroll search position</strong><br>
Select a visible result after a distant scroll down across 10000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-10000-up-false-2974a828ee98.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=10000/up=false: 6.739 µs before, 20.51 ns after, -99.70% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=10000/up=true">Scroll search position</strong><br>
Select a visible result after a distant scroll up across 10000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-10000-up-true-753585314e0c.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=10000/up=true: 4.979 µs before, 18.8 ns after, -99.62% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=100000/up=false">Scroll search position</strong><br>
Select a visible result after a distant scroll down across 100000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-100000-up-false-ae6dfb1321e7.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=100000/up=false: 73.44 µs before, 22.94 ns after, -99.97% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkScrollSearchPosition/matches=100000/up=true">Scroll search position</strong><br>
Select a visible result after a distant scroll up across 100000 matches; reset each iteration.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkScrollSearchPosition-matches-100000-up-true-5678ba67694f.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkScrollSearchPosition/matches=100000/up=true: 60.47 µs before, 21.98 ns after, -99.96% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=a">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;a&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-mixed-case-query-a-5b21d07411f5.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=a: 71.56 ms before, 21.21 ms after, -70.36% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=MATCH">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;MATCH&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-mixed-case-query-MATCH-8f92520e755b.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=MATCH: 27.04 ms before, 15.17 ms after, -43.91% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=missing">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;missing&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-mixed-case-query-missing-fd45d3b80d12.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=missing: 47.73 ms before, 20.24 ms after, -57.59% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=performance-match">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;performance-match&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-mixed-case-query-performance-match-8d2b9b612b93.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=mixed-case/query=performance-match: 71.42 ms before, 26.29 ms after, -63.19% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=a">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;a&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-ordinary-query-a-83ea329205a5.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=a: 56.75 ms before, 30.69 ms after, -45.92% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=MATCH">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;MATCH&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-ordinary-query-MATCH-e6145c953605.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=MATCH: 31.32 ms before, 20.35 ms after, -35.05% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=missing">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;missing&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-ordinary-query-missing-a9278e828d67.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=missing: 40.54 ms before, 26.13 ms after, -35.55% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=performance-match">Search large diff</strong><br>
Search 50000 diff lines with wide and combining characters for &quot;performance-match&quot;; rendered layout and initial query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchLargeDiff-lines-50000-text-ordinary-query-performance-match-959d84b5d3e0.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchLargeDiff/lines=50000/text=ordinary/query=performance-match: 58.76 ms before, 31.53 ms after, -46.34% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchOrdinaryText/match">Search ordinary text</strong><br>
Search 2,000 ordinary diff lines for &quot;match&quot;; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchOrdinaryText-match-e717e6d8da02.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchOrdinaryText/match: 1.305 ms before, 418.2 µs after, -67.96% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchOrdinaryText/MATCH">Search ordinary text</strong><br>
Search 2,000 ordinary diff lines for &quot;MATCH&quot;; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchOrdinaryText-MATCH-bd0b348d84bd.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchOrdinaryText/MATCH: 825.3 µs before, 268.4 µs after, -67.48% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchOrdinaryText/missing">Search ordinary text</strong><br>
Search 2,000 ordinary diff lines for &quot;missing&quot;; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchOrdinaryText-missing-5d45b84791aa.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchOrdinaryText/missing: 1.146 ms before, 376.5 µs after, -67.14% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchQueryLifetime/clear=false">Search query lifetime</strong><br>
Alternate aa, aaa, aaaa, and a queries across 2,000 lines of 200 a characters; keep search active; rendered layout and initial a query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchQueryLifetime-clear-false-f852edaf1e6a.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchQueryLifetime/clear=false: 33.51 ms before, 2.931 ms after, -91.25% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchQueryLifetime/clear=true">Search query lifetime</strong><br>
Alternate aa, aaa, aaaa, and a queries across 2,000 lines of 200 a characters; clear search after each query; rendered layout and initial a query warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchQueryLifetime-clear-true-d41619faae14.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchQueryLifetime/clear=true: 34.96 ms before, 11.25 ms after, -67.83% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchRepeatedPrefix/query=1">Search repeated prefix</strong><br>
Search 2,000 lines of 200 a characters for a; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchRepeatedPrefix-query-1-47f1e0fa171d.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchRepeatedPrefix/query=1: 20.59 ms before, 2.467 ms after, -88.02% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchRepeatedPrefix/query=32">Search repeated prefix</strong><br>
Search 2,000 lines of 200 a characters for a 32-character query ending in b; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchRepeatedPrefix-query-32-519b7141bcb9.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchRepeatedPrefix/query=32: 76.46 ms before, 3.952 ms after, -94.83% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkSearchRepeatedPrefix/query=128">Search repeated prefix</strong><br>
Search 2,000 lines of 200 a characters for a 128-character query ending in b; rendered layout warmed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkSearchRepeatedPrefix-query-128-0fe2efba8214.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkSearchRepeatedPrefix/query=128: 221.1 ms before, 3.268 ms after, -98.52% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/80x24/empty">View draw</strong><br>
Redraw a 80x24 view with empty content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-80x24-empty-1e831f9bbbc4.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/80x24/empty: 96.33 µs before, 4.731 µs after, -95.09% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/80x24/full">View draw</strong><br>
Redraw a 80x24 view with full content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-80x24-full-7e676d236249.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/80x24/full: 181.8 µs before, 39.43 µs after, -78.31% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/80x24/short">View draw</strong><br>
Redraw a 80x24 view with short content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-80x24-short-f00153ca7e14.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/80x24/short: 84.91 µs before, 4.684 µs after, -94.48% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/80x24/unicode">View draw</strong><br>
Redraw a 80x24 view with unicode content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-80x24-unicode-8e23f05e1a94.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/80x24/unicode: 228.4 µs before, 29.99 µs after, -86.87% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/160x48/empty">View draw</strong><br>
Redraw a 160x48 view with empty content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-160x48-empty-2f015b2c6bf9.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/160x48/empty: 388.2 µs before, 16.35 µs after, -95.79% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/160x48/full">View draw</strong><br>
Redraw a 160x48 view with full content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-160x48-full-86728f2ecaa0.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/160x48/full: 824.3 µs before, 167.3 µs after, -79.70% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/160x48/short">View draw</strong><br>
Redraw a 160x48 view with short content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-160x48-short-4c132fd9bdab.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/160x48/short: 351.7 µs before, 16.22 µs after, -95.39% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewDraw/160x48/unicode">View draw</strong><br>
Redraw a 160x48 view with unicode content.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewDraw-160x48-unicode-7ed344e64b68.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewDraw/160x48/unicode: 884.1 µs before, 126.9 µs after, -85.64% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWrite/colored">View write</strong><br>
Write 2,000 colored lines into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWrite-colored-cc0c7ac6042f.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWrite/colored: 15.53 ms before, 4.481 ms after, -71.14% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWrite/metadata">View write</strong><br>
Write 2,000 metadata lines into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWrite-metadata-4b28ce99258a.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWrite/metadata: 8.356 ms before, 2.715 ms after, -67.51% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWrite/plain">View write</strong><br>
Write 2,000 plain lines into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWrite-plain-f98a34856771.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWrite/plain: 21.88 ms before, 2.957 ms after, -86.48% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWrite/unicode">View write</strong><br>
Write 2,000 unicode lines into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWrite-unicode-4810b5a5bb91.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWrite/unicode: 11.79 ms before, 4.062 ms after, -65.55% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWriteGitColoredDiff">View write git colored diff</strong><br>
Stream 50,000 mixed Unicode lines with separately colored Git prefixes into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWriteGitColoredDiff-7ab80bf2c172.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWriteGitColoredDiff: 690.6 ms before, 160.4 ms after, -76.78% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWriteLargeDiff/colored">View write large diff</strong><br>
Stream 50,000 colored diff lines into a new view, including parsing and allocation.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWriteLargeDiff-colored-947efcf322a4.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWriteLargeDiff/colored: 398.6 ms before, 128.3 ms after, -67.82% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWriteLargeDiff/metadata">View write large diff</strong><br>
Stream 50,000 metadata diff lines into a new view, including parsing and allocation.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWriteLargeDiff-metadata-91cd695ddc2e.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWriteLargeDiff/metadata: 284.3 ms before, 75.7 ms after, -73.38% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkViewWriteLongLine">View write long line</strong><br>
Write one 13,001-byte line into a new view.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkViewWriteLongLine-b5dc13b04049.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkViewWriteLongLine: 2.927 ms before, 221.2 µs after, -92.44% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gocui/BenchmarkWriteToSearchedView">Write to searched view</strong><br>
Write 2,000 lines into a searched view and read search status.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gocui-BenchmarkWriteToSearchedView-a7a4044fb1de.svg" width="420" height="60" align="middle" alt="gocui/BenchmarkWriteToSearchedView: 25.31 ms before, 3.828 ms after, -84.88% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/BenchmarkCommitDescriptionInput/bytes=256">Commit description input</strong><br>
Type 250 bytes through the commit-description editor.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-BenchmarkCommitDescriptionInput-bytes-256-8ddc3c26c312.svg" width="420" height="60" align="middle" alt="gui/BenchmarkCommitDescriptionInput/bytes=256: 10.38 ms before, 288.7 µs after, -97.22% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/BenchmarkCommitDescriptionInput/bytes=1024">Commit description input</strong><br>
Type 1020 bytes through the commit-description editor.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-BenchmarkCommitDescriptionInput-bytes-1024-a67a117ae643.svg" width="420" height="60" align="middle" alt="gui/BenchmarkCommitDescriptionInput/bytes=1024: 169.9 ms before, 1.344 ms after, -99.21% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/BenchmarkCommitDescriptionInput/bytes=4096">Commit description input</strong><br>
Type 4090 bytes through the commit-description editor.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-BenchmarkCommitDescriptionInput-bytes-4096-1414431d6a5e.svg" width="420" height="60" align="middle" alt="gui/BenchmarkCommitDescriptionInput/bytes=4096: 4.811 s before, 6.764 ms after, -99.86% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/context/BenchmarkFilteredListNavigation/items=100">Filtered list navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count 100 filtered branches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-context-BenchmarkFilteredListNavigation-items-100-015295c5292d.svg" width="420" height="60" align="middle" alt="gui/context/BenchmarkFilteredListNavigation/items=100: 2.652 µs before, 39.62 ns after, -98.51% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/context/BenchmarkFilteredListNavigation/items=1000">Filtered list navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count 1000 filtered branches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-context-BenchmarkFilteredListNavigation-items-1000-cccc9aaa8bb6.svg" width="420" height="60" align="middle" alt="gui/context/BenchmarkFilteredListNavigation/items=1000: 24.99 µs before, 40.28 ns after, -99.84% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/context/BenchmarkFilteredListNavigation/items=10000">Filtered list navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count 10000 filtered branches.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-context-BenchmarkFilteredListNavigation-items-10000-9361c893fe52.svg" width="420" height="60" align="middle" alt="gui/context/BenchmarkFilteredListNavigation/items=10000: 277.2 µs before, 40.75 ns after, -99.99% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildCommitFileTree/files=100">Build commit file tree</strong><br>
Build and sort a tree of 100 commit files supplied in reverse order.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildCommitFileTree-files-100-1375c947d4f4.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildCommitFileTree/files=100: 59.15 µs before, 55.81 µs after, -5.65% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildCommitFileTree/files=1000">Build commit file tree</strong><br>
Build and sort a tree of 1000 commit files supplied in reverse order.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildCommitFileTree-files-1000-5fbfa0fb7218.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildCommitFileTree/files=1000: 2.077 ms before, 496.2 µs after, -76.11% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkBuildCommitFileTree/files=10000">Build commit file tree</strong><br>
Build and sort a tree of 10000 commit files supplied in reverse order.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkBuildCommitFileTree-files-10000-4fef09af0a3c.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkBuildCommitFileTree/files=10000: 204.8 ms before, 6.904 ms after, -96.63% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeNavigation/files=100">File tree navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count rows in a 100-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeNavigation-files-100-0a982211298c.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeNavigation/files=100: 2.306 µs before, 45.17 ns after, -98.04% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeNavigation/files=1000">File tree navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count rows in a 1000-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeNavigation-files-1000-91e7264a3dbb.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeNavigation/files=1000: 21.52 µs before, 47.07 ns after, -99.78% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeNavigation/files=10000">File tree navigation</strong><br>
Reset selection to the midpoint, move one row, retrieve it, and count rows in a 10000-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeNavigation-files-10000-242f69f4477b.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeNavigation/files=10000: 236.3 µs before, 43.61 ns after, -99.98% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeRangeSelection/files=100">File tree range selection</strong><br>
Retrieve a range spanning a 100-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeRangeSelection-files-100-4c3110e7ee44.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeRangeSelection/files=100: 31.19 µs before, 2.872 µs after, -90.79% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeRangeSelection/files=1000">File tree range selection</strong><br>
Retrieve a range spanning a 1000-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeRangeSelection-files-1000-25a9869ad39e.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeRangeSelection/files=1000: 2.602 ms before, 28.73 µs after, -98.90% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkFileTreeRangeSelection/files=10000">File tree range selection</strong><br>
Retrieve a range spanning a 10000-file tree.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkFileTreeRangeSelection-files-10000-333daad69bc5.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkFileTreeRangeSelection/files=10000: 289.7 ms before, 322 µs after, -99.89% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=100">Refresh after directory removal</strong><br>
Refresh a tree after removing a 100-file directory; fixture setup excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRefreshAfterDirectoryRemoval-files-100-33bafcee01d2.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=100: 1.656 ms before, 94.87 µs after, -94.27% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=1000">Refresh after directory removal</strong><br>
Refresh a tree after removing a 1000-file directory; fixture setup excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRefreshAfterDirectoryRemoval-files-1000-5f1ed92f00ff.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=1000: 181.6 ms before, 919.9 µs after, -99.49% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=5000">Refresh after directory removal</strong><br>
Refresh a tree after removing a 5000-file directory; fixture setup excluded.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRefreshAfterDirectoryRemoval-files-5000-a8effa23ea70.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRefreshAfterDirectoryRemoval/files=5000: 4.743 s before, 5.321 ms after, -99.89% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=100">Restore after single removal</strong><br>
Restore selection after removing one file from 100 remaining files.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreAfterSingleRemoval-files-100-11d3c7368b4e.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=100: 13.3 µs before, 3.353 µs after, -74.78% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=1000">Restore after single removal</strong><br>
Restore selection after removing one file from 1000 remaining files.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreAfterSingleRemoval-files-1000-a2d2a42a0a97.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=1000: 130.9 µs before, 33.62 µs after, -74.31% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=5000">Restore after single removal</strong><br>
Restore selection after removing one file from 5000 remaining files.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreAfterSingleRemoval-files-5000-b39cad51c11c.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreAfterSingleRemoval/files=5000: 778.5 µs before, 184.7 µs after, -76.28% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=100/removed=false">Restore file selection</strong><br>
Restore selection among 100 files; selected file retained.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-100-removed-false-0a764fa3c327.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=100/removed=false: 247.5 ns before, 62.15 ns after, -74.89% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=100/removed=true">Restore file selection</strong><br>
Restore selection among 100 files; selected file removed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-100-removed-true-40a695a8e3c2.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=100/removed=true: 1.321 ms before, 20.13 µs after, -98.48% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=1000/removed=false">Restore file selection</strong><br>
Restore selection among 1000 files; selected file retained.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-1000-removed-false-cfc5395738be.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=1000/removed=false: 276.5 ns before, 64.59 ns after, -76.64% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=1000/removed=true">Restore file selection</strong><br>
Restore selection among 1000 files; selected file removed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-1000-removed-true-293fa4cf5641.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=1000/removed=true: 137 ms before, 188.4 µs after, -99.86% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=5000/removed=false">Restore file selection</strong><br>
Restore selection among 5000 files; selected file retained.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-5000-removed-false-faf8b5c1d40f.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=5000/removed=false: 312.8 ns before, 76.42 ns after, -75.57% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/filetree/BenchmarkRestoreFileSelection/files=5000/removed=true">Restore file selection</strong><br>
Restore selection among 5000 files; selected file removed.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-filetree-BenchmarkRestoreFileSelection-files-5000-removed-true-558842f2c4c3.svg" width="420" height="60" align="middle" alt="gui/filetree/BenchmarkRestoreFileSelection/files=5000/removed=true: 3.883 s before, 2.603 ms after, -99.93% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraph">Render commit graph</strong><br>
Render a graph of 50 commits generated with seed 1234.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraph-0d467b36b902.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraph: 178.3 µs before, 123 µs after, -30.98% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=box/selected=false">Render commit graph variants</strong><br>
Render 50 commits generated with seed 1234, using box symbols; no selected commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-50-symbols-box-selected-false-aa5158aa5e71.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=box/selected=false: 145.2 µs before, 115.1 µs after, -20.71% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=box/selected=true">Render commit graph variants</strong><br>
Render 50 commits generated with seed 1234, using box symbols; select the middle commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-50-symbols-box-selected-true-67c35865c0f8.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=box/selected=true: 169.3 µs before, 127.5 µs after, -24.70% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=branch/selected=false">Render commit graph variants</strong><br>
Render 50 commits generated with seed 1234, using branch symbols; no selected commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-50-symbols-branch-selected-false-9aa0c8e3616b.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=branch/selected=false: 170.9 µs before, 107.1 µs after, -37.35% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=branch/selected=true">Render commit graph variants</strong><br>
Render 50 commits generated with seed 1234, using branch symbols; select the middle commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-50-symbols-branch-selected-true-fc9f31352239.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=50/symbols=branch/selected=true: 169 µs before, 126.2 µs after, -25.32% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=box/selected=false">Render commit graph variants</strong><br>
Render 500 commits generated with seed 1234, using box symbols; no selected commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-500-symbols-box-selected-false-021e11cebb35.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=box/selected=false: 10.95 ms before, 6.417 ms after, -41.38% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=box/selected=true">Render commit graph variants</strong><br>
Render 500 commits generated with seed 1234, using box symbols; select the middle commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-500-symbols-box-selected-true-f40f1fd89716.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=box/selected=true: 12.73 ms before, 6.648 ms after, -47.76% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=branch/selected=false">Render commit graph variants</strong><br>
Render 500 commits generated with seed 1234, using branch symbols; no selected commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-500-symbols-branch-selected-false-dd1b4b08e6bb.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=branch/selected=false: 11.65 ms before, 9.291 ms after, -20.25% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=branch/selected=true">Render commit graph variants</strong><br>
Render 500 commits generated with seed 1234, using branch symbols; select the middle commit.</p></td>
<td><img src="docs-master/dev/performance-comparisons/gui-presentation-graph-BenchmarkRenderCommitGraphVariants-commits-500-symbols-branch-selected-true-4d262970c807.svg" width="420" height="60" align="middle" alt="gui/presentation/graph/BenchmarkRenderCommitGraphVariants/commits=500/symbols=branch/selected=true: 12.03 ms before, 6.718 ms after, -44.16% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="utils/BenchmarkEscapeSpecialChars/controls">Escape special chars</strong><br>
Escape control characters in a filename containing tabs and newlines.</p></td>
<td><img src="docs-master/dev/performance-comparisons/utils-BenchmarkEscapeSpecialChars-controls-e7496b440420.svg" width="420" height="60" align="middle" alt="utils/BenchmarkEscapeSpecialChars/controls: 2.223 µs before, 117.8 ns after, -94.70% change"></td>
</tr>
<tr>
<td width="380" valign="top"><p><strong title="utils/BenchmarkEscapeSpecialChars/ordinary">Escape special chars</strong><br>
Escape control characters in a plain filename.</p></td>
<td><img src="docs-master/dev/performance-comparisons/utils-BenchmarkEscapeSpecialChars-ordinary-3b83e34221ae.svg" width="420" height="60" align="middle" alt="utils/BenchmarkEscapeSpecialChars/ordinary: 2.148 µs before, 24.8 ns after, -98.85% change"></td>
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
