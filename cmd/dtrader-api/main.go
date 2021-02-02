package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	r := router.New()

	fmt.Println("Server starting at port 5050")
	log.Fatal(fasthttp.ListenAndServe(":5050", r.Handler))
}
