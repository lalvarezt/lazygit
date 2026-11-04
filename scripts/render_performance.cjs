const crypto = require('node:crypto');

function escapeXml(value) {
  return String(value).replace(/[&<>"']/g, chr => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' })[chr]);
}

function time(ns) {
  const [unit, scale] = ns >= 1e9 ? ['s', 1e9] : ns >= 1e6 ? ['ms', 1e6] : ns >= 1e3 ? ['µs', 1e3] : ['ns', 1];
  return `${Number((ns / scale).toPrecision(4))} ${unit}`;
}

function text(x, y, value, className, anchor = 'start') {
  return `<text x="${x}" y="${y}" class="${className}" text-anchor="${anchor}">${escapeXml(value)}</text>`;
}

function renderTiming(row) {
  const maximum = Math.max(row.before, row.after);
  const delta = (row.after - row.before) / row.before;
  const percentage = delta.toLocaleString('en-US', {
    style: 'percent', signDisplay: 'exceptZero', minimumFractionDigits: 2, maximumFractionDigits: 2,
  });
  const label = `${row.id}: ${time(row.before)} before, ${time(row.after)} after, ${percentage} change`;
  const bars = ['before', 'after'].map((side, i) => {
    const y = 10 + i * 31;
    const width = Math.max(1, 180 * row[side] / maximum);
    return [
      text(12, y + 10, side === 'before' ? 'Before' : 'After', 'bar-label'),
      `<rect x="62" y="${y}" width="180" height="12" rx="3" class="track"/>`,
      `<rect x="62" y="${y}" width="${width.toFixed(3)}" height="12" rx="3" class="${side}"/>`,
      text(314, y + 10, time(row[side]), 'value', 'end'),
    ].join('\n');
  }).join('\n');
  const svg = `<svg xmlns="http://www.w3.org/2000/svg" width="420" height="60" viewBox="0 0 420 60" class="performance-timing" role="img" aria-labelledby="timing-title" data-benchmark="${escapeXml(row.id)}" data-before="${row.before}" data-after="${row.after}">
<title id="timing-title">${escapeXml(label)}</title>
<style>
.performance-timing{color-scheme:light dark;--performance-fg:#1f2328;--performance-muted:#59636e;--performance-track:#f0f2f5;--performance-before:#8c959f;--performance-after:#0f766e;--performance-loss:#cf222e}
@media(prefers-color-scheme:dark){.performance-timing{--performance-fg:#f0f6fc;--performance-muted:#9198a1;--performance-track:#161b22;--performance-before:#9198a1;--performance-after:#2dd4bf;--performance-loss:#f85149}}
text{font-family:system-ui,-apple-system,BlinkMacSystemFont,"Segoe UI",sans-serif;fill:var(--performance-fg);font-size:14px}.bar-label{font-size:12px;fill:var(--performance-muted)}.value{font-variant-numeric:tabular-nums}.track{fill:var(--performance-track)}.before{fill:var(--performance-before)}.after{fill:var(--performance-after)}.gain{fill:var(--performance-after);font-size:15px;font-weight:600}.loss{fill:var(--performance-loss);font-size:15px;font-weight:600}
</style>
${bars}
${text(408, 35, percentage, delta < 0 ? 'gain' : delta > 0 ? 'loss' : 'value', 'end')}
</svg>\n`;
  return { svg, label };
}

function renderComparison({ rows }, imageDirectory) {
  const images = [];
  const tableRows = rows.map(row => {
    const parts = row.id.split('/');
    const benchmarkIndex = parts.findIndex(part => part.startsWith('Benchmark'));
    const name = parts[benchmarkIndex].slice('Benchmark'.length)
      .replace(/([A-Z]+)([A-Z][a-z])/g, '$1 $2')
      .replace(/([a-z0-9])([A-Z])/g, '$1 $2')
      .split(' ').map((word, i) => i === 0 ? word : word.toLowerCase()).join(' ');
    const slug = row.id.replace(/[^a-zA-Z0-9_-]+/g, '-');
    const { svg, label } = renderTiming(row);
    const digest = crypto.createHash('sha256').update(svg).digest('hex').slice(0, 12);
    const fileName = `${slug}-${digest}.svg`;
    images.push({ fileName, svg });
    return `<tr>
<td width="380" valign="top"><p><strong title="${escapeXml(row.id)}">${escapeXml(name)}</strong><br>
${escapeXml(row.description)}</p></td>
<td><img src="${imageDirectory}/${fileName}" width="420" height="60" align="middle" alt="${escapeXml(label)}"></td>
</tr>`;
  }).join('\n');
  const table = `<table>
<thead><tr><th align="left">Benchmark / workload</th><th align="left">Before / after and change</th></tr></thead>
<tbody>
${tableRows}
</tbody>
</table>`;
  return { table, images };
}

module.exports = { renderComparison };
