package ienv

import (
	"fmt"
	"net/http"
	"os"
)

func Handler() http.HandlerFunc {
	return handler
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "<html>\n")
	vals := os.Environ()
	for _, v := range vals {
		fmt.Fprintf(w, "<li>%s</li>\n", v)
	}
	fmt.Fprintf(w, "</html>\n")
}

func EnvInfo() []string {
	return os.Environ()
}
