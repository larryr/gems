package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/larryr/gems/infosrv"
)

var quiet bool

func main() {
	flag.BoolVar(&quiet, "q", false, "quiet: set to stop dumping to log on startup")
	flag.Parse()

	log.Print("infosrv (0.0.0.0:9999)")
	if !quiet {
		infosrv.LogDump(log.Default())
	}
	log.Print("...listening at: (0.0.0.0:9999)")

	infosrv.SetupHandlers("")
	log.Fatal(http.ListenAndServe("0.0.0.0:9999", nil))
}
