package utilities

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/entgigi/custom-ingress-rest/commons"
)

func GetSrvPort() string {
	var port string
	found := true
	if port, found = os.LookupEnv(commons.SrvPrtEnvVar); !found {
		port = ":8080"
	}
	return port
}

// GetWatchNamespace returns the Namespace the operator should be watching for changes
func GetWatchNamespace() (string, error) {
	// WatchNamespaceEnvVar is the constant for env variable WATCH_NAMESPACE
	// which specifies the Namespace to watch.
	// An empty value means the operator is running with cluster scope.

	ns, found := os.LookupEnv(commons.WatchNamespaceEnvVar)
	if !found {
		return "", fmt.Errorf("%s must be set", commons.WatchNamespaceEnvVar)
	}
	return ns, nil
}

func TruncateString(s string, max int) string {
	if max > len(s) {
		return s
	}
	return s[:max]
}

func GenerateSha256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}
