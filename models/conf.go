package models

// Conf is a model for configuration
type Conf struct {
	Port            string
	AuthorizedHosts []string
	Database        struct {
		MysqlHost     string
		MysqlPort     string
		MysqlUser     string
		MysqlPassword string
		MysqlDatabase string
	}
	Log struct {
		Formatter string
		Output    string
		Level     string
	}
}
