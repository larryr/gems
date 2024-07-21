package msgrelay

import (
	"fmt"
	"log"
	"net/http"

	"github.com/larryr/gems/infosrv/ienv"
	"github.com/larryr/gems/msgrelay/health"
	"github.com/larryr/gems/msgrelay/mm"
)

type fctTyp struct {
	name string
	fct  func(w http.ResponseWriter, _ *http.Request)
}

var fctnames = ""

func SetupHandlers(prefix string) {

	fcts := []fctTyp{
		{"/", handler},
		{"/help/", handler},
		{"/mm/", mm.Handler},
		{"/health/", health.Handler},
	}

	for _, fct := range fcts {
		http.HandleFunc(prefix+fct.name, fct.fct)
		fctnames = fmt.Sprintf("%s%s\n", fctnames, fct.name)
	}
}

/* /help
 * show allowable routes of this service's http endpoint
 */
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Help: %s\n%s", r.URL.Path[1:], fctnames)
}

// dump env to log
func LogDump(l *log.Logger) {
	strDump(l, ienv.EnvInfo())
}

func strDump(l *log.Logger, vals []string) {
	for _, v := range vals {
		l.Printf("%s", v)
	}
}
