package app

type Configuration struct {
	DB struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Location string `json:"location"`
		Database string `json:"database"`
	} `json:"db"`
	LineBot struct {
		Secret string `json:"secret"`
		Token  string `json:"token"`
	} `json:"line_bot"`
	AuthToken string `json:"auth_token"`
}
