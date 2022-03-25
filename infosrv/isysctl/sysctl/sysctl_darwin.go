package sysctl

import "fmt"

// GetAll returns all sysctls. This is equivalent
// to running the command sysctl -a.
func GetAll() (map[string]string, error) {
	return nil, fmt.Errorf("sorry, you loose, you're on a mac")
}
