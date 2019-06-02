package sort

// https://eleni.blog/2019/05/11/parallel-test-execution-in-go/
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSortBasic(t *testing.T) {
	numbersToSort := []int32{6, 2, 7, 2}
	sortedNumbers := []int32{2, 2, 6, 7}

	res := CountSort(numbersToSort, 8)
	assert.True(t, assert.ObjectsAreEqualValues(sortedNumbers, res))
}

func TestCountSortTable(t *testing.T) {
	testCases := []struct {
		Name     string
		Numbers  []int32
		Expected []int32
		Range    int32
	}{
		{
			Name:     "All the numbers in the range [1-9]",
			Numbers:  []int32{4, 1, 9, 6, 3, 8, 7, 2, 5},
			Expected: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Range:    int32(10),
		},
		{
			Name:     "3 numbers in the range [1-9]",
			Numbers:  []int32{4, 1, 9},
			Expected: []int32{1, 4, 9},
			Range:    int32(10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			res := CountSort(tc.Numbers, tc.Range)
			assert.True(t, assert.ObjectsAreEqualValues(tc.Expected, res))
		})
	}
}

func TestCountSortParallel(t *testing.T) {
	testCases := []struct {
		Name     string
		Numbers  []int32
		Expected []int32
		Range    int32
	}{
		{
			Name:     "All the numbers in the range [1-9]",
			Numbers:  []int32{4, 1, 9, 6, 3, 8, 7, 2, 5},
			Expected: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Range:    int32(10),
		},
		{
			Name:     "3 numbers in the range [1-9]",
			Numbers:  []int32{4, 1, 9},
			Expected: []int32{1, 4, 9},
			Range:    int32(10),
		},
	}

	for _, tc := range testCases {
		tc := tc // We run our tests twice one with this line & one without
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			res := CountSort(tc.Numbers, tc.Range)
			assert.True(t, assert.ObjectsAreEqualValues(tc.Expected, res))
		})
	}
}

func TestCountSort(t *testing.T) {
	testCases := []struct {
		Name     string
		Numbers  []int32
		Expected []int32
		Range    int32
	}{
		{
			Name:     "All the numbers in the range [1-9]",
			Numbers:  []int32{4, 1, 9, 6, 3, 8, 7, 2, 5},
			Expected: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Range:    int32(10),
		},
		{
			Name:     "3 numbers in the range [1-9]",
			Numbers:  []int32{4, 1, 9},
			Expected: []int32{1, 4, 9},
			Range:    int32(10),
		},
		{
			Name:     "Repeated numbers in the range [1-9]",
			Numbers:  []int32{4, 1, 9, 5, 4, 1, 4, 9},
			Expected: []int32{1, 1, 4, 4, 4, 5, 9, 9},
			Range:    int32(10),
		},
		{
			Name:     "Repeated numbers in the range [1-5]",
			Numbers:  []int32{4, 1, 4, 5},
			Expected: []int32{1, 4, 4, 5},
			Range:    int32(10),
		},
		{
			Name:     "Repeated numbers in the range [1-3]",
			Numbers:  []int32{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3},
			Expected: []int32{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3},
			Range:    int32(10),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			res := CountSort(tc.Numbers, tc.Range)
			assert.True(t, assert.ObjectsAreEqualValues(tc.Expected, res))
		})
	}
}
