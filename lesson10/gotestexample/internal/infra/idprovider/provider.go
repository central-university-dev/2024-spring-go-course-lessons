package idprovider

import gonanoid "github.com/matoous/go-nanoid/v2"

type NanoIDProvider struct{}

func (n NanoIDProvider) Provide() string {
	id, err := gonanoid.New(7)
	if err != nil {
		panic(err)
	}

	return id
}
