// Code generated by Wire. DO NOT EDIT.

//go:generate go tool code.cestus.io/tools/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func InitializeFooStore() FooStorer {
	foo := _wireFooValue
	store := NewFooStore(foo)
	return store
}

var (
	_wireFooValue = Foo("foo hello value")
)
