package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	volcvikingdb "github.com/cloudwego/eino-ext/components/indexer/volc_vikingdb"
	"github.com/cloudwego/eino/schema"
)

func IndexerRAG() {
	ctx := context.Background()
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("EMBEDDER"),
	})
	if err != nil {
		panic(err)
	}
	idx, err := volcvikingdb.NewIndexer(ctx, &volcvikingdb.IndexerConfig{
		Host:       "api-vikingdb.volces.com",
		Region:     "cn-beijing",
		AK:         os.Getenv("VDB_AK"),
		SK:         os.Getenv("VDB_SK"),
		Scheme:     "https",
		Collection: "AwesomeEino",
		EmbeddingConfig: volcvikingdb.EmbeddingConfig{
			Embedding:  embedder,
			UseBuiltin: false,
		},
	})
	if err != nil {
		panic(err)
	}
	docs := []*schema.Document{
		{Content: "原神启动！！！"},
	}
	ids, err := idx.Store(ctx, docs)
	if err != nil {
		panic(err)
	}
	for i, id := range ids {
		println("doc", i+1, "id:", id)
	}
}
