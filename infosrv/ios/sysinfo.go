package ios

import "fmt"

// SysInfo contains system/platform specific information.
type SysInfo struct {
	SystemName string
	NodeName   string // System node name (hostname)
	Release    string // Kernel release
	Version    string
	Machine    string // Machine hardware name
	DomainName string
	Uptime     int64
}

func systeminfo() []string {
	si, err := getSysInfo()
	if err != nil {
		return nil
	}
	var s []string
	s = append(s, fmt.Sprintf("systemname=%s", si.SystemName))
	s = append(s, fmt.Sprintf("nodename=%s", si.NodeName))
	s = append(s, fmt.Sprintf("release=%s", si.Release))
	s = append(s, fmt.Sprintf("version=%s", si.Version))
	s = append(s, fmt.Sprintf("machine=%s", si.Machine))
	s = append(s, fmt.Sprintf("domainname=%s", si.DomainName))
	s = append(s, fmt.Sprintf("uptime=%d", si.Uptime))
	return s
}
