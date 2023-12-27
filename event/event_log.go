package event

type Event struct {
	kind string
}

func NewEvent(kind string) Event {
	return Event{
		kind: kind,
	}
}

func (my Event) Kind() string {
	return my.kind
}

type Loggable interface {
	Append(Event) error
	Latest() (Event, error)
}

type InMemoryLog struct {
	events []Event
}

func (my *InMemoryLog) Append(anEvent Event) error {
	my.events = append(my.events, anEvent)
	return nil
}

func (my *InMemoryLog) Latest() (Event, error) {
	return my.events[len(my.events)-1], nil
}

var _ Loggable = &InMemoryLog{}
