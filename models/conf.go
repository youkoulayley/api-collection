package models

// Conf is a model for configuration
type Conf struct {
	Port     string
	Database struct {
		MysqlHost     string
		MysqlPort     string
		MysqlUser     string
		MysqlPassword string
		MysqlDatabase string
	}
}
