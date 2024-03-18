package main

import "testing"

type TestOption struct {
	arr    []string
	expect int
	name   string
}

var testCase = []TestOption{
	{
		arr:    []string{"ab", "bc", "ab", "cd", "bc"},
		expect: 1,
		name:   "Case 1",
	},
	{
		arr:    []string{"ab", "ed", "ab", "cd", "bc"},
		expect: 2,
		name:   "Case 2",
	},
	{
		arr:    []string{"ab", "ed", "bc", "ab", "bc"},
		expect: 1,
		name:   "Case 3",
	},
	{
		arr:    []string{"bc", "ed", "ab", "cd", "bc"},
		expect: 2,
		name:   "Case 4",
	},
	{
		arr:    []string{"ab", "ed", "cd", "bc", "ab"},
		expect: 1,
		name:   "Case 5",
	},
	{
		arr:    []string{"ab", "ab"},
		expect: -1,
		name:   "Case 6",
	},
}

func Test_MinDistance1(t *testing.T) {
	for _, v := range testCase {
		got := MinDistance1(v.arr, "ab", "bc")
		if got != v.expect {
			t.Errorf("expect %d, got %d", v.expect, got)
		}
	}
}

func Test_MinDistance2(t *testing.T) {
	for _, v := range testCase {
		got := MinDistance2(v.arr, "ab", "bc")
		if got != v.expect {
			t.Errorf("%s expect %d, got %d", v.name, v.expect, got)
		}
	}
}
