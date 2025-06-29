package main

func BuildRAG() {
	docs := TransDoc()
	IndexerRAG(docs)
	results := RetrieverRAG("欲渡黄河冰塞川")
	for _, doc := range results {
		println(doc.ID)
		println("================================================")
		println(doc.Content)
	}
}
