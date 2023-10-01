package test

import (
	"errors"
	"testing"

	. "riocabado"
)



type Name string

type DemoCore00 struct{}
func(DemoCore00) Example(ctx Context) (value uint, e error) {
	return 12, nil;
}
type DemoCore01 struct{}
func(DemoCore01) Example(ctx Context) (value string, e error) {
	return "@", nil;
}
type DemoCore02 struct{}
func(DemoCore02) Example(ctx Context) (value Name, e error) {
	return "riocabado", nil;
}
type DemoCore03 struct{}
func(DemoCore03) Example(ctx Context) (value Name, e error) {
	return "", errors.New("@");
}


func TestFunction(test *testing.T) {

}
