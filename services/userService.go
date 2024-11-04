package services

import (
	"onez19/config"
	"onez19/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func LoginUser(user models.User) (string, error) {
	selectedUser := new(models.User)
	row := config.DB.QueryRow("SELECT username, password, first_name, last_name FROM user WHERE username = ?", user.Username)

	err := row.Scan(&selectedUser.Username, &selectedUser.Password, &selectedUser.FirstName, &selectedUser.LastName)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}
