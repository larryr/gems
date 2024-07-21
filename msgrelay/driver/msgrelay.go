package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/larryr/gems/msgrelay"
)

var quiet bool
var port = 9901
var name = "msgrelay"

func main() {
	flag.BoolVar(&quiet, "q", false, "quiet: set to stop dumping to log on startup")
	flag.Parse()

	log.Printf("%v (0.0.0.0:%v)", name, port)
	if !quiet {
		msgrelay.LogDump(log.Default())
	}
	log.Printf("...listening at: (0.0.0.0:%v)", port)

	msgrelay.SetupHandlers("")
	addr := fmt.Sprintf("0.0.0.0:%v", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
