package docx

/**
插件：github.com/nguyenthenguyen/docx
上传地址：https://cuttlefish.baidu.com/shopmis?_wkts_=1671617291075#/taskCenter/majorTask
过滤：ignore_word := ["/","党","新建文档","《","》",":","*","<",">","|","?","."]

请用以下题目《用例的作用》，写一篇320字以上的作文，并且分成至少3个段落
*/

import (
	"github.com/nguyenthenguyen/docx"
)

func Use() {
	title := "测试标题"
	content := "hhh点多1"

	r, err := docx.ReadDocxFile("./docx/pattern/template.docx")
	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()

	docx1.Replace("content", content, -1)
	docx1.WriteToFile("./docx/done/" + title + ".docx")
	r.Close()
}

func getContent() {

}

func Demo() {
	// Read from docx file
	r, err := docx.ReadDocxFile("./docx/TestDocument.docx")
	// Or read from memory
	// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)

	// Or read from a filesystem object:
	// r, err := docx.ReadDocxFromFS(file string, fs fs.FS)

	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()
	// Replace like https://golang.org/pkg/strings/#Replace
	docx1.Replace("old_1_1", "new_1_1", -1)
	docx1.Replace("old_1_2", "new_1_2", -1)
	docx1.ReplaceLink("http://example.com/", "https://github.com/nguyenthenguyen/docx", 1)
	docx1.ReplaceHeader("out with the old", "in with the new")
	docx1.ReplaceFooter("Change This Footer", "new footer")
	docx1.WriteToFile("./new_result_1.docx")

	docx2 := r.Editable()
	docx2.Replace("old_2_1", "new_2_1", -1)
	docx2.Replace("old_2_2", "new_2_2", -1)
	docx2.WriteToFile("./new_result_2.docx")

	r.Close()
}
