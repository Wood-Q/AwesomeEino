package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/retriever/milvus"
)

func RetrieverRAG() {
	ctx := context.Background()
	timeout := 30 * time.Second
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("EMBEDDER"),
		Timeout: &timeout,
	})
	if err != nil {
		panic(err)
	}
	retriever, err := milvus.NewRetriever(ctx, &milvus.RetrieverConfig{
		Client:            MilvusCli,
		Collection:        "AwesomeEino",
		Partition:         nil,
		VectorField:       "vector",
		OutputFields: []string{
			"id",
			"content",
			"metadata",
		},
		TopK:      1,
		Embedding: embedder,
	})
	if err != nil {
		panic(err)
	}

	results, err := retriever.Retrieve(ctx, "原神")
	if err != nil {
		panic(err)
	}

	log.Printf("Results: %v", results)
}
