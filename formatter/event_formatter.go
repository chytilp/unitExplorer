package formatter

import (
	"fmt"

	"github.com/chytilp/unitExplorer/request"
)

type EventFormatter struct {
	filters *[]request.Base
	results *[]request.Event
}

func NewEventFormatter() *EventFormatter {
	return &EventFormatter{}
}

func (e *EventFormatter) SetFilters(filters []request.Base) {
	e.filters = &filters
}

func (e *EventFormatter) SetResults(results []request.Event) {
	e.results = &results
}

func (e *EventFormatter) Print() error {
	if e.filters == nil {
		return fmt.Errorf("no filters assign")
	}
	if e.results == nil {
		return fmt.Errorf("no results assign")
	}
	for _, event := range *e.results {
		fmt.Printf("event -> id: %s, name: %s\n", event.Id, event.Name)
	}
	return nil
}
