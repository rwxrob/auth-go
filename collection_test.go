package auth

import (
	"testing"
)

func TestCollection(t *testing.T) {
	a := new(Data)
	a.Name = "a"
	b := new(Data)
	b.Name = "b"
	c := Collection{a, b}
	_ = c
	// TODO
	//fmt.Println(c)
}
