package ios

import (
	"fmt"
	"syscall"
)

// linuxSysInfo defines the internal type used to represent linux system information.
type linuxSysInfo struct {
	SystemName string
	NodeName   string
	Release    string
	Version    string
	Machine    string
	DomainName string
}

// GetSysInfo for linux platform
func getSysInfo() (*SysInfo, error) {

	// Get linux system information.
	info, err := getLinuxSysInfo()
	if err != nil {
		return nil, fmt.Errorf("getLinuxSysInfo: %s", err.Error())
	}

	// Get uptime.
	uptime, err := getUptime()
	if err != nil {
		return nil, fmt.Errorf("getUptime: %s", err.Error())
	}

	return &SysInfo{
		SystemName: info.SystemName,
		NodeName:   info.NodeName,
		Release:    info.Release,
		Version:    info.Version,
		Machine:    info.Machine,
		DomainName: info.DomainName,
		Uptime:     uptime,
	}, nil
}

// getLinuxSysInfo returns information of a linux-based systems.
func getLinuxSysInfo() (*linuxSysInfo, error) {
	var uname syscall.Utsname

	err := syscall.Uname(&uname)
	if err != nil {
		return nil, fmt.Errorf("Uname: %s", err.Error())
	}

	info := &linuxSysInfo{
		SystemName: toString(uname.Sysname),
		NodeName:   toString(uname.Nodename),
		Release:    toString(uname.Release),
		Version:    toString(uname.Version),
		Machine:    toString(uname.Machine),
		DomainName: toString(uname.Domainname),
	}

	return info, nil
}

// getUptime returns the machine uptime
func getUptime() (int64, error) {
	var sysinfo syscall.Sysinfo_t

	err := syscall.Sysinfo(&sysinfo)
	if err != nil {
		return 0, fmt.Errorf("Sysinfo: %s", err.Error())
	}

	// Note that Uptime will be int32 or int64 based
	// on the architecture.
	return int64(sysinfo.Uptime), nil
}

func toString(a [65]int8) string {
	var s []byte
	for _, c := range a {
		if c == 0 {
			break
		}
		s = append(s, byte(c))
	}

	return string(s)
}
