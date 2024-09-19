import * as fs from "node:fs";

interface NodePart {
  renderedLength: number;
  gzipLength: number;
  brotliLength: number;
}

interface BundleData {
  nodeParts: { [key: string]: NodePart };
}

function getBundleSize(jsonFile: string): number {
  const data: BundleData = JSON.parse(fs.readFileSync(jsonFile, "utf8"));
  const nodeParts = data.nodeParts;

  let totalSize = 0;
  for (const uid in nodeParts) {
    totalSize += nodeParts[uid].renderedLength || 0;
  }
  return totalSize;
}

function formatBytes(bytes: number): string {
  const sizes = ["Bytes", "KB", "MB"];
  if (bytes === 0) return "0 Bytes";
  const i = Math.floor(Math.log(bytes) / Math.log(1024));
  return `${Number.parseFloat((bytes / 1024 ** i).toFixed(2))} ${sizes[i]}`;
}

const prBundleSize = getBundleSize("__bundle__/result.json");
const baseBundleSize = getBundleSize("__bundle__/base.json");
const diff = prBundleSize - baseBundleSize;
const diffPercentage = ((diff / baseBundleSize) * 100).toFixed(2);

const table = `
## 📦 バンドルサイズ比較

|                   | 現在のブランチ | ベースブランチ | 増加分       | 増加率     |
|-------------------|---------------|---------------|-------------|-----------|
| バンドルサイズ    | ${formatBytes(prBundleSize)} | ${formatBytes(
  baseBundleSize
)} | ${diff >= 0 ? "+" : ""}${formatBytes(diff)} | ${
  diff >= 0 ? "+" : ""
}${diffPercentage}% |
`;

console.log(table);
