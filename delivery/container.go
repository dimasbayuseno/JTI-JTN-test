package delivery

import (
	"initial/configuration"
	"initial/domain/sample/samplerepository"
	"initial/domain/sample/sampleservice"
)

type Container struct {
	SampleService sampleservice.Service
}

func SetupContainer() Container {
	configuration.Env()
	database, err := configuration.NewDatabase()
	if err != nil {
		panic(err)
	}

	// Initialize samplerepository
	sampleDataRepository := samplerepository.New(database)

	// Initialize sampleservice
	sampleService := sampleservice.New(sampleDataRepository)

	return Container{
		SampleService: sampleService,
	}

}
