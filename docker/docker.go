package docker

import (
	"log"
	"net/http"
	"path/filepath"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"golang.org/x/net/context"

	"github.com/docker/go-connections/tlsconfig"
)

var cli docker.Client

func TestDocker(l *log.Logger, host string, api string, path string) {

	options := tlsconfig.Options{
		CAFile:             filepath.Join(path, "ca.pem"),
		CertFile:           filepath.Join(path, "cert.pem"),
		KeyFile:            filepath.Join(path, "key.pem"),
		InsecureSkipVerify: false,
	}

	tlsc, err := tlsconfig.Client(options)
	if err != nil {
		panic(err)
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsc,
		},
		CheckRedirect: docker.CheckRedirect,
	}

	c, err := docker.NewClient(host, api, httpClient, nil)
	if err != nil {
		panic(err)
	}
	cli = *c
	l.Println("Getting container...")
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	l.Printf("NB container :%v\n", len(containers))
	for _, container := range containers {
		l.Printf("Stopping container ", container.ID[:10], "... ")

		l.Println("Success")
	}
}
