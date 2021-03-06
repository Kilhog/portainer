// +build !windows

package cli

const (
	defaultBindAddress     = ":9000"
	defaultDataDirectory   = "/data"
	defaultAssetsDirectory = "."
	defaultTemplatesURL    = "https://raw.githubusercontent.com/portainer/templates/master/templates.json"
	defaultTLSVerify       = "false"
	defaultTLSCACertPath   = "/certs/ca.pem"
	defaultTLSCertPath     = "/certs/cert.pem"
	defaultTLSKeyPath      = "/certs/key.pem"
)
