package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/musaubrian/maintenance/gno"
	"github.com/musaubrian/maintenance/view"
)

func main() {
	lg := gno.Log

	var port string
	flag.StringVar(&port, "p", "", "port to expose the page")
	flag.Parse()
	if len(port) < 1 {
		lg(gno.ERROR, "Missing port")
	}

	port = ":" + port
	main := view.Main()
	http.Handle("/", templ.Handler(main))
	lg(gno.INFO, "Maintenance page running at port "+port)
	log.Fatal(http.ListenAndServe(port, nil))
}
