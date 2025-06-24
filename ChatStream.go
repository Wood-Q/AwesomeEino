package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func ChatStream() {
	err := godotenv.Load() // 加载环境变量
	if err != nil {
		log.Fatal("Error loading .env file") // 处理加载错误
	}
	ctx := context.Background()

	timeout := 30 * time.Second
	// 初始化模型
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   "doubao-1.5-pro-32k-250115",
		Timeout: &timeout,
	})
	if err != nil {
		panic(err)
	}

	// 准备消息
	messages := []*schema.Message{
		schema.SystemMessage("你是一个助手"),
		schema.UserMessage("你好"),
	}

	// 获取流式回复
	reader, err := model.Stream(ctx, messages)
	if err != nil {
		panic(err)
	}
	defer reader.Close() // 注意要关闭

	// 处理流式内容
	for {
		chunk, err := reader.Recv()
		if err != nil {
			break
		}
		print(chunk.Content)
	}
}
