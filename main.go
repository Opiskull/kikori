package main

import (
	"github.com/opiskull/kikori/logserver"
	"github.com/opiskull/kikori/webserver"
)

func main() {
	server := logserver.CreateLogServer()
	server.Wait()

	web := webserver.CreateWebServer()
	web.Start(":1323")
}
