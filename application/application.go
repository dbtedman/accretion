package application

import "github.com/dbtedman/accretion/event"

type Application struct {
	log event.Loggable
}

func NewApplication(log event.Loggable) Application {
	return Application{
		log: log,
	}
}

func (my *Application) Start() error {
	_ = my.log.Append(event.NewEvent("application_startup"))
	return nil
}
