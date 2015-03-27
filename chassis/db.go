//
// Chassis.
//

package chassis

import (
	"database/sql"
	log "github.com/Sirupsen/logrus"
	pq "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

type InTransaction func(tx *sql.Tx) error

// Opens the specified database.
func OpenDatabase(url string) (*sql.DB, error) {
	log.Infof("Opening database %s", url)
	db, err := sql.Open(driverName, url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Runs the specified function in a database transaction.
func WithTransaction(db *sql.DB, f InTransaction) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("%v", err)
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		if err != nil {
			log.Warnf("%v\n", err.Error())
		}
	}()
	if err := f(tx); err != nil {
		panic(err)
	}
	return nil
}

func IsUniqueViolation(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code == "23505"
	}
	return false
}
