package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/larryr/gems/showchrome"
)

var quiet bool

func main() {
	flag.BoolVar(&quiet, "q", false, "quiet: set to stop dumping to log on startup")
	flag.Parse()

	log.Print("showchrome (0.0.0.0:7002)")
	if !quiet {
		showchrome.LogDump(log.Default())
	}
	log.Print("...listening at: (0.0.0.0:7002)")

	showchrome.SetupHandlers("")
	log.Fatal(http.ListenAndServe("0.0.0.0:7002", nil))
}
