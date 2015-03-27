//
// Chassis.
//

package chassis

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	maxBodyLen        = 1024 * 1024
	rfc3339MilliFixed = "2006-01-02T15:04:05.000Z07:00"
)

func TimeToJson(t time.Time) string {
	return t.Format(rfc3339MilliFixed)
}

func ReadJson(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxBodyLen))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, v); err != nil {
		return err
	}
	return nil
}

func WriteJson(w http.ResponseWriter, code int, v interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
