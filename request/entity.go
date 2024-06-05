package request

import (
	"fmt"
	"time"
)

type Base struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Domain struct {
	Base
}

func (d *Domain) String() string {
	return fmt.Sprintf("Id: %s, Name: %s", d.Id, d.Name)
}

type Competitor struct {
	CompetitorId int64
	Base
}

type Competition struct {
	CompetitionId int64
	Base
	Url *string `json:"url"`
}

type Event struct {
	Base
	StartDate   *time.Time   `json:"startDate"`
	FullData    bool         `json:"fullData"`
	Type        string       `json:"type"`
	Competition Competition  `json:"competition"`
	Competitors []Competitor `json:"competitors"`
}

func (e *Event) String() string {
	return fmt.Sprintf("Id: %s, Name: %s", e.Id, e.Name)
}

type Selection struct {
	Base
	Odds float64 `json:"odds"`
}

type Market struct {
	Base
	StartDate         *time.Time  `json:"startDate"`
	FullData          bool        `json:"fullData"`
	WinningSelections int         `json:"winningSelections"`
	Selections        []Selection `json:"selections"`
}
