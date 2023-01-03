package simp

import (
	"fmt"
	"github.com/larryr/gems/infosrv/ienv"
	"log"
	"net/http"
)

type fctTyp struct {
	name string
	fct  func(w http.ResponseWriter, _ *http.Request)
}

var fctnames = ""

func SetupHandlers(prefix string) {

	fcts := []fctTyp{
		{"/", handler},
		{"/info/", handler},
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
