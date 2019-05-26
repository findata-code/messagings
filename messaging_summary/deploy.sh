#!/usr/bin/env bash

gcloud functions deploy GetSummaryMessage --trigger-topic messages --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)