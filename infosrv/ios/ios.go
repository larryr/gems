package ios

import (
	"fmt"
	"net/http"
	"os"
	"os/user"
	"runtime"
)

func Handler() http.HandlerFunc {
	return handler
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "<html>\n")
	vals := OSInfo()
	for _, v := range vals {
		fmt.Fprintf(w, "<li>%s</li>\n", v)
	}
	fmt.Fprintf(w, "</html>\n")
}

func OSInfo() []string {
	var s []string

	h, _ := os.Hostname()
	s = append(s, fmt.Sprintf("host=%s", h))

	ex, _ := os.Executable()
	s = append(s, fmt.Sprintf("executable=%s", ex))

	s = append(s, fmt.Sprintf("uid=%d", os.Getuid()))
	s = append(s, fmt.Sprintf("gid=%d", os.Getgid()))
	usr, _ := user.Current()
	s = append(s, fmt.Sprintf("login name=%s", usr.Username))
	s = append(s, fmt.Sprintf("user name=%s", usr.Name))

	s = append(s, fmt.Sprintf("pid=%d", os.Getpid()))
	s = append(s, fmt.Sprintf("ppid=%d", os.Getppid()))
	s = append(s, fmt.Sprintf("pagesize=%d", os.Getpagesize()))

	s = append(s, fmt.Sprintf("gomaxprocs=%d", runtime.GOMAXPROCS(-1)))
	s = append(s, fmt.Sprintf("numcpu=%d", runtime.NumCPU()))

	//var sys syscall.Sysinfo_t
	//syscall.Sysinfo(&sys)
	//s = append(s, fmt.Sprintf("uptime=%d", sys.Uptime))
	s = append(s, systeminfo()...)

	return s
}
