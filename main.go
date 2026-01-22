package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
)

func main() {
	srcDir := flag.String("src", ".", "源目录（默认：当前目录）")
	dstDir := flag.String("dst", ".", "目标目录（默认：当前目录）")
	remove := flag.Bool("rm", false, "解压完成后删除原 .livp 文件")
	onlyMov := flag.Bool("only-mov", false, "仅导出 .mov 视频文件（Live Photo 视频）")
	onlyImage := flag.Bool("only-image", false, "仅导出图片文件（heic / jpg / png）")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `
livp-extract
用于批量解压 Apple Live Photo (.livp) 文件。

每个 .livp 文件将被解压为：
  - 一张图片（.heic / .jpg / .png）
  - 一个视频文件（.mov）

用法:
  livp-extract [参数]

参数:
`)
		flag.PrintDefaults()

		fmt.Fprintf(flag.CommandLine.Output(), `
示例:
  livp-extract
      在当前目录查找 .livp 并解压到当前目录

  livp-extract -src ./input -dst ./output
      指定源目录和目标目录

  livp-extract --only-mov
      仅导出 .mov 视频文件

  livp-extract --only-image -rm
      仅导出图片并删除原 .livp 文件
`)
	}

	flag.Parse()

	// 参数互斥校验
	if *onlyMov && *onlyImage {
		fmt.Println("错误：--only-mov 与 --only-image 不能同时使用")
		os.Exit(1)
	}

	srcAbs, err := filepath.Abs(*srcDir)
	if err != nil {
		log.Fatal(err)
	}

	dstAbs, err := filepath.Abs(*dstDir)
	if err != nil {
		log.Fatal(err)
	}

	livpFiles, err := ScanLivpFiles(srcAbs)
	if err != nil {
		log.Fatal(err)
	}

	if len(livpFiles) == 0 {
		fmt.Println("未找到 .livp 文件")
		return
	}

	bar := progressbar.NewOptions(len(livpFiles),
		progressbar.OptionSetDescription("解压中"),
		progressbar.OptionShowCount(),
		progressbar.OptionClearOnFinish(),
	)

	for _, livp := range livpFiles {
		bar.Describe(filepath.Base(livp))

		if err := ExtractLivp(livp, dstAbs, *onlyMov, *onlyImage); err != nil {
			log.Println("失败:", err)
		} else if *remove {
			_ = RemoveFile(livp)
		}

		bar.Add(1)
	}

	fmt.Println("完成 ✔")
}
