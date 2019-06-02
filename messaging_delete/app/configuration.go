package app

type Configuration struct {
	LineBot struct {
		Secret string `json:"secret"`
		Token  string `json:"token"`
	} `json:"line_bot"`
	AuthToken string `json:"auth_token"`
	API       struct {
		GetExpenses      string `json:"get_expenses"`
		GetLastNExpenses string `json:"get_last_n_expenses"`
		GetIncomes       string `json:"get_incomes"`
		GetLatestReset   string `json:"get_latest_reset"`
		DeleteExpense    string `json:"delete_expense"`
	} `json:"api"`
}
