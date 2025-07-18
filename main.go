package main

import (
	"AwesomeEino/stage9"
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // 加载环境变量
	if err != nil {
		log.Fatal("Error loading .env file") // 处理加载错误
	}
	ctx := context.Background()
	// stage9.OrcGraphWithState(ctx, map[string]string{"role": "cute", "content": "你好啊"})

	r, err := stage10.Buildtest(ctx)
	if err != nil {
		panic(err)
	}
	variables := map[string]any{
		"role": "可爱的女子高中生",
		"task": "安慰一下我",
	}
	output, err := r.Invoke(ctx, variables)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
