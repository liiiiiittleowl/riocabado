package riocabado




type PrototypeCore[T any] struct{
	prototype T
}
func(the *PrototypeCore[T]) GetValue() (value T, e error) {
	return the.prototype, nil;
}

func Prototype[T any](value T) Core[T] {
	return &PrototypeCore[T]{
		prototype: value,
	};
}




type FactoryCore[T any] struct{
	function func() (value T, e error)
}
func(the *FactoryCore[T]) GetValue() (value T, e error) {
	return the.function();
}

func Factory[T any](function func() (value T, e error)) Core[T] {
	return &FactoryCore[T]{
		function: function,
	};
}

