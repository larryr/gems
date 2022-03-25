package ifiles

import (
	"fmt"
	"net/http"

	"github.com/larryr/gems/infosrv/ifiles/mountinfo"
)

func Handler() http.HandlerFunc {
	return handler
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "<html>\n")
	vals := FSInfo()
	for _, v := range vals {
		fmt.Fprintf(w, "<li>%s</li>\n", v)
	}
	fmt.Fprintf(w, "</html>\n")
}

func FSInfo() []string {

	s := files()
	s = append(s, mounts()...)

	return s
}

func files() []string {
	s := []string{"key files"}

	s = append(s, mydata()...)
	s = append(s, k8sfiles()...)
	return s
}
func mydata() []string {
	s := []string{"data volume"}
	return s
}

func mounts() []string {
	s := []string{"mounts"}
	mnts, err := mountinfo.GetMounts(nil)
	if err != nil {
		return s
	}

	for _, m := range mnts {
		s = append(s, toString(m))
	}
	return s
}

func toString(i *mountinfo.Info) string {
	return fmt.Sprintf("%x:%x %x:%x fsTyp=%s [%s] [%s] [%s] src=%s root=%s mnt=%s",
		i.ID, i.Parent, i.Major, i.Minor,
		i.FSType,
		i.Options, i.Optional, i.VFSOptions,
		i.Source,
		i.Root,
		i.Mountpoint,
	)
}
