#!/usr/bin/env bash

gcloud functions deploy LineToPubsub --trigger-http --runtime go111 --set-env-vars CONFIG=$(base64 -i ./current.config.json)