// Code generated by Wire. DO NOT EDIT.

//go:generate go tool code.cestus.io/tools/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFooBar() FooBar {
	bar := provideBar()
	fooBar := provideFooBar(bar, bar)
	return fooBar
}
