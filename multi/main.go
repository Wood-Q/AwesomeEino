package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/eino/flow/agent/multiagent/host"
	"github.com/cloudwego/eino/schema"
)

func main(){
	// 记录程序开始时间
	startTime := time.Now()
	fmt.Printf("程序开始执行时间: %s\n", startTime.Format("2006-01-02 15:04:05.000"))

	ctx := context.Background()
	h, err := newHost(ctx)
	if err != nil {
		panic(err)
	}
	adder,err:=newAddSpecialist(ctx)
	if err!= nil {
		panic(err)
	}
	suber,err:=newSubSpecialist(ctx)
	if err!= nil {
		panic(err)
	}
	hostMA, err := host.NewMultiAgent(ctx, &host.MultiAgentConfig{
		Host: *h,
		Specialists: []*host.Specialist{
			adder,
			suber,
		},
		Summarizer: &host.Summarizer{
			ChatModel:    h.ChatModel,
			SystemPrompt: "请总结一下各个专家的回答",
		},
	})
	if err!= nil {
		panic(err)
	}
	cb := &logCallback{}
	msg := &schema.Message{
			Role:    schema.User,
			Content: "帮我计算1239+231-222",
		}
	out, err := hostMA.Generate(ctx, []*schema.Message{msg}, host.WithAgentCallbacks(cb))
	if err != nil {
		panic(err)
	}
	println(out.Content)

	// 计算并打印运行时间
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("\n程序结束执行时间: %s\n", endTime.Format("2006-01-02 15:04:05.000"))
	fmt.Printf("总运行时间: %v\n", duration)
	fmt.Printf("总运行时间(毫秒): %.2f ms\n", float64(duration.Nanoseconds())/1000000)

	// for { // 多轮对话，除非用户输入了 "exit"，否则一直循环
	// 	println("\n\nYou: ") // 提示轮到用户输入了

	// 	var message string
	// 	scanner := bufio.NewScanner(os.Stdin) // 获取用户在命令行的输入
	// 	for scanner.Scan() {
	// 		message += scanner.Text()
	// 		break
	// 	}

	// 	if err := scanner.Err(); err != nil {
	// 		panic(err)
	// 	}

	// 	if message == "exit" {
	// 		return
	// 	}

	// 	msg := &schema.Message{
	// 		Role:    schema.User,
	// 		Content: message,
	// 	}

	// 	out, err := hostMA.Generate(ctx, []*schema.Message{msg}, host.WithAgentCallbacks(cb))
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	println("\nAnswer:")
	// 	println(out.Content)
	// }
}
type logCallback struct{}

func (l *logCallback) OnHandOff(ctx context.Context, info *host.HandOffInfo) context.Context {
	println("\nHandOff to", info.ToAgentName, "with argument", info.Argument)
	return ctx
}

// type loggerCallback struct {
// 	callbacks.HandlerBuilder
// }

// func (cb *loggerCallback) OnStart(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
// 	fmt.Println("==================")
// 	inputStr, _ := json.MarshalIndent(input, "", "  ") // nolint: byted_s_returned_err_check
// 	fmt.Printf("[OnStart] %s\n", string(inputStr))
// 	return ctx
// }

// func (cb *loggerCallback) OnEnd(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
// 	fmt.Println("=========[OnEnd]=========")
// 	outputStr, _ := json.MarshalIndent(output, "", "  ") // nolint: byted_s_returned_err_check
// 	fmt.Println(string(outputStr))
// 	return ctx
// }

// func (cb *loggerCallback) OnError(ctx context.Context, info *callbacks.RunInfo, err error) context.Context {
// 	fmt.Println("=========[OnError]=========")
// 	fmt.Println(err)
// 	return ctx
// }

// func (cb *loggerCallback) OnEndWithStreamOutput(ctx context.Context, info *callbacks.RunInfo,
// 	output *schema.StreamReader[callbacks.CallbackOutput]) context.Context {

// 	var graphInfoName = react.GraphName

// 	go func() {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				fmt.Println("[OnEndStream] panic err:", err)
// 			}
// 		}()

// 		defer output.Close() // remember to close the stream in defer

// 		fmt.Println("=========[OnEndStream]=========")
// 		for {
// 			frame, err := output.Recv()
// 			if errors.Is(err, io.EOF) {
// 				// finish
// 				break
// 			}
// 			if err != nil {
// 				fmt.Printf("internal error: %s\n", err)
// 				return
// 			}

// 			s, err := json.Marshal(frame)
// 			if err != nil {
// 				fmt.Printf("internal error: %s\n", err)
// 				return
// 			}

// 			if info.Name == graphInfoName { // 仅打印 graph 的输出, 否则每个 stream 节点的输出都会打印一遍
// 				fmt.Printf("%s: %s\n", info.Name, string(s))
// 			}
// 		}

// 	}()
// 	return ctx
// }

// func (cb *loggerCallback) OnStartWithStreamInput(ctx context.Context, info *callbacks.RunInfo,
// 	input *schema.StreamReader[callbacks.CallbackInput]) context.Context {
// 	defer input.Close()
// 	return ctx
// }
