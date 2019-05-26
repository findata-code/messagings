#!/usr/bin/env bash

gcloud functions deploy LineToPubsub --trigger-http --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)