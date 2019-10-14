#!/bin/bash

curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json	| jq -j --arg ID $HERO_ID '.[] | select(.id == ($ID|tonumber)) | .connections.relatives'
