// +build ignore

package lib

import (
	"context"

	"specials-gitea.kpn.org/gcp/go-oceanstor-sdk/dorado"
)

// input param
const (
	StoragePoolName      = ""
	HyperMetroDomainName = ""
)

// GetClient get dorado.Client
func GetClient() (*dorado.Client, error) {
	username := "admin"
	password := ""
	localIps := []string{"", ""}
	remoteIps := []string{"", ""}

	return dorado.NewClient(context.Background(), localIps, remoteIps, username, password, nil)
}
