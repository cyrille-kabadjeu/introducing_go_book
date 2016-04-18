package functionUtil

import "testing"

type testPair struct {
	values[]int
	sum int
}

var testData = []testPair{
	{ []int{1,2}, 3 },
	{ []int{1,1,1,1,1,1}, 6 },
	{ []int{-1,1}, 0 },
}

func TestAverage(t *testing.T) {
	for _, pair := range testData {
		v := Add(pair.values...)
		if v != pair.sum {
			t.Error(
				"For", pair.values,
				"expected", pair.sum,
				"got", v,
			)
		}
	}
}