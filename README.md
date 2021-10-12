bluge plugin
===

Analysis plugins for the [bluge](https://github.com/blugelabs/bluge/) indexing/search library.


# Usage example

blog: https://zenn.dev/ikawaha/articles/20201230-84b042603ccbbce645d5

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blugelabs/bluge"
	segment "github.com/blugelabs/bluge_segment_api"
	"github.com/ikawaha/blugeplugin/analysis/lang/ja"
)

func main() {
	// サンプルなので in memory で済ませる
	config := bluge.InMemoryOnlyConfig()

	// index writer を用意する
	w, err := bluge.OpenWriter(config)
	if err != nil {
		log.Fatalf("error opening writer: %v", err)
	}
	defer w.Close()

	// 対象ドキュメント（詳細は別途記載）
	docs := NewDocuments()

	// indexing
	for _, doc := range docs {
		doc := doc
		if err := w.Update(doc.ID(), doc); err != nil {
			log.Fatalf("error updating document: %v", err)
		}
		// 表示
		fmt.Printf("indexed document with id:%s\n", doc.ID())
		doc.EachField(func(field segment.Field) {
			fmt.Printf("\t%s: %s\n", field.Name(), field.Value())
		})
	}

	// index reader を用意する
	r, err := w.Reader()
	if err != nil {
		log.Fatalf("error getting index reader: %v", err)
	}
	defer r.Close()

	// クエリ
	q := "踊る人形"
	query := bluge.NewMatchQuery(q).SetAnalyzer(ja.Analyzer()).SetField("body")
	req := bluge.NewTopNSearch(10, query).WithStandardAggregations()
	fmt.Printf("query: search field %q, value %q\n", query.Field(), q)

	// search
	ite, err := r.Search(context.Background(), req)
	if err != nil {
		log.Fatalf("error executing search: %v", err)
	}
	// 検索結果
	for {
		match, err := ite.Next()
		if err != nil {
			log.Fatalf("error iterator document matches: %v", err)
		}
		if match == nil {
			break
		}
		if err := match.VisitStoredFields(func(field string, value []byte) bool {
			fmt.Printf("%s: %q\n", field, string(value))
			return true
		}); err != nil {
			log.Fatalf("error loading stored fields: %v", err)
		}
	}
}

func NewDocuments() []*bluge.Document {
	docs := []struct {
		ID     string
		Author string
		Text   string
	}{
		{
			ID:     "1:赤い蝋燭と人魚",
			Author: "小川未明",
			Text:   "人魚は南の方の海にばかり棲んでいるのではありません",
		},
		{
			ID:     "2:吾輩は猫である",
			Author: "夏目漱石",
			Text:   "吾輩は猫である。名前はまだない",
		},
		{
			ID:     "3:狐と踊れ",
			Author: "神林長平",
			Text:   "踊っているのでなければ踊らされているのだろうさ",
		},
		{
			ID:     "4:ダンスダンスダンス",
			Author: "村上春樹",
			Text:   "音楽の鳴っている間はとにかく踊り続けるんだ。おいらの言っていることはわかるかい？",
		},
	}
	var ret []*bluge.Document
	for _, v := range docs {
		auth := bluge.NewTextField("author", v.Author).WithAnalyzer(ja.Analyzer())
		body := bluge.NewTextField("body", v.Text).WithAnalyzer(ja.Analyzer())
		doc := bluge.NewDocument(v.ID).AddField(auth).AddField(body)
		ret = append(ret, doc)
	}
	return ret
}
```
