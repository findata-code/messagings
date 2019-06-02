#!/usr/bin/env bash
gcloud functions deploy DeleteExpenseMessage --trigger-topic messages --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)