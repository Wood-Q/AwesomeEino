package stage9

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

type State struct {
	History map[string]any
}

func genFunc(ctx context.Context) *State {
	return &State{
		History: make(map[string]any),
	}
}

func OrcGraphWithState(ctx context.Context, input map[string]string) {
	g := compose.NewGraph[map[string]string, *schema.Message](
		compose.WithGenLocalState(genFunc),
	)
	lambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output map[string]string, err error) {
		//在节点内部处理state
		_ = compose.ProcessState[*State](ctx, func(_ context.Context, state *State) error {
			state.History["tsundere_action"] = "我喜欢你"
			state.History["cute_action"] = "摸摸头"
			return nil
		})
		if input["role"] == "tsundere" {
			return map[string]string{"role": "tsundere", "content": input["content"]}, nil
		}
		if input["role"] == "cute" {
			return map[string]string{"role": "cute", "content": input["content"]}, nil
		}
		return map[string]string{"role": "user", "content": input["content"]}, nil
	})
	TsundereLambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output []*schema.Message, err error) {
		_ = compose.ProcessState[*State](ctx, func(_ context.Context, state *State) error {
			input["content"] = input["content"] + state.History["tsundere_action"].(string)
			return nil
		})
		return []*schema.Message{
			{
				Role:    schema.System,
				Content: "你是一个高冷傲娇的大小姐，每次都会用傲娇的语气回答我的问题",
			},
			{
				Role:    schema.User,
				Content: input["content"],
			},
		}, nil
	})

	CuteLambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output []*schema.Message, err error) {
		// _ = compose.ProcessState[*State](ctx, func(_ context.Context, state *State) error {
		// 	input["content"] = input["content"] + state.History["action"].(string)
		// 	return nil
		// })
		return []*schema.Message{
			{
				Role:    schema.System,
				Content: "你是一个可爱的小女孩，每次都会用可爱的语气回答我的问题",
			},
			{
				Role:    schema.User,
				Content: input["content"],
			},
		}, nil
	})

	cutePreHandler := func(ctx context.Context, input map[string]string, state *State) (map[string]string, error) {
		input["content"] = input["content"] + state.History["cute_action"].(string)
		return input, nil
	}

	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	if err != nil {
		panic(err)
	}
	//注册节点
	err = g.AddLambdaNode("lambda", lambda)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("tsundere", TsundereLambda)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("cute", CuteLambda, compose.WithStatePreHandler(cutePreHandler))
	if err != nil {
		panic(err)
	}
	err = g.AddChatModelNode("model", model)
	if err != nil {
		panic(err)
	}
	//加入分支
	g.AddBranch("lambda", compose.NewGraphBranch(func(ctx context.Context, in map[string]string) (endNode string, err error) {
		if in["role"] == "tsundere" {

			return "tsundere", nil
		}
		if in["role"] == "cute" {
			return "cute", nil
		}
		return "tsundere", nil
	}, map[string]bool{"tsundere": true, "cute": true}))

	//链接节点
	err = g.AddEdge(compose.START, "lambda")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("tsundere", "model")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("cute", "model")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("model", compose.END)
	if err != nil {
		panic(err)
	}
	//编译
	r, err := g.Compile(ctx)
	if err != nil {
		panic(err)
	}
	//执行
	answer, err := r.Invoke(ctx, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer.Content)
}
