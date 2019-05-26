package app

type Configuration struct {
	AuthToken string `json:"auth_token"`
	API       struct {
		GetExpenses    string `json:"get_expenses"`
		GetIncomes     string `json:"get_incomes"`
		GetLatestReset string `json:"get_latest_reset"`
	} `json:"api"`
	LineBot struct {
		Secret string `json:"secret"`
		Token  string `json:"token"`
	} `json:"line_bot"`
}
