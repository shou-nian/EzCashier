package utils

import (
	"fmt"
	"os"
)

type ConnectionName string

const (
	MysqlConnection  ConnectionName = "mysql"
	ServerConnection ConnectionName = "server"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n ConnectionName) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case MysqlConnection:
		// URL for MySQL connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
	case ServerConnection:
		// URL for server connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
