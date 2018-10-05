package api

import (
	"net/http"
)

var (
	IdMw Middleware = createParamMv("id")
)

var routes = Routes{

	Route{"GetInformations", http.MethodGet, "/infos/", getInfo, chain(LogMw)},

	Route{"GetEnvironment", http.MethodGet, "/environment/", getEnvironment, chain(LogMw)},
	Route{"PostEnvironment", http.MethodPost, "/environment/", loadEnvironment, chain(LogMw, BodyMw)},
	Route{"PutEnvironment", http.MethodPut, "/environment/", updateEnvironment, chain(LogMw, BodyMw)},
	Route{"DeleteEnvironment", http.MethodDelete, "/environment/", deleteEnvironment, chain(LogMw)},

	Route{"EnvironmentCheck", http.MethodPost, "/check/", checkEnvironment, chain(LogMw, BodyMw)},

	Route{"GetTasks", http.MethodGet, "/tasks/", getTasks, chain(LogMw)},
	Route{"GetTask", http.MethodGet, "/tasks/{id}", getTaskDetails, chain(LogMw, IdMw)},
	Route{"PutTask", http.MethodPut, "/tasks/{id}", runTask, chain(LogMw, IdMw)},

	Route{"GetNodes", http.MethodGet, "/nodes/", getNodes, chain(LogMw)},
	Route{"GetNode", http.MethodGet, "/nodes/{id}", getNodeDetails, chain(LogMw, IdMw)},

	Route{"PostStorage", http.MethodPost, "/storage/", saveValue, chain(LogMw, BodyMw)},
	Route{"PutStorage", http.MethodPut, "/storage/", saveValue, chain(LogMw, BodyMw)},
	Route{"GetStorage", http.MethodGet, "/storage/{id}", getValue, chain(LogMw, IdMw)},
	Route{"DeleteStorage", http.MethodDelete, "/storage/{id}", deleteValue, chain(LogMw, IdMw)},

	Route{"GetKeys", http.MethodGet, "/storage/", getKeys, chain(LogMw)},

	Route{"PostEnvironmentParam", http.MethodPost, "/envparam/", saveEnvironmentParam, chain(LogMw, BodyMw)},
	Route{"PutEnvironmentParam", http.MethodPut, "/envparam/", saveEnvironmentParam, chain(LogMw, BodyMw)},
	Route{"GetEnvironmentParam", http.MethodGet, "/envparam/", getEnvironmentParam, chain(LogMw, IdMw)},
	Route{"DeleteEnvironmentParam", http.MethodDelete, "/envparam/", deleteEnvironmentParam, chain(LogMw, IdMw)},

	Route{"PostEnvironmentSession", http.MethodPost, "/envsession/", saveEnvironmentSession, chain(LogMw, BodyMw)},
	Route{"PutEnvironmentSession", http.MethodPut, "/envsession/", saveEnvironmentSession, chain(LogMw, BodyMw)},
	Route{"GetEnvironmentSession", http.MethodGet, "/envsession/", getEnvironmentSession, chain(LogMw, IdMw)},
	Route{"DeleteEnvironmentSession", http.MethodDelete, "/envsession/", deleteEnvironmentSession, chain(LogMw, IdMw)},
}
