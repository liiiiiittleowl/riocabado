package riocabado




const DEFAULT_NAME = `*)>+{%(_,$&}/@$]&!@#"'<.?@[-#%|}"*()_+{|^='~!#>,.?/<\\`;


var pool = map[any]any {};

func Register[T any](core Core[T]) {
	RegisterWithName[T](DEFAULT_NAME, core);
}
func RegisterWithName[T any](name string, core Core[T]) {
	key := struct{name string; zero T} {name: name};
	pool[key] = core;
}

func Get[T any]() (value T, e error) {
	return GetWithName[T](DEFAULT_NAME);
}
func GetWithName[T any](name string) (value T, e error) {
	key := struct{name string; zero T} {name: name};

	prototype, has := pool[key];
	if !has {return *new(T), NoCoreFoundError{};}

	core, _ := prototype.(Core[T]);
	return core.GetValue();
}




type Core[T any] interface{
	GetValue() (value T, e error)
}

type NoCoreFoundError struct{}
func(NoCoreFoundError) Error() (message string) {
	return `No Builder Found`;
}