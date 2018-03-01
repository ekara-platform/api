package main

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
	Route{"Environment", http.MethodDelete, "/environment/", deleteEnvironment, chain(LogMw)},

	Route{"Task", http.MethodGet, "/tasks/", getTasks, chain(LogMw)},
	Route{"Task", http.MethodGet, "/tasks/{id}", getTaskDetails, chain(LogMw, IdMw)},
	Route{"Task", http.MethodPut, "/tasks/{id}", runTask, chain(LogMw, IdMw)},

	Route{"Node", http.MethodGet, "/nodes/", getNodes, chain(LogMw)},
	Route{"Node", http.MethodGet, "/nodes/{id}", getNodeDetails, chain(LogMw, IdMw)},
}
