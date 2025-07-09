package stage9

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func OrcChain() {
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
	lambda := compose.InvokableLambda(func(ctx context.Context, input string) (output []*schema.Message, err error) {
		desuwa := input + "回答结尾加上desuwa"
		output = []*schema.Message{
			{
				Role:    schema.User,
				Content: desuwa,
			},
		}
		return output, nil
	})

	chain := compose.NewChain[string, *schema.Message]()
	chain.AppendLambda(lambda).AppendChatModel(model)
	r, err := chain.Compile(ctx)
	if err != nil {
		panic(err)
	}
	answer, err := r.Invoke(ctx, "你好，请告诉我你的名字")
	if err != nil {
		panic(err)
	}
	fmt.Println(answer.Content)
}
