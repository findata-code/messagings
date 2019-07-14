#!/usr/bin/env bash

gcloud functions deploy IncomeMessage --trigger-topic messages --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)
gcloud functions deploy GetIncome --trigger-http --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)
gcloud functions deploy DeleteIncome --trigger-http --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)
gcloud functions deploy GetLastNIncomes --trigger-http --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)