#!/usr/bin/env bash

gcloud functions deploy ResetMessage --trigger-topic messages --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)