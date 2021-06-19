package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FarmSummaryRequestBody struct {
	FarmingStatus     string  `json:"farmingStatus"`
	ChiaFarmed        float64 `json:"chiaFarmed"`
	PlotCount         int     `json:"plotCount"`
	BlockRewards      float64 `json:"blockRewards"`
	ExpectedTimeToWin string  `json:"expectedTimeToWin"`
	ActivePlotters    int     `json:"activePlotters"`
	PassedFilter      int     `json:"passedFilter"`
	ProofsFound       int     `json:"proofsFound"`
}

type FarmSummary struct {
	Id                primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	FarmingStatus     string             `json:"farmingStatus"`
	ChiaFarmed        float64            `json:"chiaFarmed"`
	PlotCount         int                `json:"plotCount"`
	BlockRewards      float64            `json:"blockRewards"`
	ExpectedTimeToWin string             `json:"expectedTimeToWin"`
	ActivePlotters    int                `json:"activePlotters"`
	PassedFilter      int                `json:"passedFilter"`
	ProofsFound       int                `json:"proofsFound"`
	Timestamp         time.Time          `json:"timestamp"`
}

func NewFarmSummary(fsr FarmSummaryRequestBody) *FarmSummary {
	id := primitive.NewObjectID()
	dt := time.Now()
	fs := FarmSummary{
		Id:                id,
		FarmingStatus:     fsr.FarmingStatus,
		ChiaFarmed:        fsr.ChiaFarmed,
		PlotCount:         fsr.PlotCount,
		BlockRewards:      fsr.BlockRewards,
		ExpectedTimeToWin: fsr.ExpectedTimeToWin,
		ActivePlotters:    fsr.ActivePlotters,
		PassedFilter:      fsr.PassedFilter,
		ProofsFound:       fsr.ProofsFound,
		Timestamp:         dt,
	}
	return &fs
}
