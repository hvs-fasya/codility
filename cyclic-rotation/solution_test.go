package cyclic_rotation

import (
	"reflect"
	"testing"
)

type Example struct {
	Times int
	Arr   []int
}

func Test(t *testing.T) {
	samples := []struct {
		Sample   Example
		Expected []int
	}{
		{Example{Times: 1, Arr: nil}, nil},
		{Example{Times: 3, Arr: []int{}}, []int{}},
		{Example{Times: 3, Arr: []int{3, 8, 9, 7, 6}}, []int{9, 7, 6, 3, 8}},
		{Example{Times: 1, Arr: []int{0, 0, 0}}, []int{0, 0, 0}},
		{Example{Times: 0, Arr: []int{1, 2, 3}}, []int{1, 2, 3}},
		{Example{Times: 4, Arr: []int{1, 2, 3}}, []int{3, 1, 2}},
	}
	t.Run("CyclicRotationSolution", func(t *testing.T) {
		for _, sample := range samples {
			res := Solution(sample.Sample.Times, sample.Sample.Arr)
			if !reflect.DeepEqual(res, sample.Expected) {
				t.Errorf("for sample K: %d Arr: %v expected %v; got %v", sample.Sample.Times, sample.Sample.Arr, sample.Expected, res)
			}
		}
	})
}
