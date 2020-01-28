// +build sqlserver

package gormigrate

import (
	_ "github.com/chronosphereio/gorm/dialects/mssql"
)

func init() {
	databases = append(databases, database{
		name:    "mssql",
		connEnv: "SQLSERVER_CONN_STRING",
	})
}
