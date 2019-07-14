#!/usr/bin/env bash

gcloud functions deploy ResetMessage --trigger-topic messages --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)
gcloud functions deploy GetLatestReset --trigger-http --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)