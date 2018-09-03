package api

import (
	"net/http"
)

var (
	IdMw Middleware = createParamMv("id")
)

var routes = Routes{

	Route{"Informations", http.MethodGet, "/infos/", getInfo, chain(LogMw)},

	Route{"Environment", http.MethodGet, "/environment/", getEnvironment, chain(LogMw)},
	Route{"Environment", http.MethodPost, "/environment/", loadEnvironment, chain(LogMw, BodyMw)},
	Route{"Environment", http.MethodPut, "/environment/", updateEnvironment, chain(LogMw, BodyMw)},
	Route{"Environment", http.MethodDelete, "/environment/", deleteEnvironment, chain(LogMw)},

	Route{"EnvironmentCheck", http.MethodPost, "/check/", checkEnvironment, chain(LogMw, BodyMw)},

	Route{"Task", http.MethodGet, "/tasks/", getTasks, chain(LogMw)},
	Route{"Task", http.MethodGet, "/tasks/{id}", getTaskDetails, chain(LogMw, IdMw)},
	Route{"Task", http.MethodPut, "/tasks/{id}", runTask, chain(LogMw, IdMw)},

	Route{"Node", http.MethodGet, "/nodes/", getNodes, chain(LogMw)},
	Route{"Node", http.MethodGet, "/nodes/{id}", getNodeDetails, chain(LogMw, IdMw)},

	Route{"Storage", http.MethodPost, "/storage/", saveValue, chain(LogMw, BodyMw)},
	Route{"Storage", http.MethodPut, "/storage/", saveValue, chain(LogMw, BodyMw)},
	Route{"Storage", http.MethodGet, "/storage/{id}", getValue, chain(LogMw, IdMw)},
	Route{"Storage", http.MethodDelete, "/storage/{id}", deleteValue, chain(LogMw, IdMw)},

	Route{"EnvironmentParam", http.MethodPost, "/envparam/", saveEnvironmentParam, chain(LogMw, BodyMw)},
	Route{"EnvironmentParam", http.MethodPut, "/envparam/", saveEnvironmentParam, chain(LogMw, BodyMw)},
	Route{"EnvironmentParam", http.MethodGet, "/envparam/", getEnvironmentParam, chain(LogMw, IdMw)},
	Route{"EnvironmentParam", http.MethodDelete, "/envparam/", deleteEnvironmentParam, chain(LogMw, IdMw)},
}
