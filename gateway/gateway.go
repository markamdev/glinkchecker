//go:build !linux

package gateway

import "fmt"

func DetectDefaultGateway() (string, error) {
	return "", fmt.Errorf("Not supported on this OS")
}
