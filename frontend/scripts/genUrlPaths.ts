import chokidar from "chokidar";
import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// src/routesディレクトリへのパスを設定
const baseDir = path.join(__dirname, "..", "src", "routes");

// 保存先のpath.jsonのパスを設定
const outputPath = path.join(__dirname, "..", "path.json");

function findIndexTsxFiles(dir: string): string[] {
  let urls: string[] = [];
  const files = fs.readdirSync(dir);
  for (const file of files) {
    const fullPath = path.join(dir, file);
    const stat = fs.statSync(fullPath);
    if (stat.isDirectory()) {
      urls = urls.concat(findIndexTsxFiles(fullPath));
    } else if (stat.isFile() && file === "index.tsx") {
      const relativeDirPath = path.relative(baseDir, dir);
      const urlPath = "/" + relativeDirPath.split(path.sep).join("/");
      urls.push(urlPath);
    }
  }
  return urls;
}

function writeUrlsToFile(urls: string[]) {
  // URLパスをJSONとして保存
  fs.writeFileSync(outputPath, JSON.stringify(urls, null, 2), "utf-8");
}

function updateUrls() {
  const urlPaths = findIndexTsxFiles(baseDir);
  writeUrlsToFile(urlPaths);
}

// 初期実行
updateUrls();

// chokidarを使用してディレクトリを監視
const watcher = chokidar.watch(baseDir, {
  persistent: true,
  ignoreInitial: true,
});

// ファイルの追加、変更、削除を監視
watcher.on("all", (event, filePath) => {
  if (
    ["add", "change", "unlink"].includes(event) &&
    filePath.endsWith("index.tsx")
  ) {
    updateUrls();
  }
});
