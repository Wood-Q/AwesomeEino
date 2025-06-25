package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/retriever/volc_vikingdb"
)

func RetrieverRAG(query string) {
	ctx := context.Background()
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("EMBEDDER"),
	})
	if err != nil {
		panic(err)
	}
	ret, err := volc_vikingdb.NewRetriever(ctx, &volc_vikingdb.RetrieverConfig{
		Host:       "api-vikingdb.volces.com",
		Region:     "cn-beijing",
		AK:         os.Getenv("VDB_AK"),
		SK:         os.Getenv("VDB_SK"),
		Scheme:     "https",
		Collection: "AwesomeEino",
		Index:      "Find",
		EmbeddingConfig: volc_vikingdb.EmbeddingConfig{
			UseBuiltin: false,
			Embedding:  embedder,
		},
	})
	docs, err := ret.Retrieve(ctx, query)
	if err != nil {
		panic(err)
	}
	print("找到辣：", docs[0].Content)
}
