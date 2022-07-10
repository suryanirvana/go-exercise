package main

import (
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type User struct {
	Login				string	`json:"login"`
	Name				string	`json:"name"`
	PublicRepos			int		`json:"public_repos"`
}

type UserRequest struct {
	Username 			string
	Url					string
}

func NewUserRequest(username string) (*UserRequest, error) {
	if username == "" {
		return nil, fmt.Errorf("Username can't be nil")
	}

	userRequest := UserRequest{
		Username: username,
		Url: fmt.Sprintf("https://api.github.com/users/%s", username),
	}

	return &userRequest, nil
}

func (request *UserRequest) Get() (*User, error) {
	response, err := http.Get(request.Url)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	defer response.Body.Close()

	var user *User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return user, nil
}

func main() {
	username := "suryanirvana"
	userRequest, _ := NewUserRequest(username)

	user, err := userRequest.Get()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(user)
	}
}
