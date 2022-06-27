package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unreal-kz/sosivio-data-vizualization-service/pkg/models"
	"github.com/unreal-kz/sosivio-data-vizualization-service/pkg/utils"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	newDataVisual := models.GetAllData()
	res, _ := json.Marshal(newDataVisual)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateData(w http.ResponseWriter, r *http.Request) {
	createDS := []models.DataVisual{}
	utils.ParseBody(r, &createDS)
	for _, v := range createDS {
		ds := v.CreateData()
		res, _ := json.Marshal(ds)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetDataById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dsID := vars["id"]
	ID, err := strconv.ParseInt(dsID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	dsDetails, _ := models.GetDataById(ID)
	res, _ := json.Marshal(dsDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func UpdateData(w http.ResponseWriter, r *http.Request) {

// }

func DeleteData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dsID := vars["id"]
	ID, err := strconv.ParseInt(dsID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	dsDetails := models.DeleteData(ID)
	res, _ := json.Marshal(dsDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
