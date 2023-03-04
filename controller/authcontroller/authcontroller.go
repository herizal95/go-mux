package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/herizal95/gomux/config"
	"github.com/herizal95/gomux/helper"
	"github.com/herizal95/gomux/models/entity"
	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {

	var input entity.Users
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Passowrd), bcrypt.DefaultCost)

	data := entity.Users{
		Uid:      uuid.NewV4(),
		Name:     input.Name,
		Email:    input.Email,
		Passowrd: string(hashPassword),
	}

	// insert ke database
	if err := config.DB.Create(&data).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJson(w, http.StatusOK, response)

}

func Logout(w http.ResponseWriter, r *http.Request) {

}
