#!/usr/bin/env bash

gcloud functions deploy ResetMessage --trigger-topic messages --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)
gcloud functions deploy GetLatestReset --trigger-http --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)