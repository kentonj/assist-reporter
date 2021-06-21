#!/bin/bash
image_tag=$(git rev-parse HEAD)
repository="kzart94/assist-reporter"
# build and tag with current commit sha and also latest tag
docker build . -t "$repository:$image_tag" -t "$repository:latest"
# push both images
docker push "$repository:$image_tag"
docker push "$repository:latest"
