package main

import (
	"net/http"

	"github.com/hyperioxx/frontman/config"
	"github.com/hyperioxx/frontman/service"
)


type FrontmanPlugin struct {}

func (p *FrontmanPlugin) Name() string {
    return "Header Change Plugin"
}

func (p *FrontmanPlugin) PreRequest(req *http.Request, sr service.ServiceRegistry, cfg config.Config) error {
    req.Header.Add("MyHeader", "TEST1")
    return nil
}

func (p *FrontmanPlugin) PostResponse(resp *http.Response, sr service.ServiceRegistry, cfg config.Config) error {
    resp.Header.Add("MyHeader", "TEST1")
    return nil
}

func (p *FrontmanPlugin) Close() error {
    return nil
}