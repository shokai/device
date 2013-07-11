package app

import (
	"github.com/hoisie/web"
	"log"
	"net/http/httputil"
	"os"
	"strings"
)

const host = "0.0.0.0:9999"

func index(c *web.Context, val string) {
	type data struct {
		Name string
	}
	d := data{val}

	context := Context{*c}
	context.WriteTemplate("templates/index.template", d)
}

func one(c *web.Context, val string) {
	dump, _ := httputil.DumpRequest(c.Request, false)
	c.Server.Logger.Println(string(dump))

	type data struct {
		Title string
		Host  string
	}
	d := data{
		"<script>alert('you have been pwned')</script>",
		host,
	}

	context := Context{*c}
	buf := context.RenderTemplate("templates/one.template", d)
	redirectTo := "data:text/html;charset=UTF-8," + strings.Replace(buf.String(), "\n", "", -1)
	c.Redirect(302, redirectTo)
}

func NewServer() *web.Server {
	server := web.NewServer()
	server.Get("/apps/one/(.*)", one)
	server.Get("/(.*)", index)
	server.SetLogger(log.New(os.Stdout, "app ", log.Ldate|log.Ltime|log.Lshortfile))
	return server
}
