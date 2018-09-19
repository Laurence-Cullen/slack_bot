package slack_structs

type Event struct {
	Type string `json:"type"`
	User string `json:"user"`
	Text string `json:"text"`
	Ts string `json:"ts"`
	Channel string `json:"channel"`
	EventTs string `json:"event_ts"`
}

type EventWrapper struct {
	Token string `json:"token"`
	TeamId string `json:"team_id"`
	ApiAppId string `json:"api_app_id"`
	Event Event `json:"event"`
	Type string `json:"type"`
	EventId string `json:"event_id"`
	EventTime int64 `json:"event_time"`
	AuthedUsers []string `json:"authed_users"`
}