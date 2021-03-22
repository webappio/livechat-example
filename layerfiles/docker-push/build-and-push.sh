#!/usr/bin/env bash

set -eu -o pipefail

# store old images for removal at end (to avoid running out of disk)
old_images="$(docker images -a | tail -n +2 | grep -v '<none>' | awk '{print $1 ":" $2}' | sort | uniq)"

export COMPOSE_DOCKER_CLI_BUILD=1
export DOCKER_BUILDKIT=1

docker-compose build --parallel --progress=plain

# clear unused docker data to avoid running out of disk over time
if [ -n "${old_images}" ]; then
  echo "Deleting old images..."
  docker rmi -f ${old_images} || true
fi

export IMAGE_TAG=$(git describe --tags --always)

BUILT_IMAGES="$(docker images | awk '{print $1 ":" $2}' | tail -n +2 | grep -v '<none>' | sort | uniq || true)"
# PUSH_JOBS=()
for img in $BUILT_IMAGES; do
  shortimg="$(echo $img | awk -F/ '{print $NF}' | awk -F: '{print $1}')"
  docker tag $img $REGISTRY/$shortimg:$IMAGE_TAG
  # docker push $REGISTRY/$shortimg:$IMAGE_TAG&
  # PUSH_JOBS+=($!)
done

for i in {1..2}; do
  echo "Clearing old containers if applicable..."
  timeout 10 docker container prune --force || true
  echo "Clearing old volumes if applicable..."
  timeout 10 docker volume prune --force || true
  echo "Clearing old images if applicable..."
  timeout 10 docker image prune --force || true
  echo "Clearing build cache if appplicable..."
  timeout 10 docker builder prune --keep-storage 10G --force || true
done

# echo "Waiting for push jobs to complete..."
# for job in ${PUSH_JOBS[@]}; do
#   wait $job
# done
echo "Done!"
