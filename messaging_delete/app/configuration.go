package app

type Configuration struct {
	LineBot struct {
		Secret string `json:"secret"`
		Token  string `json:"token"`
	} `json:"line_bot"`
	AuthToken string `json:"auth_token"`
	API       struct {
		GetIncomes       string `json:"get_incomes"`
		GetExpenses      string `json:"get_expenses"`
		GetLastNExpenses string `json:"get_last_n_expenses"`
		GetLastNIncomes  string `json:"get_last_n_incomes"`
		GetLatestReset   string `json:"get_latest_reset"`
		DeleteIncome     string `json:"delete_income"`
		DeleteExpense    string `json:"delete_expense"`
	} `json:"api"`
}
