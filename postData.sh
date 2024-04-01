#! /bin/bash

if [ $# -eq 2 ]; then
	curl -X POST http://localhost:8080/post/m5 -H "Content-Type: application/json" -d '{"uid": "'"$1"'", "area": "'"$2"'"}'
else
	echo "pls ./postData.sh [uid] [aria]"
fi