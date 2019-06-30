package messaging_income

import (
	"messaging_income/app/http_handler"
	"messaging_income/app/message_handler"
)

var IncomeMessage = message_handler.IncomeMessage
var GetIncome = http_handler.GetIncome
var GetLastNIncomes = http_handler.GetLastNIncomes
var DeleteIncome = http_handler.DeleteIncome
