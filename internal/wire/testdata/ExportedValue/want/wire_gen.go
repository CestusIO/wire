// Code generated by Wire. DO NOT EDIT.

//go:generate go tool code.cestus.io/tools/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"example.com/bar"
)

// Injectors from wire.go:

func injectedMessage() string {
	string2 := _wireStringValue
	return string2
}

var (
	_wireStringValue = bar.PublicMsg
)
