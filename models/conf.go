package models

// Conf is a model for configuration
type Conf struct {
	Version         string
	Port            string
	Cors struct {
		AuthorizedHosts []string
		AuthorizedMethods []string
		AuthorizedHeaders []string
	}
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
