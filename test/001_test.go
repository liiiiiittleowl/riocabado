package test

import (
	"testing"
	"time"

	. "riocabado"
)




type Demo00 struct{
	value int64
}




func Test002(test *testing.T) {
	prototype := Demo00{value: time.Now().UnixNano()};
	Register(Prototype[Demo00](prototype));

	value, e := Get[Demo00]();
	if e != nil {
		if _, ok := e.(NoCoreFoundError); ok {
			test.Error(`[ERROR] 出现NoCoreFoundError异常`);
		} else {
			test.Errorf(`[ERROR] %s`, e.Error());
		}
	}
	if value != prototype {
		test.Errorf(`[ERROR] 返回值异常 value: %v`, value);
	}
}




func Test003(test *testing.T) {
	num := 0;

	Register(Factory[struct{}](func() (value struct{}, e error) {
		num++;
		return struct{}{}, nil;
	}));

	prev := num;
	for i, len := 0, 12; i < len; i++ {
		_, e := Get[struct{}]();
		if e != nil {test.Errorf(`[ERROR] %s`, e.Error());}

		if prev == num || num != prev + 1 {
			test.Errorf(`[ERROR] FactoryCore执行异常`);
		}

		prev = num;
	}
}
