package stage6

import (
	"context"
	"os"
	"strconv"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
)

func TransDoc() []*schema.Document {
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
			ID:      uuid.New().String(),
			Content: string(bs),
		},
	}

	// 执行分割
	results, err := splitter.Transform(ctx, docs)
	if err != nil {
		panic(err)
	}

	for i, doc := range results {
		doc.ID = docs[0].ID + "_" + strconv.Itoa(i)
		println(doc.ID)
	}

	return results
}
