package main

import (
	"components"
	_ "components/helper"
	"flag"
)

var (
	apiAddr = flag.String("addr", ":9000", "The API server address.")
)

func main() {
	app := components.NewApp()
	app.ListenAndServe(*apiAddr)
}
