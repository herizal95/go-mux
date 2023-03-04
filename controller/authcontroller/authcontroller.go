package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/herizal95/gomux/config"
	"github.com/herizal95/gomux/helper"
	"github.com/herizal95/gomux/models/entity"
	"github.com/herizal95/gomux/utils"
	"github.com/jinzhu/gorm"
	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var input entity.Users
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// ambil data users by username
	var user entity.Users
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "username atau password salah"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return

		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJson(w, http.StatusInternalServerError, response)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Passowrd), []byte(input.Passowrd)); err != nil {
		response := map[string]string{"message": "username atau password salah"}
		helper.ResponseJson(w, http.StatusUnauthorized, response)
		return
	}

	// create token jwt
	expTime := time.Now().Add(time.Minute * 5)
	claims := &utils.JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "belajar golang by herry",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasikan alg to signin
	tokenAlg := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sigin token
	token, err := tokenAlg.SignedString(utils.JWT_KEY)

	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// set token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	// login berhasil
	response := map[string]string{"message": "Login Berhasil"}
	helper.ResponseJson(w, http.StatusOK, response)
	return

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
		Username: input.Username,
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

	// hapus token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	// login berhasil
	response := map[string]string{"message": "Logout berhasil"}
	helper.ResponseJson(w, http.StatusOK, response)
	return

}
