#!/usr/bin/env bash

gcloud functions deploy GetSummaryMessage --trigger-topic messages --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)