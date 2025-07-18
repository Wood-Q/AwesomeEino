package stage9

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/compose"
)

func OrcGraph(choice string) {
	ctx := context.Background()
	g := compose.NewGraph[string, string]()
	lambda0 := compose.InvokableLambda(func(ctx context.Context, input string) (output string, err error) {
		if input == "1" {
			return "毫猫", nil
		} else if input == "2" {
			return "耄耋", nil
		} else if input == "3" {
			return "device", nil
		}
		return "", nil
	})
	lambda1 := compose.InvokableLambda(func(ctx context.Context, input string) (output string, err error) {
		return "喵！", nil
	})
	lambda2 := compose.InvokableLambda(func(ctx context.Context, input string) (output string, err error) {
		return "哈！", nil
	})
	lambda3 := compose.InvokableLambda(func(ctx context.Context, input string) (output string, err error) {
		return "没有人类了!!!", nil
	})
	//加入节点
	err := g.AddLambdaNode("lambda0", lambda0)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("lambda1", lambda1)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("lambda2", lambda2)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("lambda3", lambda3)
	if err != nil {
		panic(err)
	}
	//加入分支
	err = g.AddBranch("lambda0", compose.NewGraphBranch(func(ctx context.Context, in string) (endNode string, err error) {
		if in == "毫猫" {
			return "lambda1", nil
		} else if in == "耄耋" {
			return "lambda2", nil
		} else if in == "device" {
			return "lambda3", nil
		}
		// 否则，返回 compose.END，表示流程结束
		return compose.END, nil
	}, map[string]bool{"lambda1": true, "lambda2": true, "lambda3": true, compose.END: true}))
	if err != nil {
		panic(err)
	}
	//分支连接
	err = g.AddEdge(compose.START, "lambda0")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("lambda1", compose.END)
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("lambda2", compose.END)
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("lambda3", compose.END)
	if err != nil {
		panic(err)
	}
	// 编译
	r, err := g.Compile(ctx)
	if err != nil {
		panic(err)
	}
	// 执行
	answer, err := r.Invoke(ctx, choice)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)

}
