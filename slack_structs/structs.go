package slack_structs

type Event struct {
	Type string
	User string
	Text string
	Ts string
	Channel string
	EventTs string `json:"event_ts"`
}

type EventPacket struct {
	Token string
	TeamId string `json:"team_id"`
	ApiAppId string `json:"api_app_id"`
	Event Event
	Type string
	EventId string `json:"event_id"`
	EventTime int64 `json:"event_time"`
	AuthedUsers []string `json:"authed_users"`
}