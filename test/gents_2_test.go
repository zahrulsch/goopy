package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zahrulsch/goopy"
)

type TcStructChild struct {
	Address string `json:"address"`
}

type TcStruct struct {
	Name string        `json:"name"`
	Age  int           `json:"age"`
	Desc TcStructChild `json:"description"`
}

func TestGenTS2(t *testing.T) {

	var tcs TcStruct

	t.Run("test againts struct", func(t *testing.T) {
		res := escapeChar(goopy.GenTS(tcs, 4))
		expect := escapeChar("{ name: string; age: number; description: { address: string } }")

		assert.Equal(t, expect, res)
	})

	t.Run("test againts struct pointer", func(t *testing.T) {
		res := escapeChar(goopy.GenTS(&tcs, 4))
		expect := escapeChar("{ name: string; age: number; description: { address: string } } | undefined")

		assert.Equal(t, expect, res)
	})
}
