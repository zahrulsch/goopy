package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zahrulsch/goopy"
)

type Tc struct {
	T      interface{}
	Expect string
}

func TestGenTS1(t *testing.T) {
	tcs := []Tc{
		{
			T:      int(0),
			Expect: "number",
		},
		{
			T:      int32(0),
			Expect: "number",
		},
		{
			T:      int64(0),
			Expect: "number",
		},
		{
			T:      uint(0),
			Expect: "number",
		},
		{
			T:      uint32(0),
			Expect: "number",
		},
		{
			T:      uint64(0),
			Expect: "number",
		},
		{
			T:      float32(0),
			Expect: "number",
		},
		{
			T:      float64(0),
			Expect: "number",
		},
		{
			T:      string("test"),
			Expect: "string",
		},
		{
			T:      false,
			Expect: "boolean",
		},
	}

	for _, tc := range tcs {
		testName := fmt.Sprintf("test againts value concrete %T", tc.T)
		t.Run(testName, func(t *testing.T) {
			res := goopy.GenTS(tc.T)
			assert.Equal(t, tc.Expect, res)
		})
	}

	t.Run("test againts pointer string", func(t *testing.T) {
		var tar *string
		res := goopy.GenTS(tar)
		assert.Equal(t, "string | undefined", res)
	})

	t.Run("test againts pointer int", func(t *testing.T) {
		var tar *int
		res := goopy.GenTS(tar)
		assert.Equal(t, "number | undefined", res)
	})

	t.Run("test againts pointer bool", func(t *testing.T) {
		var tar *bool
		res := goopy.GenTS(tar)
		assert.Equal(t, "boolean | undefined", res)
	})
}
