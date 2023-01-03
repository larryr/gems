package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/larryr/gems/simp"
)

var quiet bool

func main() {
	flag.BoolVar(&quiet, "q", false, "quiet: set to stop dumping to log on startup")
	flag.Parse()

	log.Print("simp (0.0.0.0:7001)")
	if !quiet {
		simp.LogDump(log.Default())
	}
	log.Print("...listening at: (0.0.0.0:7001)")

	simp.SetupHandlers("")
	log.Fatal(http.ListenAndServe("0.0.0.0:7001", nil))
}
