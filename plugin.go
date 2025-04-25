package myrrmiddlewate

import (
	"fmt"
	"net/http"
)

const PluginName = "myrrmiddlewate"

type Plugin struct{}

func (p *Plugin) Init() error {
	return nil
}

func (p *Plugin) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware test\n")
		next.ServeHTTP(w, r)
	})
}

func (p *Plugin) Name() string {
	return PluginName
}
