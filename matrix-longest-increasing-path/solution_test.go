package matrix_longest_increasing_path

import (
	"reflect"
	"testing"
)

type Example struct {
	Matrix [][]int
}

func Test(t *testing.T) {
	samples := []struct {
		Sample   Example
		Expected int
	}{
		{Sample: Example{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}}, Expected: 4},
		{Sample: Example{[][]int{{3, 4, 5}, {9, 6, 6}, {8, 7, 1}}}, Expected: 6},
		{Sample: Example{[][]int{{1}}}, Expected: 1},
		{Sample: Example{[][]int{{}}}, Expected: 0},
	}

	t.Run("MatrixLongestIncreasingPathSolution", func(t *testing.T) {
		for _, sample := range samples {
			res := Solution(sample.Sample.Matrix)
			if !reflect.DeepEqual(res, sample.Expected) {
				t.Errorf("for sample: matrix: %v expected %d; got %d", sample.Sample.Matrix, sample.Expected, res)
			}
		}
	})
}
