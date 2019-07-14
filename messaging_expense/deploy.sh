#!/usr/bin/env bash

gcloud functions deploy ExpenseMessage --trigger-topic messages --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)
gcloud functions deploy GetExpense --trigger-http --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)
gcloud functions deploy DeleteExpense --trigger-http --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)
gcloud functions deploy GetLastNExpenses --trigger-http --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)