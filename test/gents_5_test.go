package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zahrulsch/goopy"
)

func TestGenTS5(t *testing.T) {
	tcs := []Tc{
		{
			T:      map[string]string{},
			Expect: "Record<string, string>",
		},
		{
			T:      map[string]*string{},
			Expect: "Record<string, string | undefined>",
		},
		{
			T:      map[int32]*string{},
			Expect: "Record<number, string | undefined>",
		},
		{
			T:      map[int32]Address{},
			Expect: "Record<number, { city: string; pos_code: number }>",
		},
		{
			T:      map[int32]*Address{},
			Expect: "Record<number, { city: string; pos_code: number } | undefined>",
		},
	}

	for _, tc := range tcs {
		testName := fmt.Sprintf("test againt %T", tc.T)
		t.Run(testName, func(t *testing.T) {
			res := escapeChar(goopy.GenTS(tc.T, 4))
			assert.Equal(t, escapeChar(tc.Expect), res)
		})
	}
}
