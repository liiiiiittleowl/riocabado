package riocabado

import "sync"




type SingletonCore[T any] struct{
	value T
	has bool

	locker sync.Locker
	getter func(ctx Context) (value T, e error)
}
func(the *SingletonCore[T]) Example(ctx Context) (value T, e error) {
	if the.has {return the.value, nil;}

	the.locker.Lock();
	defer the.locker.Unlock();

	if the.has {
		return the.value, nil;
	} else {
		value, e = the.getter(ctx);
		if e != nil {return *new(T), e;}

		the.value = value;
		the.has = true;
		return value, nil;
	}
}

func Singleton[T any](
	getter func(ctx Context) (value T, e error),
) Core[T] {
	return &SingletonCore[T]{
		value: *new(T),
		has: false,

		locker: new(sync.Mutex),
		getter: getter,
	};
}




type PrototypeCore[T any] struct{
	prototype T
}
func(the PrototypeCore[T]) Example(ctx Context) (value T, e error) {
	return the.prototype, nil;
}

func Prototype[T any](value T) Core[T] {
	return PrototypeCore[T]{
		prototype: value,
	};
}




type FactoryCore[T any] struct{
	function func() (value T, e error)
}
func(the FactoryCore[T]) Example(ctx Context) (value T, e error) {
	return the.function();
}

func Factory[T any](function func() (value T, e error)) Core[T] {
	return FactoryCore[T]{
		function: function,
	};
}

