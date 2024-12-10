package autocorrect

import (
	"testing"

	"github.com/longbridge/assert"
)

func Test_fullwidth(t *testing.T) {
	cases := map[string]string{
		"你好,这是一个句子.":        "你好，这是一个句子。",
		"刚刚买了一部iPhone,好开心!": "刚刚买了一部 iPhone，好开心！",
		"蚂蚁集团上市后有多大的上涨空间?":  "蚂蚁集团上市后有多大的上涨空间？",
		"我们需要一位熟悉 JavaScript、HTML5,至少理解一种框架(如 Backbone.js、AngularJS、React 等)的前端开发者.": "我们需要一位熟悉 JavaScript、HTML5，至少理解一种框架 (如 Backbone.js、AngularJS、React 等) 的前端开发者。",
		"蚂蚁疾奔:蚂蚁集团两地上市~全速推进!":                                                        "蚂蚁疾奔：蚂蚁集团两地上市～全速推进！",
		"蚂蚁集团是阿里巴巴(BABA.N)旗下金融科技子公司":                                                 "蚂蚁集团是阿里巴巴 (BABA.N) 旗下金融科技子公司",
		"Dollar的演示 $阿里巴巴.US$ 股票标签":                                                   "Dollar 的演示 $阿里巴巴.US$ 股票标签",
		// https://developer.mozilla.org/en-US/docs/Glossary/Entity
		"确保&quot;&gt;HTML Entity&lt;&quot;的字符&#34;不会被处理&#34; Ruby&amp;Go": "确保&quot;&gt;HTML Entity&lt;&quot;的字符&#34;不会被处理&#34; Ruby&amp;Go",
	}

	for source, exptected := range cases {
		actual := Format(source)
		assert.Equal(t, exptected, actual)
	}
}
