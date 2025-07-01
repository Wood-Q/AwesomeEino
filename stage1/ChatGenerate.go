package stage1

import (
	"context"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
)

func ChatGenerate() {
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

	// 生成回复
	response, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}

	// 处理回复
	println(response.Content)

	// 获取 Token 使用情况
	if usage := response.ResponseMeta.Usage; usage != nil {
		println("提示 Tokens:", usage.PromptTokens)
		println("生成 Tokens:", usage.CompletionTokens)
		println("总 Tokens:", usage.TotalTokens)
	}
}
