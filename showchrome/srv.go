package showchrome

import (
	"fmt"
	"github.com/larryr/gems/infosrv/ienv"
	"log"
	"net/http"
	"strings"
)

type fctTyp struct {
	name string
	fct  func(w http.ResponseWriter, _ *http.Request)
}

var fctnames = ""
var ch *ChromeCtl

func SetupHandlers(prefix string) {
	ch = &ChromeCtl{}

	fcts := []fctTyp{
		{"/", handler},
		{"/chrome/", chromeHandler},
		{"/info/env/", ienv.Handler()},
	}
	for _, fct := range fcts {
		http.HandleFunc(prefix+fct.name, fct.fct)
		fctnames = fmt.Sprintf("%s%s\n", fctnames, fct.name)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello larry: %s\n%s", r.URL.Path[1:], fctnames)
}

func LogDump(l *log.Logger) {
	strDump(l, ienv.EnvInfo())
}

func strDump(l *log.Logger, vals []string) {
	for _, v := range vals {
		l.Printf("%s", v)
	}
}

func chromeHandler(w http.ResponseWriter, r *http.Request) {
	cmdargs := cmdArgs(r.URL.Path[1:])
	if len(cmdargs) < 2 {
		fmt.Fprintf(w, "unknown command: %s\n", r.URL.Path)
		return
	}
	cmd := cmdargs[1]
	args := cmdargs[2:]
	result := "bad command"
	switch cmd {
	case "launch":
		result = ch.Launch(args)
	}
	fmt.Fprintf(w, "%s", result)
}

func cmdArgs(path string) []string {
	a := strings.Split(path, "/")
	return a
}
