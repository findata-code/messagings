#!/usr/bin/env bash
gcloud functions deploy DeleteExpenseMessage --trigger-topic messages --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)
gcloud functions deploy DeleteIncomeMessage --trigger-topic messages --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)