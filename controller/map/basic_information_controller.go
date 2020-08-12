package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeffri/golang-test/db"
	"github.com/jeffri/golang-test/initialize"
)

func ReturnAllBasicInformation(w http.ResponseWriter, r *http.Request) {
	var basic initialize.BasicInformation
	var arrBasicInformation []initialize.BasicInformation
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM basic_information")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&basic.Id_basic_information, &basic.Employee_code, &basic.First_name, &basic.Last_name, &basic.Gender, &basic.Birthdate, &basic.Add_postal_code, &basic.Id_prefecture, &basic.Adress, &basic.Adress_kana, &basic.Adress_detail, &basic.Adress_detail_kana, &basic.Add_phone_number, &basic.Marital_status, &basic.Dormitory_status); err != nil {
			log.Fatal(err.Error())

		} else {
			arrBasicInformation = append(arrBasicInformation, basic)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrBasicInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
