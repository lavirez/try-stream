package main

import (
	"context"
	"fmt"

	"github.com/GetStream/stream-go2/v8"
)

type User struct {
	ID       string  `json:"user_id"`
	Email    string  `json:"email"`
	UserName string  `json:"username"`
	Token    *string `json:"token"`
}

func (user *User) addUserToStream(client *stream.Client) {
	userData := stream.User{
		ID: user.ID,
		Data: map[string]any{
			"email":    user.Email,
			"username": user.UserName,
		},
	}
	resp, err := client.Users().Add(context.TODO(), userData, false)
	if err != nil {
		return
	}
	fmt.Println(resp)
}

func (user *User) generateStreamToken(client *stream.Client) { 

}

func decodeJWTToken(client *stream.Client) {
}


func (user *User) generateJWTToknen(client *stream.Client) {
}

func main() {
	u := User{
		ID:       "92903hfj2030asd-asj393ds-a8490asdn-as24jad",
		Email:    "info@alire.me",
		UserName: "alireza",
		Token:    nil,
	}
	u.addUserToStream()
}
