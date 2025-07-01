package stage3

import (
	"context"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
)

func EmbedText() {
	ctx := context.Background()

	// 初始化嵌入器
	timeout := 30 * time.Second
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("EMBEDDER"),
		Timeout: &timeout,
	})
	if err != nil {
		panic(err)
	}

	// 生成文本向量
	texts := []string{
		"这是第一段示例文本",
		"这是第二段示例文本",
	}

	embeddings, err := embedder.EmbedStrings(ctx, texts)
	if err != nil {
		panic(err)
	}

	// 使用生成的向量
	for i, embedding := range embeddings {
		println("文本", i+1, "的向量维度:", len(embedding))
	}
}
