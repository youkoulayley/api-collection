package models

// Conf is a model for configuration
type Conf struct {
	Version         string
	Port            string
	AuthorizedHosts []string
	Database        struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
	Log struct {
		Formatter string
		Output    string
		Level     string
	}
}
