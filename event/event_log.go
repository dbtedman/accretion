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
	Push(Event) error
	Pop() (Event, error)
}

type InMemoryLog struct {
	events []Event
}

func (my *InMemoryLog) Push(anEvent Event) error {
	my.events = append(my.events, anEvent)
	return nil
}

func (my *InMemoryLog) Pop() (Event, error) {
	popped := my.events[len(my.events)-1]
	my.events = my.events[:len(my.events)-1]
	return popped, nil
}

var _ Loggable = &InMemoryLog{}
