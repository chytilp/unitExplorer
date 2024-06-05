package request

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestUnmarshallDomainsJson(t *testing.T) {
	jsonString := `
	{
		"payload": [
			{
				"id": "1",
				"name": "football"
			},
			{
				"id": "2",
				"name": "tennis"
			}
		],
		"session": {},
		"taskId": "task-list-domains-80cbdde4-4289-4b56-b653-01793476d683"
	}`
	data := []byte(jsonString)
	payload := DomainPayload{}
	err := json.Unmarshal(data, &payload)
	if err != nil {
		t.Errorf("err: %#v", err)
		return
	}
	assert.Equal(t, payload.TaskId, "task-list-domains-80cbdde4-4289-4b56-b653-01793476d683")
	assert.Equal(t, payload.Session, map[string]interface{}{})
	assert.Equal(t, len(payload.Payload), 2)
	assert.Equal(t, payload.Payload[0], Domain{Base{Id: "1", Name: "football"}})
}

func TestUnmarshallEventsJson(t *testing.T) {
	jsonString := `
	{
		"payload": [
			{
				"competition": {
					"id": "indonesia_liga_1",
					"name": "Indonesia Liga 1",
					"url": null
				},
				"competitors": [
					{
						"id": "football|PSIS Semarang",
						"name": "PSIS Semarang"
					},
					{
						"id": "football|RANS Nusantara FC",
						"name": "RANS Nusantara FC"
					}
				],
				"fullData": true,
				"id": "FBL-934156",
				"name": "PSIS Semarang vs RANS Nusantara FC",
				"startDate": "2024-04-22T15:08:44.145790",
				"type": "MATCH"
			},
			{
				"competition": {
					"id": "premier_league_2_(division_1)",
					"name": "Premier League 2 (Division 1)",
					"url": null
				},
				"competitors": [
					{
						"id": "football|Leeds United (U21)",
						"name": "Leeds United (U21)"
					},
					{
						"id": "football|Everton FC (U21)",
						"name": "Everton FC (U21)"
					}
				],
				"fullData": true,
				"id": "FBL-1205306",
				"name": "Leeds United (U21) vs Everton FC (U21)",
				"startDate": "2024-04-22T15:08:44.146089",
				"type": "MATCH"
			}
		],
		"session": {"id": "1", "name": "football"},
		"taskId": "task-list-events-06f13f6c-aedb-4229-84ca-b3bf03e51463"
	}`
	data := []byte(jsonString)
	payload := EventPayload{}
	err := json.Unmarshal(data, &payload)
	if err != nil {
		t.Errorf("err: %#v", err)
		return
	}
	assert.Equal(t, payload.TaskId, "task-list-events-06f13f6c-aedb-4229-84ca-b3bf03e51463")
	assert.Equal(t, payload.Session, map[string]interface{}{"Id": "1", "Name": "football"})
	assert.Equal(t, len(payload.Payload), 2)
	assert.Equal(t, payload.Payload[0], Event{
		Base:        Base{Id: "FBL-934156", Name: "PSIS Semarang vs RANS Nusantara FC"},
		StartDate:   time.Date(2024, 4, 22, 15, 8, 44, 145790, time.Local("")),
		FullData:    true,
		Type:        "MATCH",
		Competition: Competition{Base{Id: "indonesia_liga_1", Name: "Indonesia Liga 1"}, nil},
		Competitors: []Competitor{
			Competitor{Base{Id: "football|PSIS Semarang", Name: "PSIS Semarang"}},
			Competitor{Base{Id: "football|RANS Nusantara FC", Name: "RANS Nusantara FC"}},
		},
	})
}

/*func TestUnmarshallMarketsJson(t *testing.T) {
	jsonString := `
	{
		"payload": [
			{
				"fullData": true,
				"id": "72735c20-bdb7-4abb-8ec4-25e0403559a8",
				"name": "Match Odds - Half Time",
				"selections": [
					{
						"id": "0",
						"name": "Leeds United (U21)",
						"odds": 16.0
					},
					{
						"id": "1",
						"name": "Everton FC (U21)",
						"odds": 1.13
					},
					{
						"id": "2",
						"name": "Draw",
						"odds": 4.9
					}
				],
				"startDate": null,
				"winningSelections": 1
			},
			{
				"fullData": true,
				"id": "698dd98f-aa98-4fb4-9080-99cb0a7dc65d",
				"name": "Match Odds",
				"selections": [
					{
						"id": "0",
						"name": "Leeds United (U21)",
						"odds": 3.78
					},
					{
						"id": "1",
						"name": "Everton FC (U21)",
						"odds": 1.64
					},
					{
						"id": "2",
						"name": "Draw",
						"odds": 4.04
					}
				],
				"startDate": null,
				"winningSelections": 1
			}
		],
		"session": {
			"competition": {
				"id": "premier_league_2_(division_1)",
				"name": "Premier League 2 (Division 1)",
				"url": null
			},
			"competitors": [
				{
					"id": "football|Leeds United (U21)",
					"name": "Leeds United (U21)"
				},
				{
					"id": "football|Everton FC (U21)",
					"name": "Everton FC (U21)"
				}
			],
			"fullData": true,
			"id": "FBL-1205306",
			"name": "Leeds United (U21) vs Everton FC (U21)",
			"startDate": "2024-04-22T15:08:44.146089",
			"type": "MATCH"
		},
		"taskId": "task-list-markets-f60bc82f-f8cf-4cbe-a566-30c76878c95b"
	}`
	data := []byte(jsonString)
	payload := MarketPayload{}
	err := json.Unmarshal(data, &payload)
	if err != nil {
		t.Errorf("err: %#v", err)
		return
	}
	assert.Equal(t, payload.TaskId, "task-list-domains-80cbdde4-4289-4b56-b653-01793476d683")
	assert.Equal(t, payload.Session, map[string]interface{}{})
	assert.Equal(t, len(payload.Payload), 2)
	assert.Equal(t, payload.Payload[0], Market{Base{Id: "1", Name: "football"}})
}*/
