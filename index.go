package riocabado




var _Pool = map[any]any{};

func Register[T any](core Core[T]) {
	if core == nil {return;}

	var key T;
	_Pool[key] = core;
}
func Get[T any]() (value T, e error) {
	return DependOn[T](Context{records: map[any]struct{}{}})
}
func DependOn[T any](ctx Context) (value T, e error) {
	var key T;
	if _, has := ctx.records[key]; has {
		return *new(T), CircularDependencyError{};
	}

	prototype, has := _Pool[key];
	if !has {return *new(T), NoCoreFoundError{};}

	ctx.records[key] = struct{}{};
	defer delete(ctx.records, key);

	core, _ := prototype.(Core[T]);
	return core.Example(ctx);
}




type Context struct{
	records map[any]struct{}
}


type Core[T any] interface{
	Example(ctx Context) (value T, e error)
}


type CircularDependencyError struct{}
func(CircularDependencyError) Error() (message string) {
	return "circular dependency";
}

type NoCoreFoundError struct{}
func(NoCoreFoundError) Error() (message string) {
	return "no core found";
}
