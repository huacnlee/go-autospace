package autocorrect

import (
	"strings"
	"testing"

	"github.com/longbridge/assert"
)

func TestUnformat(t *testing.T) {
	assert.Equal(t, " Hello world ", Unformat(" Hello world "))
	assert.Equal(t, "100中文", Unformat("100 中文"))
	assert.Equal(t, "中文100", Unformat("中文 100"))

	raw := "据港交所最新权益披露资料显示，2019 年 12 月 27 日，三生制药获 JP Morgan Chase & Co.每股均价 9.582 港元，增持 270.3 万股，总价约 2590 万港元。"
	expected := "据港交所最新权益披露资料显示，2019年12月27日，三生制药获JP Morgan Chase & Co.每股均价9.582港元，增持270.3万股，总价约2590万港元。"
	assert.Equal(t, expected, Unformat(raw))
}

type customUnformat struct{}

func (c customUnformat) Unformat(text string) string {
	return strings.ReplaceAll(text, "BBBB", "AAAA")
}

func TestUnformatWithOptions(t *testing.T) {
	assert.Equal(t, "增持270.3万股AAAA", Unformat("增持270.3万股BBBB", customUnformat{}))
}

func Test_UnformatHTMLWithOptions(t *testing.T) {
	out, err := UnformatHTML("<p>测试 options 你好 BBBB</p>", customUnformat{})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "<p>测试options你好AAAA</p>", out)
}
