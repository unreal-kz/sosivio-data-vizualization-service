package models

import (
	"github.com/jinzhu/gorm"
	"github.com/unreal-kz/sosivio-data-vizualization-service/pkg/config"
)

var db *gorm.DB

type DataVisual struct {
	gorm.Model
	Id          int64   `json:"Id"`
	WhatHappend string  `gorm:""json:"whathappend"`
	Occurences  int64   `json:"occurences"`
	Percentage  float64 `json:"percentage"`
	Description string  `json:"description"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&DataVisual{})
}

func (dv *DataVisual) CreateData() *DataVisual {
	db.NewRecord(dv)
	db.Create(&dv)
	return dv
}

func GetAllData() []DataVisual {
	var dataVisuals []DataVisual
	db.Find(&dataVisuals)
	return dataVisuals
}

func GetDataById(ID int64) (*DataVisual, *gorm.DB) {
	var getData DataVisual
	db := db.Where("ID=?", ID).Find(&getData)
	return &getData, db
}

func DeleteData(ID int64) DataVisual {
	var dataVisual DataVisual
	db.Where("ID=?", ID).Delete(dataVisual)
	return dataVisual
}
