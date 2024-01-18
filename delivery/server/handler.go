package server

import (
	"initial/delivery"
	"initial/domain/sample"
)

type handler struct {
	sampleHandler sample.HttpSampleHandler
}

func SetupHandler(container delivery.Container) handler {
	return handler{sampleHandler: sample.NewSampleHandler(container.SampleService)}
}
