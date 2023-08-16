package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zahrulsch/goopy"
)

func TestGenTS4(t *testing.T) {
	tcs := []Tc{
		{
			T:      []string{},
			Expect: "Array<string>",
		},
		{
			T:      &[]string{},
			Expect: "Array<string> | undefined",
		},
		{
			T:      []*string{},
			Expect: "Array<string | undefined>",
		},
		{
			T:      []int32{},
			Expect: "Array<number>",
		},
		{
			T:      &[]int32{},
			Expect: "Array<number> | undefined",
		},
		{
			T:      []*int32{},
			Expect: "Array<number | undefined>",
		},
		{
			T:      []bool{},
			Expect: "Array<boolean>",
		},
		{
			T:      &[]bool{},
			Expect: "Array<boolean> | undefined",
		},
		{
			T:      []*bool{},
			Expect: "Array<boolean | undefined>",
		},
		{
			T:      &[]*bool{},
			Expect: "Array<boolean | undefined> | undefined",
		},
	}

	for _, tc := range tcs {
		t.Run("test againts slice", func(t *testing.T) {
			res := goopy.GenTS(tc.T)
			assert.Equal(t, tc.Expect, res)
		})
	}

	t.Run("test againts slice of struct", func(t *testing.T) {
		var address []Address
		res := goopy.GenTS(address)
		expect := "Array<{ city: string; pos_code: number }>"

		assert.Equal(t, expect, res)

		var persons []Person
		res = goopy.GenTS(persons)
		expect = "Array<{ id: number; name: string; address: { city: string; pos_code: number } | undefined }>"

		assert.Equal(t, expect, res)
	})
}
