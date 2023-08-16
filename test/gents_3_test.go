package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zahrulsch/goopy"
)

type Address struct {
	City string `json:"city"`
	Pos  int    `json:"pos_code"`
}

type Person struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address *Address `json:"address"`
}

func TestGenTS3(t *testing.T) {
	var target Person
	res := goopy.GenTS(target)
	expect := `{ id: number; name: string; address: { city: string; pos_code: number } | undefined }`
	assert.Equal(t, expect, res)

	res = goopy.GenTS(&target)
	expect = `{ id: number; name: string; address: { city: string; pos_code: number } | undefined } | undefined`
	assert.Equal(t, expect, res)
}
