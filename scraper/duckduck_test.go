package scraper

import (
	"embed"
	"testing"

	"github.com/ayuxsys/ddg/must"
)

//go:embed testdata/headless_resp.html
var testData embed.FS

func TestParseDDGHeadless(t *testing.T) {
	html := must.Eval(testData.ReadFile("testdata/headless_resp.html"))
	t.Log(parseDDGHeadless(string(html)))
}

// func TestParseDDG(t *testing.T) {
// 	html := `
// 	<h2 class="result__title">
// 		<a rel="nofollow" class="result__a"
// 		   href="https://en.wikipedia.org/wiki/Coffee">
// 			Coffee - Wikipedia
// 		</a>
// 	</h2>`
//
// 	res := parseDDG(html)
// 	t.Log(res)
// }
