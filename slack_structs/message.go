package slack_structs

type MessageWrapper struct {
	Ok bool
	Token string
	Channel string
	Ts string
	Message Message
}

type Message struct {
	Text string
	Username string
	BotId string `json:"bot_id"`
	Attachments []Attachment
	Type string
	Subtype string
	Ts string
}

type Attachment struct {
	Text string
	Id int
	Fallback string
}