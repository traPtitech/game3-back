package config

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/traPtitech/game3-back/internal/pkg/util"
)

func AppAddr() string {
	return util.GetEnvOrDefault("APP_ADDR", ":8080")
}

func MySQL() *mysql.Config {
	c := mysql.NewConfig()

	neoshowcase := util.GetEnvOrDefault("NEOSHOWCASE", "0")
	if neoshowcase == "1" {
		c.User = util.GetEnvOrDefault("NS_MARIADB_USER", "root")
		c.Passwd = util.GetEnvOrDefault("NS_MARIADB_PASSWORD", "pass")
		c.Net = util.GetEnvOrDefault("DB_NET", "tcp")
		c.Addr = fmt.Sprintf(
			"%s:%s",
			util.GetEnvOrDefault("NS_MARIADB_HOSTNAME", "localhost"),
			util.GetEnvOrDefault("NS_MARIADB_PORT", "3306"),
		)
		c.DBName = util.GetEnvOrDefault("NS_MARIADB_DATABASE", "app")
	} else {
		c.User = util.GetEnvOrDefault("DB_USER", "root")
		c.Passwd = util.GetEnvOrDefault("DB_PASSWORD", "pass")
		c.Net = util.GetEnvOrDefault("DB_NET", "tcp")
		c.Addr = fmt.Sprintf(
			"%s:%s",
			util.GetEnvOrDefault("DB_HOST", "localhost"),
			util.GetEnvOrDefault("DB_PORT", "3306"),
		)
		c.DBName = util.GetEnvOrDefault("DB_NAME", "app")
	}
	c.Collation = "utf8mb4_general_ci"
	c.AllowNativePasswords = true
	c.ParseTime = true

	return c
}
