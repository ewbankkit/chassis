//
// Chassis.
//

package chassis

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"os"
	"path/filepath"
)

// Read the service configuration file.
func ReadConfigFile(v interface{}) error {
	cwd, err := os.Getwd()
	if err != nil {
		return nil
	}
	name := "config.json"
	env := os.Getenv("EUCALYPTUS_ENV")
	if env != "" {
		name = "config." + env + ".json"
	}
	name = filepath.Join(cwd, name)

	log.Infof("Reading configuration file %s", name)
	file, err := os.Open(name)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
}
