package database

var connection string

func init() {
	connection = "Postgres"
}

func GetDatabase() string {
	return connection
}