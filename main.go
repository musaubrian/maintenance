package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/a-h/templ"
	g "github.com/musaubrian/maintenance/gno"
	"github.com/musaubrian/maintenance/view"
)

func main() {

	var port string
	flag.StringVar(&port, "p", "", "port to expose the page")
	flag.Parse()
	if len(port) < 1 {
		g.Log(g.ERROR, "Missing port")
	}
	err := loadEnv()
	if err != nil {
		g.Log(g.ERROR, "Could not set env from `.env`, try running with EMAIL=random@domain <program>")
	}

	email := os.Getenv("EMAIL")
	port = ":" + port
	main := view.Main(email)
	http.Handle("/", templ.Handler(main))
	g.Log(g.INFO, "Maintenance page running at port "+port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func loadEnv() error {
	b, err := os.ReadFile(".env")
	if err != nil {
		return err
	}
	res := os.Getenv("EMAIL")
	if len(res) > 1 {
		g.Log(g.INFO, "EMAIL exists in environment, skipping loading `.env`")
		return nil
	}
	g.Log(g.INFO, "Loaded `.env`")
	err = os.Setenv("EMAIL", strings.Split(string(b), "=")[1])
	if err != nil {
		return err
	}
	return nil
}
