package main

import (
	"os"

	"github.com/musaubrian/maintenance/gno"
)

func main() {
	if len(os.Args) != 2 {
		gno.Log(gno.ERROR, "Expected one argument (port)")
	}
	port := os.Args[1]

	g := gno.New()
	g.BootstrapBuild("build", "maintenance", ".")
	// g.AddCommand("templ", "generate")
	g.AddCommand("./build/maintenance", "-p", port)
	g.Build()
	g.RunCommandsSync()
}
