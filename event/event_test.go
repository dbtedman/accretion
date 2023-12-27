package event_test

import (
	"github.com/dbtedman/accretion/application"
	"github.com/dbtedman/accretion/event"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplicationStartupPublishesEvent(t *testing.T) {
	var log event.Loggable = &event.InMemoryLog{}
	app := application.NewApplication(log)
	err := app.Start()
	assert.Nil(t, err, "Application startup should not return error.")

	anEvent, err := log.Latest()

	assert.NotNil(t, anEvent, "A non null event should be popped.")
	assert.Equal(t, anEvent.Kind(), "application_startup")
	assert.Nil(t, err, "A nil error message should be returned.")
}
