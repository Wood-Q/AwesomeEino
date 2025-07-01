package stage7

import (
	"AwesomeEino/stage4"
	"AwesomeEino/stage5"
	"AwesomeEino/stage6"
)

func BuildRAG() {
	docs := stage6.TransDoc()
	stage4.IndexerRAG(docs)
	results := stage5.RetrieverRAG("欲渡黄河冰塞川")
	for _, doc := range results {
		println(doc.ID)
		println("================================================")
		println(doc.Content)
	}
}
