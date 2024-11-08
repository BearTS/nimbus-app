#!/bin/bash

VERSION=$(grep -o '\"version\": \"[^\"]*\"' package.json | sed 's/[^0-9a-z.-]//g'| sed 's/version//g')
LATEST="latest"

# if branch is unstable in git for circle ci
if [ -n "$CIRCLE_BRANCH" ]; then
  if [ "$CIRCLE_BRANCH" != "master" ]; then
    LATEST="$LATEST-$CIRCLE_BRANCH"
  fi
fi


sh build.sh

docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --tag nimbus-server:$VERSION \
  --tag nimbus-server:$LATEST \
  --push \
  .
