package autocorrect

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/longbridge/assert"
)

func readFile(filename string) (out string) {
	data, err := ioutil.ReadFile(fmt.Sprintf("./_fixtures/%s", filename))
	if err != nil {
		panic(err)
	}
	return string(data)
}

func BenchmarkFormatHTML(b *testing.B) {
	raw := readFile("example.txt")
	for i := 0; i < b.N; i++ {
		// about 1.07ms/op
		FormatHTML(raw)
	}
}

func TestFormatHTMLWithFixtuires(t *testing.T) {
	expected := readFile("example.expected.txt")
	out, err := FormatHTML(readFile("example.txt"))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)
	assert.EqualHTML(t, expected, out)
}

func TestFormatHTMLWithSameTextInAttribute(t *testing.T) {
	html := `<p data-value="每股均价9.582港元，增持270.3万股"><script>var value = "三生制药获JP Morgan以";</script><pre>每股均价9.582港元</pre><textarea>文本框test</textarea><style>//增持270.3万股</style>三生制药获JP Morgan以<i>每股均价9.582港元，增持270.3万股</i>增持</p>`
	expected := `<p data-value="每股均价9.582港元，增持270.3万股"><script>var value = "三生制药获JP Morgan以";</script><pre>每股均价9.582港元</pre><textarea>文本框test</textarea><style>//增持270.3万股</style>三生制药获 JP Morgan 以<i>每股均价 9.582 港元，增持 270.3 万股</i>增持</p>`
	out, err := FormatHTML(html)
	if err != nil {
		t.Error(err)
	}
	assert.EqualHTML(t, expected, out)
}

func TestFormatHTMLWithEscapedHTML(t *testing.T) {
	html := `<p>据2019年12月27日，三生制药获JP Morgan Chase &amp; Co.每股均价9.582港元，增持270.3万股</p>`
	expected := `<p>据 2019 年 12 月 27 日，三生制药获 JP Morgan Chase &amp; Co.每股均价 9.582 港元，增持 270.3 万股</p>`
	out, err := FormatHTML(html)
	if err != nil {
		t.Error(err)
	}
	assert.EqualHTML(t, expected, out)

	html = `<p>据2019年12月27日，三生制药获JP Morgan Chase & Co.每股均价9.582港元，增持270.3万股</p>`
	expected = `<p>据 2019 年 12 月 27 日，三生制药获 JP Morgan Chase & Co.每股均价 9.582 港元，增持 270.3 万股</p>`
	out, err = FormatHTML(html)
	if err != nil {
		t.Error(err)
	}
	assert.EqualHTML(t, expected, out)
}

func TestFormatHTML_halfwidth(t *testing.T) {
	html := `<p>自动转换全角“字符、数字”：我们将在（１６：３２）出发去ＣＢＤ中心。</p>`

	out, err := FormatHTML(html)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "<p>自动转换全角 “字符、数字”：我们将在（16:32）出发去 CBD 中心。</p>", out)
}

func TestUnformatHTML(t *testing.T) {
	raw := "<p>Hello world this is english.</p><p><strong>2018 至 2019 财年</strong>，印度电力部门总进口额为 7100 亿卢比（约合 672 亿人民币）其中 2100 亿卢比来自中国</p><p>占比 29.6%，这意味着中国是印度电力设备的国外主要供应商。</p>"
	out, err := UnformatHTML(raw)
	assert.NoError(t, err)
	assert.EqualHTML(t, "<p>Hello world this is english.</p><p><strong>2018至2019财年</strong>，印度电力部门总进口额为7100亿卢比（约合672亿人民币）其中2100亿卢比来自中国</p><p>占比29.6%，这意味着中国是印度电力设备的国外主要供应商。</p>", out)

}

func BenchmarkFormatHTML_large(b *testing.B) {
	raw := readFile("large.html")

	for i := 0; i < b.N; i++ {
		// about 1.07ms/op
		FormatHTML(raw)
	}
}
