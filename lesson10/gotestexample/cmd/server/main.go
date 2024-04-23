package main

import (
	"log"
	"net/http"

	"gotestexample/internal/app/command"
	"gotestexample/internal/app/query"
	"gotestexample/internal/di"
	"gotestexample/internal/infra/httpapi"
	"gotestexample/internal/infra/idprovider"
	"gotestexample/internal/infra/inmemory"
)

func main() {
	idp := idprovider.NanoIDProvider{}
	repo := inmemory.New()

	cmdCreateURL := command.NewCreateShortURLCmd(idp, repo)
	queryGetURL := query.NewGetLongURLQuery(repo)

	container := di.New(cmdCreateURL, queryGetURL)

	router := httpapi.InitRouter(container)

	//nolint:gosec // it's ok
	if err := http.ListenAndServe(":3001", router); err != nil {
		log.Fatal(err)
	}
}
