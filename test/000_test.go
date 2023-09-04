package test

import (
	"errors"
	"testing"

	. "riocabado"
)




type DemoCore00[T any] struct{value T}
func(the DemoCore00[T]) GetValue() (value T, e error) {
	return the.value, nil;
}
type DemoCore01[T any] struct{}
func(DemoCore01[T]) GetValue() (value T, e error) {
	return *new(T), errors.New(`Error !!!`);
}


func Is[T any](val any) (is bool) {
	_, ok := val.(T);
	return ok;
}




func Test000(test *testing.T) {
	Register[int](DemoCore00[int]{value: 12});
	Register[string](DemoCore00[string]{value: `@`});

	value00, e := Get[int]();
	if e != nil {
		if _, ok := e.(NoCoreFoundError); ok {
			test.Error(`[ERROR] 出现NoCoreFoundError异常`);
		} else {
			test.Errorf(`[ERROR] %s`, e.Error());
		}
	}
	if value00 != 12 {
		test.Errorf(`[ERROR] 返回值异常 value: %d`, value00);
	}

	value01, e := Get[string]();
	if e != nil {
		if _, ok := e.(NoCoreFoundError); ok {
			test.Error(`[ERROR] 出现NoCoreFoundError异常`);
		} else {
			test.Errorf(`[ERROR] %s`, e.Error());
		}
	}
	if value01 != `@` {
		test.Errorf(`[ERROR] 返回值异常 value: %s`, value01);
	}

	_, e = Get[uint]();
	if e == nil {
		test.Error(`[ERROR] 未出现预期的NoCoreFoundError异常`);
	}
}


func Test001(test *testing.T) {
	RegisterWithName[int](`000`, DemoCore00[int]{value: 12});
	RegisterWithName[int](`001`, DemoCore00[int]{value: 24});

	RegisterWithName[string](`000`, DemoCore00[string]{value: `#`});

	Register[int](DemoCore00[int]{value: 36});


	value00, e := GetWithName[int](`000`);
	if e != nil {
		if _, ok := e.(NoCoreFoundError); ok {
			test.Error(`[ERROR] 出现NoCoreFoundError异常`);
		} else {
			test.Errorf(`[ERROR] %s`, e.Error());
		}
	}
	if value00 != 12 {
		test.Errorf(`[ERROR] 返回值异常 value: %d`, value00);
	}

	value01, e := GetWithName[int](`001`);
	if e != nil {
		if _, ok := e.(NoCoreFoundError); ok {
			test.Error(`[ERROR] 出现NoCoreFoundError异常`);
		} else {
			test.Errorf(`[ERROR] %s`, e.Error());
		}
	}
	if value01 != 24 {
		test.Errorf(`[ERROR] 返回值异常 value: %d`, value01);
	}

	value02, e := GetWithName[string](`000`);
	if e != nil {
		if _, ok := e.(NoCoreFoundError); ok {
			test.Error(`[ERROR] 出现NoCoreFoundError异常`);
		} else {
			test.Errorf(`[ERROR] %s`, e.Error());
		}
	}
	if value02 != `#` {
		test.Errorf(`[ERROR] 返回值异常 value: %s`, value02);
	}

	_, e = GetWithName[string](`001`);
	if e == nil {
		test.Error(`[ERROR] 未出现预期的NoCoreFoundError异常`);
	}

	value03, e := GetWithName[int](DEFAULT_NAME);
	if e != nil {
		if _, ok := e.(NoCoreFoundError); ok {
			test.Error(`[ERROR] 出现NoCoreFoundError异常`);
		} else {
			test.Errorf(`[ERROR] %s`, e.Error());
		}
	}
	if value03 != 36 {
		test.Errorf(`[ERROR] 返回值异常 value: %d`, value03);
	}
}
