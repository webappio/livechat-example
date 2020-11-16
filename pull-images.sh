#!/usr/bin/env bash
find services -maxdepth 2 -name Dockerfile -exec grep -o 'FROM [0-9a-zA-Z:.]*' {} \; | \
  awk '{print $2}' | sort | uniq | xargs -I{} docker pull {}
docker-compose pull
