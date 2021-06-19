package main

import (
	"context"
	"log"
	"time"

	"assist/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client

func postFarmSummary(c *gin.Context) {
	var farmSummaryRequest models.FarmSummaryRequestBody
	err := c.ShouldBindJSON(&farmSummaryRequest)
	if err != nil {
		panic(err)
	}
	fm := models.NewFarmSummary(farmSummaryRequest)
	farmStatusCollection := DB.Database("assist").Collection("farm-status")
	res, err := farmStatusCollection.InsertOne(context.TODO(), fm)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote farm summary with id: %s\n", res.InsertedID)
	c.JSON(200, gin.H{
		"message": "g2g",
		"id":      res.InsertedID,
	})
}

func getLatestFarmSummary(c *gin.Context) {
	farmStatusCollection := DB.Database("assist").Collection("farm-status")
	var farmSummary models.FarmSummary
	latestFarmStatusOpts := options.FindOne()
	latestFarmStatusOpts.SetSort(bson.M{"$natural": -1})
	err := farmStatusCollection.FindOne(context.TODO(), bson.M{}, latestFarmStatusOpts).Decode(&farmSummary)
	if err != nil {
		log.Fatal(err)
	}
	location, err := time.LoadLocation("America/Denver")
	farmSummary.Timestamp = farmSummary.Timestamp.In(location)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, farmSummary)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	credential := options.Credential{
		Username: "root",
		Password: "secret",
	}
	clientOpts := options.Client().ApplyURI("mongodb://mongo:27017").SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	DB = client
	r := gin.Default()

	r.POST("/farm-summary", postFarmSummary)
	r.GET("/farm-summary/latest", getLatestFarmSummary)
	r.Run() // listen and serve on 0.0.0.0:8080
}
