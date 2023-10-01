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

var error00 = errors.New("@");
type DemoCore03 struct{}
func(DemoCore03) Example(ctx Context) (value uint8, e error) {
	return *new(uint8), error00;
}


func TestFunction(test *testing.T) {
	Register[uint](DemoCore00{});
	value00, e := Get[uint]();
	if e != nil {test.Errorf("出现错误 message: %v\n", e);}
	if value00 != 12 {test.Errorf("意外的返回值 value: %v\n", value00)}

	Register[string](DemoCore01{});
	value01, e := Get[string]();
	if e != nil {test.Errorf("出现错误 message: %v\n", e);}
	if value01 != "@" {test.Errorf("意外的返回值 value: %v\n", value00)}

	Register[Name](DemoCore02{});
	value02, e := Get[Name]();
	if e != nil {test.Errorf("出现错误 message: %v\n", e);}
	if value02 != "riocabado" {test.Errorf("意外的返回值 value: %v\n", value00)}

	Register[uint8](DemoCore03{});
	var _, err = Get[uint8]();
	if err == nil {test.Errorf("未返回预定错误\n");}
	if err != error00 {test.Errorf("出现意外的错误 message: %v\n", err);}


	value00, e = Get[uint]();
	if e != nil {test.Errorf("出现错误 message: %v\n", e);}
	if value00 != 12 {test.Errorf("意外的返回值 value: %v\n", value00)}
}
