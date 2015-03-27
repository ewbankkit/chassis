//
// Chassis.
//

package chassis

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

// Run a REST server on the specified routes.
func RunRestServer(addr string, routes []Route) error {
	log.Infof("Starting REST server on %s", addr)

	router := newRouter(routes)
	if err := http.ListenAndServe(addr, router); err != nil {
		return err
	}

	return nil
}

func GetPathParamValue(r *http.Request, p string) string {
	return mux.Vars(r)[p]
}

func newRouter(routes []Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := logger(route.Handler, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof(
			"%v\t%s",
			r.Method,
			r.RequestURI)
		inner.ServeHTTP(w, r)
	})
}
