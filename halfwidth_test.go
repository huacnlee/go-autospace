package autocorrect

import (
	"testing"

	"github.com/longbridge/assert"
)

func Test_halfwidth(t *testing.T) {
	source := "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ１２３４５６７８９０"
	assert.Equal(t, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", halfwidth(source))
	assert.Equal(t, "他说：我们将在16:32分出发去CBD中心。", halfwidth("他说：我们将在１６：３２分出发去ＣＢＤ中心。"))
	// Fullwidth space
	assert.Equal(t, "ジョイフル－後場売り気配 200 店舗を閉鎖へ 7 月以降、不採算店中心に", halfwidth("ジョイフル－後場売り気配　200　店舗を閉鎖へ　7 月以降、不採算店中心に"))
	// Fullwidth Numbers
	assert.Equal(t, "0 1 2 3 4 5 6 7 8 9", halfwidth("0 1 2 3 4 5 6 7 8 9"))
}

func Benchmark_halfwidth(b *testing.B) {
	source := "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ１２３４５６７８９０"
	for i := 0; i < b.N; i++ {
		// about 0.003ms/op
		halfwidth(source)
	}
}
