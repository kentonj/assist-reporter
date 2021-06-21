# assist-reporter

This is a little assistant service (hence the assist) that allows me to POST and GET metrics. Currently for a chia farm that i'd like to periodically check up on, but could really be for anything else. It's a Go microservice with a mongo db for writing and reading reported metrics.

This can be deployed with `helm`. See the `scripts` dir for a `build-image.sh` script (this should be run on the same architecture that it will be deployed on, so if you're deploying on a raspberry pi, build it there too). `./scripts/helm-deploy.sh` will deploy both containers (`mongo` and `assist-reporter`).

## Routes

`POST /farm-summary`
```json
{
  "farmingStatus": "Farming",
  "chiaFarmed": 2.0,
  "plotCount": 200,
  "blockRewards": 2.0,
  "expectedTimeToWin": "7 months 2 weeks",
  "activePlotters": 4,
  "passedFilter": 130,
  "proofsFound": 1
}
```

`GET /farm-summary/latest`
```json
{
  "msg": "g2g",
  "id": "insertedMongoId"
}
```
