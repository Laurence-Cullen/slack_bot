package slack_structs

type MessageWrapper struct {
	Ok bool `json:"ok"`
	Token string `json:"token"`
	Channel string `json:"channel"`
	Ts string `json:"ts"`
	Message Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Username string `json:"username"`
	BotId string `json:"bot_id"`
	Attachments []Attachment `json:"attachments"`
	Type string `json:"type"`
	Subtype string `json:"subtype"`
	Ts string `json:"ts"`
}

type Attachment struct {
	Text string `json:"text"`
	Id int `json:"id"`
	Fallback string `json:"fallback"`
}