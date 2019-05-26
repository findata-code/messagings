package messaging_gateway

type Configuration struct {
	PubSub struct {
		Credential string `json:"credential"`
		ProjectId  string `json:"project_id"`
		Topic      string `json:"topic"`
	} `json:"pubsub"`
	LineBot struct {
		Secret string `json:"secret"`
		Token  string `json:"token"`
	} `json:"line_bot"`
}
