#!/bin/bash
# helm deploy the services
image_tag=$(git rev-parse HEAD)

helm upgrade --install mongo charts/mongo
helm upgrade --install assist-reporter charts/assist-reporter --set image.tag="$image_tag"
