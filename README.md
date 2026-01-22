# livp-extract

**livp-extract** 是一个用于批量解压 Apple Live Photo (.livp) 文件的命令行工具。  
每个 `.livp` 文件会被解压为一张图片和一个视频文件（.mov），支持灵活的参数控制和进度条显示。

---

## 功能特点

- 批量扫描指定目录中的 `.livp` 文件
- 自动解压为：
    - 图片文件（.heic / .jpg / .png）
    - 视频文件（.mov）
- 支持命令行参数：
    - 指定源目录和目标目录
    - 仅导出视频 (`--only-mov`)
    - 仅导出图片 (`--only-image`)
    - 解压后删除原 `.livp` 文件 (`-rm`)
- 带进度条显示解压进度
- 中文帮助文档 (`-h`)
- 参数互斥检查，避免误操作
- 支持 macOS ARM / Intel，Linux 系统

---

## 安装

### 1. 克隆项目

```bash
git clone https://github.com/yourusername/livp-extract.git
cd livp-extract
```

### 2. 获取依赖

```
go mod tidy
```

### 3. 构建可执行文件

```
go build -o livp-extract
```

### 4. 可选：移动到 PATH 目录

```
mv livp-extract /usr/local/bin/
# 或者 macOS ARM
mv livp-extract /opt/homebrew/bin/
```

------

## 使用方法

```
./livp-extract [参数]
```

### 参数说明

| 参数           | 说明                               |
| -------------- | ---------------------------------- |
| `-src`         | 源目录，默认当前目录               |
| `-dst`         | 目标目录，默认当前目录             |
| `-rm`          | 解压完成后删除原 `.livp` 文件      |
| `--only-mov`   | 仅导出 `.mov` 视频文件             |
| `--only-image` | 仅导出图片文件（heic / jpg / png） |
| `-h`           | 显示帮助信息                       |

⚠️ `--only-mov` 和 `--only-image` 不能同时使用，会报错并退出。

------

### 使用示例

```
# 当前目录查找 .livp 并解压到当前目录
./livp-extract

# 指定源目录和目标目录
./livp-extract -src ./input -dst ./output

# 仅导出视频
./livp-extract --only-mov

# 仅导出图片并删除原 .livp
./livp-extract --only-image -rm
```

------

## 示例展示

### 终端运行效果

```
$ ./livp-extract -src ./LivePhotos -dst ./Exported -rm
解压中 [██████████--------------------] 5/20 IMG_1234.livp
解压中 [█████████████---------------] 10/20 IMG_1245.livp
解压中 [████████████████------------] 15/20 IMG_1256.livp
解压中 [████████████████████--------] 20/20 IMG_1267.livp
完成 ✔
```

### 解压结果目录结构示例

```
Exported/
├── IMG_1234.heic
├── IMG_1234.mov
├── IMG_1245.heic
├── IMG_1245.mov
├── IMG_1256.heic
├── IMG_1256.mov
├── IMG_1267.heic
└── IMG_1267.mov
```

## 注意事项

- macOS / Linux 可直接使用 `go build` 构建。
- 对于大目录，建议使用 SSD 或高速硬盘以提高解压速度。
- 默认解压到目标目录，文件名保持与 `.livp` 相同。

------

## 项目目录结构

```
livp-extract/
├── main.go         # 程序入口，参数解析，进度条
├── scanner.go      # 扫描目录查找 .livp 文件
├── extractor.go    # 解压逻辑
├── utils.go        # 工具函数
└── go.mod          # Go 模块管理
```

