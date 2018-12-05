package duplicates_test

import "testing"
import dups "github.com/stirt/duplicates/pkg/duplicates"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type testCase struct {
	input  []int
	output []int
}

func createTestCases() map[int]testCase {
	cases := make(map[int]testCase)
	cases[0] = testCase{input: []int{1, 2, 3, 4, 4, 5}, output: []int{4}}
	cases[1] = testCase{input: []int{2, 3, 4, 5, 6, 6, 6, 7, 7, 8, 9, 10}, output: []int{6, 7}}
	return cases
}

func TestFindDuplicates(t *testing.T) {
	testCases := createTestCases()
	for _, v := range testCases {
		found := dups.FindDuplicates(v.input)
		if !equal(found, v.output) {
			t.Fatalf("Unexpected result, expecting %v but got %v", v.output, found)
		}
	}
}
