package messaging_reset

import (
	"messaging_reset/app/http_handler"
	"messaging_reset/app/message_handler"
)

var ResetMessage = message_handler.ResetMessage
var GetLatestReset = http_handler.GetLatestReset