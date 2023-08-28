#!/bin/bash
curl -s https://techmancer.com/api/challenges/logs/web-logs-raw -O >/dev/null
echo -e "\n" >>web-logs-raw

substr='techmancer heroku/router'
masked="MASKED"
while read p; do
    if [[ "$p" == *"$substr"* ]]; then

        request_id=$(echo "$p" | cut -d "=" -f7 | cut -d " " -f1)

        if [[ "$p" == *"$masked"* ]]; then
            fwd="M"
        else
            fwd=$(echo "$p" | cut -d "=" -f8 | cut -d " " -f1 | tr -d '"')
        fi
        echo "$request_id [$fwd]"
    fi
done <web-logs-raw
