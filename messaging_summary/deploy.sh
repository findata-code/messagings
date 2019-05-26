#!/usr/bin/env bash

gclound functions deploy SummaryMessage --trigger-topic messages --runtime go111 --set-env-vars FV_TOKEN=$(cat ./.token)