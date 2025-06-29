package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // 加载环境变量
	if err != nil {
		log.Fatal("Error loading .env file") // 处理加载错误
	}
	SimpleAgent()
	//BuildRAG()
	// TransDoc()
	//RetrieverRAG()
	//IndexerRAG()
	// EmbedText()
	//TemplateChat()
	//ChatStream()
	//ChatGenerate()
}
