package inet

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func Handler() http.HandlerFunc {
	return handler
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "<html>\n")
	vals := NetInfo()
	for _, v := range vals {
		fmt.Fprintf(w, "<li>%s</li>\n", v)
	}
	fmt.Fprintf(w, "</html>\n")
}

func NetInfo() []string {
	s := interfaces()
	h := hosts()
	s = append(s, h...)
	return s
}

func interfaces() []string {
	var (
		s     []string
		addrs []net.Addr
	)
	intfs, _ := net.Interfaces()
	for _, i := range intfs {
		a, _ := i.Addrs()
		addrs = append(addrs, a...)
		s = append(s, fmt.Sprintf("intf-%d=%v,\t%s,\t%d,\t%v", i.Index, a, i.Name, i.MTU, i.Flags))
	}
	names := hostnames(addrs)

	s = append(s, names...)
	return s
}

func hostnames(addrs []net.Addr) []string {
	var names []string
	for _, c := range addrs {
		a, _, _ := net.ParseCIDR(c.String())
		hosts, err := net.LookupAddr(a.String())
		if err != nil {
			log.Printf("error:%v", err)
			continue
		}
		for _, n := range hosts {
			nstr := fmt.Sprintf("\t    %s -> %s\n", a, n)
			names = append(names, nstr)
		}
	}
	return names
}

var hostsList = []string{"infosrv", "genfract", "gopuff.com", "google.com"}

// find some hosts based on DNS address.
// use any list of DNS names in global hostsList
func hosts() []string {
	r := []string{"DNS:"}

	for _, n := range hostsList {
		r = append(r, fmt.Sprintf("Host:%s", n))

		//look up addr
		s, err := net.LookupHost(n)
		if err == nil {
			r = append(r, "host address:")
			for i, h := range s {
				r = append(r, fmt.Sprintf("\t%d-> %s", i, h))
			}
		} else {
			log.Printf("Lookup Host err: %v", err)
		}

		//look up cname
		cn, err := net.LookupCNAME(n)
		if err == nil {
			r = append(r, fmt.Sprintf("cname=[%s]", cn))
		} else {
			log.Printf("cname err: %v", err)
		}

		//look up txt
		s, err = net.LookupTXT(n)
		if err == nil {
			r = append(r, "TXT entries:")
			for i, h := range s {
				r = append(r, fmt.Sprintf("\t%d-> %s", i, h))
			}
		} else {
			log.Printf("TXT err: %v", err)
		}

		cn, srv, err := net.LookupSRV("", "", n)
		if err == nil {
			r = append(r, fmt.Sprintf("SRV records:%s", cn))
			for i, h := range srv {
				r = append(r, fmt.Sprintf("\t%d-> t=%s p=%d pri=%d w=%d", i, h.Target, h.Port, h.Priority, h.Weight))
			}
		} else {
			log.Printf("SRV err: %v", err)
		}

		ns, err := net.LookupNS(n)
		if err == nil {
			r = append(r, "NS records:")
			for i, h := range ns {
				r = append(r, fmt.Sprintf("\t%d-> %s", i, h.Host))
			}
		} else {
			log.Printf("NS err: %v", err)
		}
	}
	return r
}
