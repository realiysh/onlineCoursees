package resty

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type User struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func GetUserByID(id uint) (*User, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("http://localhost:8080/api/users/%d", id))

	if err != nil || resp.StatusCode() != 200 {
		return nil, fmt.Errorf("user not found")
	}

	var user User
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, err
	}
	return &user, nil
}
