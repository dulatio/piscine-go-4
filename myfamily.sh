#!/bin/bash

curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json	| jq -r --arg ID $HERO_ID '.[] | select(.id == ($ID|tonumber)) | .connections.relatives'
