package nwunsch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNeedlemanWunsch(t *testing.T) {
	tests := []struct {
		seq1, seq2           string
		match, mismatch, gap int
		score                int
		align1, align2       string
	}{
		{
			seq1:     "GATTACA",
			seq2:     "GCATGCU",
			match:    1,
			mismatch: -1,
			gap:      -2,
			score:    -1,
			align1:   "GATTACA",
			align2:   "GCATGCU",
		},
		{
			seq1:     "AGT",
			seq2:     "AGCT",
			match:    1,
			mismatch: -1,
			gap:      -1,
			score:    2,
			align1:   "AG-T",
			align2:   "AGCT",
		},
		// more test cases...
	}

	for _, test := range tests {
		score, align1, align2 := NeedlemanWunsch(test.seq1, test.seq2, test.match, test.mismatch, test.gap)

		assert.Equal(t, test.score, score, "unexpected score for sequences %q and %q", test.seq1, test.seq2)
		assert.Equal(t, test.align1, align1, "unexpected alignment 1 for sequences %q and %q", test.seq1, test.seq2)
		assert.Equal(t, test.align2, align2, "unexpected alignment 2 for sequences %q and %q", test.seq1, test.seq2)
	}
}
