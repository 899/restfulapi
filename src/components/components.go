package components

type APIMethods map[string]func(*Context) interface{}

type IHelper interface {
	ComponentInit()
	Name() string
}

var HelperAPI IHelper

func Start() {
	HelperAPI.ComponentInit()
}
