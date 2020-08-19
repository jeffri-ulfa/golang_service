package List_input_information

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeffri/golang-test/Go_DX_Services/db"
	"github.com/jeffri/golang-test/Go_DX_Services/initialize"
)

func permissionToDrive(w http.ResponseWriter, r *http.Request) {
	var join initialize.Join
	var arrJoin []initialize.Join
	var response initialize.Response

	db := db.Connect()
	result, err := db.Query("SELECT general_information.id_general_information, general_information.id_store_code, store_information.code_store, general_information.id_basic_information, basic_information.employee_code, basic_information.first_name, basic_information.last_name, commuting_basic_information.driver_license_expiry_date,commuting_basic_information.driver_license_expiry_datecar_insurance_document_expiry_date FROM (((general_information INNER JOIN store_information ON general_information.id_store_code = store_information.id_store_code)INNER JOIN basic_information ON general_information.id_basic_information = basic_information.id_basic_information) INNER JOIN commuting_basic_information ON general_information.id_general_information = commuting_basic_information.id_commuting_basic_information")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for result.Next() {
		if err := result.Scan(&join.id_general_information, &join.code_store, &join.employee_code, &join.first_name, &join.last_name, &join.driver_license_expiry_date, &join.car_insurance_document_expiry_date); err != nil {
			log.Fatal(err.Error())

		} else {
			arrJoin = append(arrJoin, join)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrJoin
	w.Header().Set("Content-Type", "Aplication/json", "*")
	json.NewEncoder(w).Encode(response)

}
