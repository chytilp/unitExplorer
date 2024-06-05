package request

type Source struct {
	Id    int     `json:"id"`
	Proxy *string `json:"proxy"`
}

type Request struct {
	Task         string                 `json:"task"`
	Source       Source                 `json:"source"`
	Parameters   map[string]interface{} `json:"parameters"`
	Session      map[string]interface{} `json:"session"`
	ExpectBodyIn int                    `json:"expectBodyIn"`
}

func CreateDomainRequest(sourceId int) (Request, error) {
	return Request{
		Task: "LIST_DOMAINS",
		Source: Source{
			Id:    sourceId,
			Proxy: nil,
		},
		Parameters:   map[string]interface{}{},
		Session:      map[string]interface{}{},
		ExpectBodyIn: 150,
	}, nil
}

func CreateEventRequest(sourceId int, domain Domain) (Request, error) {
	parameters := map[string]interface{}{"id": domain.Id, "name": domain.Name}
	return Request{
		Task: "LIST_EVENTS",
		Source: Source{
			Id:    sourceId,
			Proxy: nil,
		},
		Parameters:   parameters,
		Session:      map[string]interface{}{},
		ExpectBodyIn: 300,
	}, nil
}

func createCompetitor(competitor Competitor) map[string]interface{} {
	return map[string]interface{}{
		"id":   competitor.Id,
		"name": competitor.Name,
	}
}

func createEventParameter(event Event) map[string]interface{} {
	competitors := []map[string]interface{}{
		createCompetitor(event.Competitors[0]),
		createCompetitor(event.Competitors[1]),
	}
	competition := map[string]interface{}{
		"id":   event.Competition.Id,
		"name": event.Competition.Name,
		"url":  event.Competition.Url,
	}

	return map[string]interface{}{
		"competition": competition,
		"competitors": competitors,
		"fullData":    event.FullData,
		"startDate":   event.StartDate,
		"type":        event.Type,
		"id":          event.Id,
		"name":        event.Name,
	}
}

func CreateMarketRequest(sourceId int, domain Domain, event Event) (Request, error) {
	session := map[string]interface{}{"id": domain.Id, "name": domain.Name}
	parameters := createEventParameter(event)
	return Request{
		Task: "LIST_MARKETS",
		Source: Source{
			Id:    sourceId,
			Proxy: nil,
		},
		Parameters:   parameters,
		Session:      session,
		ExpectBodyIn: 300,
	}, nil
}
