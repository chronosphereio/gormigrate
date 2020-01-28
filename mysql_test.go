// +build mysql

package gormigrate

import (
	_ "github.com/chronosphereio/gorm/dialects/mysql"
)

func init() {
	databases = append(databases, database{
		name:    "mysql",
		connEnv: "MYSQL_CONN_STRING",
	})
}
