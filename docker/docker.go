package docker

import (
	_ "log"
	_ "net/http"
	_ "path/filepath"

	_ "docker.io/go-docker"
	_ "docker.io/go-docker/api/types"
	_ "github.com/docker/docker/client"
	_ "golang.org/x/net/context"

	_ "github.com/docker/go-connections/tlsconfig"
)

//var cli docker.Client

//func CreateLocalClient(url string) (*client.Client, error) {
//	return client.NewClientWithOpts(
//		client.WithHost(endpoint.URL),
//		//client.WithVersion(portainer.SupportedDockerAPIVersion),
//	)
//}
