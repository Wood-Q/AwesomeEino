package main

import (
	"AwesomeEino/stage9"
	"context"
	"fmt"
	"log"

	ccb "github.com/cloudwego/eino-ext/callbacks/cozeloop"
	"github.com/cloudwego/eino/callbacks"
	"github.com/coze-dev/cozeloop-go"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // 加载环境变量
	if err != nil {
		log.Fatal("Error loading .env file") // 处理加载错误
	}
	ctx := context.Background()
	client, err := cozeloop.NewClient()
    if err != nil {
       panic(err)
    }
    defer client.Close(ctx)
    // 在服务 init 时 once 调用
    handler := ccb.NewLoopHandler(client)
    callbacks.AppendGlobalHandlers(handler)
	// stage9.OrcGraphWithModel(ctx, map[string]string{"role": "cute", "content": "你好啊"})
	// stage9.OrcGraphWithState(ctx, map[string]string{"role": "cute", "content": "你好啊"})
	// stage9.OrcGraphWithCallback(ctx, map[string]string{"role": "cute", "content": "你好啊"})
	g := stage9.OutSideOrcGraph(ctx)
	r, err := g.Compile(ctx)
	if err != nil {
		panic(err)
	}
	result, err := r.Invoke(ctx, map[string]string{"role": "cute", "content": "你好啊"})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// r, err := stage10.Buildtest(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// variables := map[string]any{
	// 	"role": "可爱的女子高中生",
	// 	"task": "安慰一下我",
	// }
	// output, err := r.Invoke(ctx, variables)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(output)
}
