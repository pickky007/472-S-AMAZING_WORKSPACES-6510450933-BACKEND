package services

import (
	"onez19/config"
	"onez19/models"

	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers() ([]models.UserResponse, error) {
	var users []models.UserResponse

	// ดึงข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	rows, err := config.DB.Query("SELECT username, first_name, last_name FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UserResponse
		if err := rows.Scan(&user.Username, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func RegisterUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = config.DB.Exec("INSERT INTO user (username, password, first_name, last_name) VALUES (?, ?, ?, ?)",
		user.Username, hashedPassword, user.FirstName, user.LastName)
	return err
}
