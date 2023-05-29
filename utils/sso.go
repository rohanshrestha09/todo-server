package utils

import (
	"encoding/json"
	"net/http"

	"github.com/rohanshrestha09/todo/models"
)

type User interface {
	models.FacebookUser | models.GoogleUser
}

func GetSSOUserInfo[T User](token, url string) (T, error) {

	var user T

	request, _ := http.NewRequest("GET", url+token, nil)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return user, err
	}

	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(&user)

	defer response.Body.Close()

	if err != nil {
		return user, err
	}

	return user, nil
}
