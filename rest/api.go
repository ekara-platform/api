package rest

import (
	"fmt"
	"github.com/ekara-platform/api/secret"
	"github.com/ekara-platform/api/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type (
	context struct {
		logger        *log.Logger
		router        *mux.Router
		address       string
		storageEngine storage.Storage
		secretEngine  secret.Secret
		version       string
	}
)

var (
	api context
)

// Starts the API
func Init(logger *log.Logger, port int) {
	api = context{
		logger:        logger,
		router:        NewRouter(routes),
		address:       fmt.Sprintf(":%d", port),
		storageEngine: storage.GetStorage(),
		secretEngine:  secret.GetSecret(),
		version:       "1.0.0",
	}
}

func Serve() {
	api.logger.Fatal(http.ListenAndServe(api.address, api.router))
}
