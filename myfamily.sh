#!/bin/bash

curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json	| jq --arg ID $HERO_ID '.[] | select(.id == ($ID|tonumber)) | .connections.relatives' | sed 's/"//g'
