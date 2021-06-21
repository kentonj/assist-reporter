#!/bin/bash
commit_sha=$(git rev-parse HEAD)
# build and tag with current commit sha and also latest tag
docker build . -t "kzart94/assist-reporter:$commit_sha" -t "kzart94/assist-reporter:latest"
# push both images
docker push "kzart94/assist-reporter:$commit_sha"
docker push "kzart94/assist-reporter:latest"
