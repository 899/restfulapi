package components

import (
	"crypto"
	"net/http"

	"fmt"

	"runtime"

	"log"

	"github.com/funny/jsonapi"
)

type App struct {
	jsonapi *jsonapi.API
}

var (
	DefaultApp *App
)

func NewApp() *App {
	api := jsonapi.New(crypto.SHA256, jsonapi.StdLogger)

	DefaultApp = &App{
		jsonapi: api,
	}

	return DefaultApp
}

type Context struct {
	context *jsonapi.Context
}

func (app *App) HandleFunc(path string, method func(*Context) interface{}) {
	app.jsonapi.HandleFunc(path, func(ctx *jsonapi.Context) interface{} {

		myctx := &Context{context: ctx}
		return method(myctx)
	})
}

func (app *App) ListenAndServe(addr string) {
	Start()

	http.ListenAndServe(addr, app.jsonapi)
}

func HandleFunc(path string, method func(*Context) interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	log.Fatal(f.Name())
	DefaultApp.HandleFunc(path, method)
}

func (ctx *Context) Request(req interface{}) {
	ctx.context.Request(req)
}

func GetPath(module string, method string) string {
	return fmt.Sprintf("/%s/%s", module, method)
}
