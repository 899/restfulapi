package helper

import (
	"components"
	"time"
)

type HelperAPI struct{}

func init() {
	components.HelperAPI = &HelperAPI{}
}

func (helper *HelperAPI) ComponentInit() {
	components.HandleFunc("test", helper.Test)
}

func (helper *HelperAPI) Test(ctx *components.Context) interface{} {
	var req struct {
		Abc int64
	}
	ctx.Request(&req)
	return map[string]interface{}{
		"now":   time.Now().UTC().Unix(),
		"param": req.Abc,
	}
}
func (helper *HelperAPI) Name() string {
	return "helper"
}
