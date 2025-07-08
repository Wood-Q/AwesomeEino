package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type AddTool struct {
}

func GetAddTool() tool.InvokableTool {
	return &AddTool{}
}

type AddParam struct {
	A    int `json:"a"`
	B    int `json:"b"`
}

func (t *AddTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "add",
		Desc: "add two numbers",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"a": {
				Type:     "number",
				Desc:     "first number",
				Required: true,
			},
			"b": {
				Type:     "number",
				Desc:     "second number",
				Required: true,
			},
		}),
	}, nil
}

func (t *AddTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	p := &AddParam{}
	err := json.Unmarshal([]byte(argumentsInJSON), p)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", p.A+p.B), nil
}

type SubTool struct {
}

func GetSubTool() tool.InvokableTool {
	return &SubTool{}
}

type SubParam struct {
	A int `json:"a"`
	B int `json:"b"`
}

func (t *SubTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "sub",
		Desc: "sub two numbers",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"a": {
				Type:     "number",
				Desc:     "first number",
				Required: true,
			},
			"b": {
				Type:     "number",
				Desc:     "second number",
				Required: true,
			},
		}),
	}, nil
}

func (t *SubTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	p := &SubParam{}
	err := json.Unmarshal([]byte(argumentsInJSON), p)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", p.A-p.B), nil
}

type AnalyzeTool struct{
}

func GetAnalyzeTool() tool.InvokableTool {
	return &AnalyzeTool{}
}
type AnalyzeParam struct {
	Content string `json:"content"`
}

func (a *AnalyzeTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "analyze",
		Desc: "analyze the difficulty of the content",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"content": {
				Type:     "string",
				Desc:     "content to be analyzed",
				Required: true,
			},
		}),
	}, nil
}
func (a *AnalyzeTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts...tool.Option) (string, error) {
	// 解析输入参数
	p := &AnalyzeParam{}
	err := json.Unmarshal([]byte(argumentsInJSON), p)
	if err != nil {
		return "", err
	}
	//调用模型
	arkAPIKey := "56a6b406-8b6b-4bb5-b169-92117a5caa72"
	arkModelName := "doubao-1-5-pro-32k-250115"
	arkModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: arkAPIKey,
		Model:  arkModelName,
	})
	if err != nil {
		fmt.Printf("failed to create chat model: %v", err)
		return "", err
	}
	//调用模型
	AnalyzeInput:=[]*schema.Message{
		{
			Role:    schema.System,
			Content: "你是一个数学老师，你需要分析用户的问题，判断用户的问题的难度，难度分为简单，中等，困难，你需要根据用户的问题给出一个难度的评分，评分范围为1-10，1为简单，10为困难",
		},
		{
			Role:    schema.User,
			Content: p.Content,
		},
	}
	response,err:=arkModel.Generate(ctx,AnalyzeInput)
	if err!= nil {
		fmt.Printf("failed to generate: %v", err)
		return "", err
	}
	return response.Content,nil
}