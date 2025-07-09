package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent"
	"github.com/cloudwego/eino/flow/agent/multiagent/host"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
)

func newAddSpecialist(ctx context.Context) (*host.Specialist, error) {
	arkAPIKey := "56a6b406-8b6b-4bb5-b169-92117a5caa72"
	arkModelName := "doubao-1-5-pro-32k-250115"
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: arkAPIKey,
		Model:  arkModelName,
	})
	if err != nil {
		fmt.Printf("failed to create chat model: %v", err)
		return nil,err
	}
	addtool:=GetAddTool()
	raAgent, err := react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: chatModel,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: []tool.BaseTool{addtool},
		},
	})
	if err!= nil {
		return nil, err
	}
	return &host.Specialist{
		AgentMeta: host.AgentMeta{
			Name:        "add_specialist",
			IntendedUse: "add two numbers and return the result",
		},
		Invokable: func(ctx context.Context, input []*schema.Message, opts ...agent.AgentOption) (*schema.Message, error) {
			return raAgent.Generate(ctx, input, opts...)
		},
	}, nil
}
