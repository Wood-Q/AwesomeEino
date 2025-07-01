package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func SimpleAgent() {

	getGameTool := CreateTool()
	ctx := context.Background()
	handler := callbacks.NewHandlerBuilder().
		OnStartFn(func(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
			return ctx
		}).
		OnEndFn(func(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
			fmt.Println("=========[OnEnd]=========")
			outputStr, _ := json.MarshalIndent(output, "", "  ")
			fmt.Println(string(outputStr))
			fmt.Println("=========[OnEnd]=========")
			return ctx
		}).
		OnErrorFn(func(ctx context.Context, info *callbacks.RunInfo, err error) context.Context {
			return ctx
		}).
		Build()
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
	//绑定工具
	info, err := getGameTool.Info(ctx)
	if err != nil {
		panic(err)
	}
	infos := []*schema.ToolInfo{
		info,
	}
	err = model.BindTools(infos)
	if err != nil {
		panic(err)
	}
	//创建tools节点
	ToolsNode, err := compose.NewToolNode(context.Background(), &compose.ToolsNodeConfig{
		Tools: []tool.BaseTool{
			getGameTool,
		},
	})
	if err != nil {
		panic(err)
	}
	//创建完整的处理链
	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(model, compose.WithNodeName("chat_model")).
		AppendToolsNode(ToolsNode, compose.WithNodeName("tools"))

	// 编译并运行 chain
	agent, err := chain.Compile(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//运行Agent
	resp, err := agent.Invoke(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "请告诉我原神的URL是什么",
		},
	}, compose.WithCallbacks(handler))
	if err != nil {
		log.Fatal(err)
	}

	// 输出结果
	for _, msg := range resp {
		fmt.Println(msg.Content)
	}
}
