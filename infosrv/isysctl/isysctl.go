package isysctl

import (
	"fmt"
	"net/http"

	"github.com/larryr/gems/infosrv/isysctl/sysctl"
)

func Handler() http.HandlerFunc {
	return handler
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "<html>\n")
	vals := SysctlInfo()
	for _, v := range vals {
		fmt.Fprintf(w, "<li>%s</li>\n", v)
	}
	fmt.Fprintf(w, "</html>\n")
}

func SysctlInfo() []string {
	s := []string{"sysctl"}
	sysctls, err := sysctl.GetAll()
	if err != nil {
		return append(s, fmt.Sprintf("%v", err))
	}

	for n, v := range sysctls {
		s = append(s, fmt.Sprintf("%s => %s", n, v))
	}
	return s
}
