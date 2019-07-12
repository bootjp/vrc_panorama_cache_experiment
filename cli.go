package main

import (
	"log"
	"time"

	"github.com/kataras/iris/cache"

	"github.com/kataras/iris"
)

func main() {

	app := iris.Default()

	// response content and no-cache
	// It did not go well as a precedent case  https://twitter.com/27Cobalter/status/1149185587594723328
	app.Get("/no-cache", func(ctx iris.Context) {
		ctx.Header("cache-control", "no-store")
		// TODO : add response image
	})

	// app.he
	app.Get("/redirect", func(ctx iris.Context) {
		ctx.Header(cache.CacheControlHeaderValue, "no-store")
		t := time.Now()
		ctx.Redirect("/img/t?="+t.UTC().String(), 307)
	})

	// always return 404 and contents.
	app.Get("/404", func(ctx iris.Context) {
		ctx.Header(cache.CacheControlHeaderValue, "no-store")
		ctx.StatusCode(404)
		// TODO : add response image
	})

	if err := app.Run(iris.Addr(":8080")); err != nil {
		log.Fatal(err)
	}
}
