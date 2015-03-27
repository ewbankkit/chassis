//
// Chassis.
//

package chassis

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"errors"
	log "github.com/Sirupsen/logrus"
)

func MigrateDatabase(url string) error {
	log.Infof("Migrating database %s", url)

	driver := goose.DBDriver{
		Name:    driverName,
		OpenStr: url,
		Dialect: &goose.PostgresDialect{},
		Import:  "github.com/lib/pq",
	}
	if !driver.IsValid() {
		return errors.New("invalid DBDriver")
	}

	conf := &goose.DBConf{
		MigrationsDir: "migrations",
		Env:           "all",
		Driver:        driver,
		PgSchema:      "",
	}
	target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		return err
	}
	if err := goose.RunMigrations(conf, conf.MigrationsDir, target); err != nil {
		return err
	}

	return nil
}
