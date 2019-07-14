#!/usr/bin/env bash

gcloud functions deploy HelpMessage --trigger-topic messages --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)