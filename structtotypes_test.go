package goopy

import (
	"os"
	"testing"
	"time"
)

type Access struct {
	AccessType string `json:"access_type"`
}

type OtherInfo struct {
	Wa string  `json:"whatsapp"`
	V  *string `json:"viagra"`
	A  Access  `json:"access_"`
}

type Example struct {
	Other OtherInfo             `json:"other"`
	Name  string                `json:"name"`
	Age   *int                  `json:"age"`
	Adul  bool                  `json:"adult"`
	J     map[string]*OtherInfo `json:"js"`
	Ks    []string              `json:"ks"`
	Is    []*string             `json:"is"`
	Date  time.Time             `json:"date" alter_type:"string"`
	// Access interface{}        `json:"access_type"`
}

func TestStructToTypes(t *testing.T) {
	out := "example.ts"
	var i = Example{}

	_, r := StructToTYPES(i)

	f, _ := os.Create(out)
	defer f.Close()

	r = "export type Example = " + r

	f.Write([]byte(r))
}
