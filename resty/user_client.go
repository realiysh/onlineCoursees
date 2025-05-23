package resty

import (
	"course-service/models" // <-- используй то, что указано в go.mod
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func Useruser(token string) (*models.User, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetHeader("Accept", "application/json").
		Get("http://localhost:8084/api/profile")

	fmt.Println("StatusCode:", resp.StatusCode())
	fmt.Println("Response Body:", resp.String())
	fmt.Println("Error:", err)

	if err != nil || resp.StatusCode() != 200 {
		return nil, fmt.Errorf("ошибка запроса: %v, статус: %d", err, resp.StatusCode())
	}

	var user models.User
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, err
	}

	return &user, nil
}
