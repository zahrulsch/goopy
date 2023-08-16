package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zahrulsch/goopy"
)

type I struct {
	Key interface{} `json:"key"`
}

func TestGenTS6(t *testing.T) {
	tcs := []Tc{
		{
			T:      I{Key: ""},
			Expect: "{ key: string }",
		},
		{
			T:      I{Key: 0},
			Expect: "{ key: number }",
		},
		{
			T:      I{Key: false},
			Expect: "{ key: boolean }",
		},
		{
			T:      I{Key: Address{}},
			Expect: "{ key: { city: string; pos_code: number } }",
		},
		{
			T:      I{Key: &Address{}},
			Expect: "{ key: { city: string; pos_code: number } | undefined }",
		},
		{
			T:      I{Key: []string{}},
			Expect: "{ key: Array<string> }",
		},
		{
			T:      I{Key: []*string{}},
			Expect: "{ key: Array<string | undefined> }",
		},
		{
			T:      I{Key: []int{}},
			Expect: "{ key: Array<number> }",
		},
		{
			T:      I{Key: []*int{}},
			Expect: "{ key: Array<number | undefined> }",
		},
		{
			T:      I{Key: map[string]string{}},
			Expect: "{ key: { [key: string]: string } }",
		},
		{
			T:      I{Key: map[int]string{}},
			Expect: "{ key: { [key: number]: string } }",
		},
		{
			T:      I{Key: map[int]*string{}},
			Expect: "{ key: { [key: number]: string | undefined } }",
		},
		{
			T:      I{Key: map[int][]string{}},
			Expect: "{ key: { [key: number]: Array<string> } }",
		},
	}

	for _, tc := range tcs {
		testName := fmt.Sprintf("test againts %#v", tc.T)
		t.Run(testName, func(t *testing.T) {
			res := goopy.GenTS(tc.T)
			assert.Equal(t, tc.Expect, res)
		})
	}
}
