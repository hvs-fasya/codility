package binary_gap

import "testing"

func Test(t *testing.T) {
	samples := []struct {
		Sample   int
		Expected int
	}{
		{9, 2},
		{32, 0},
		{529, 4},
		{20, 1},
		{15, 0},
	}
	t.Run("BinaryGapSolution", func(t *testing.T) {
		for _, sample := range samples {
			res := Solution(sample.Sample)
			if sample.Expected != res {
				t.Errorf("for sample %d expected %d; got %d", sample.Sample, sample.Expected, res)
			}
		}
	})
}
