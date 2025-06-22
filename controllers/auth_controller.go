package controllers

import (
	"encoding/json"
	"net/http"
	"belajar-golang-api/config"
	"belajar-golang-api/models"
	"belajar-golang-api/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Identifier string `json:"identifier"` // Can be email or username
		Password   string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	var user models.User
	result := config.DB.Where("email = ? OR username = ?", userInput.Identifier, userInput.Identifier).First(&user)
	if result.Error != nil {
		utils.RespondError(w, http.StatusUnauthorized, "User not found")
		return
	}

	if !utils.CheckPasswordHash(userInput.Password, user.Password.String) {
		utils.RespondError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := map[string]interface{}{
		"token": token,
		"user": user,
	}

	utils.RespondJSON(w, http.StatusOK, response)
}

func CheckDataHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	for i := range users {
		users[i].SetDefaults()
	}

	utils.RespondJSON(w, http.StatusOK, users)
}
