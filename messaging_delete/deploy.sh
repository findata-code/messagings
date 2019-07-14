#!/usr/bin/env bash
gcloud functions deploy DeleteExpenseMessage --trigger-topic messages --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)
gcloud functions deploy DeleteIncomeMessage --trigger-topic messages --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)