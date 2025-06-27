package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino/schema"
)

func TransDoc() {
	ctx := context.Background()

	// 初始化分割器
	splitter, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		Headers: map[string]string{
			"#":   "h1",
			"##":  "h2",
			"###": "h3",
		},
		TrimHeaders: false,
	})
	if err != nil {
		panic(err)
	}

	// 准备要分割的文档
	content, err := os.OpenFile("./document.md", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer content.Close()
	bs, err := os.ReadFile("./document.md")
	if err != nil {
		panic(err)
	}
	docs := []*schema.Document{
		{
			ID:      "doc1",
			Content: string(bs),
		},
	}

	// 执行分割
	results, err := splitter.Transform(ctx, docs)
	if err != nil {
		panic(err)
	}

	// 处理分割结果
	for i, doc := range results {
		println("片段", i+1, ":", doc.Content)
		println("标题层级：")
		for k, v := range doc.MetaData {
			if k == "h1" || k == "h2" || k == "h3" {
				println("  ", k, ":", v)
			}
		}
	}
}
