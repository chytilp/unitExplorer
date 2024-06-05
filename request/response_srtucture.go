package request

type ResponseBase struct {
	TaskId  string                 `json:"taskId"`
	Session map[string]interface{} `json:"session"`
}

type DomainPayload struct {
	ResponseBase
	Payload []Domain `json:"payload"`
}

type EventPayload struct {
	ResponseBase
	Payload []Event `json:"payload"`
}

type MarketPayload struct {
	ResponseBase
	Payload []Market `json:"payload"`
}
