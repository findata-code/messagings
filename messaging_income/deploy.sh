#!/usr/bin/env bash

gcloud functions deploy IncomeMessage --trigger-topic messages --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)
gcloud functions deploy GetIncome --trigger-http --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)