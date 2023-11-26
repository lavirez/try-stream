package users

import (
	"context"

	"github.com/GetStream/stream-go2/v8"
)

type User struct {
	ID       string  `json:"user_id"`
	Email    string  `json:"email"`
	UserName string  `json:"username"`
	Token    *string `json:"token"`
}

func (user *User) retrieveStreamUser(client *stream.Client) *stream.UserResponse { 
    resp, err := client.Users().Get(context.TODO(), user.ID)
    if err != nil { 
        return nil
    }
    return resp
}

func (user *User) addUserToStream(client *stream.Client) *stream.UserResponse {
	userData := stream.User{
		ID: user.ID,
		Data: map[string]any{
			"email":    user.Email,
			"username": user.UserName,
		},
	}
	resp, err := client.Users().Add(context.TODO(), userData, false)
	if err != nil {
		return nil
	}
	return resp
}

func (user *User) getStreamToken(client *stream.Client) *string { 
    userToken, err := client.CreateUserToken(user.ID)
    if err != nil { 
        return nil
    }
    return &userToken
}

func decodeJWTToken(client *stream.Client) {
}

func (user *User) generateJWTToknen(client *stream.Client) {
}

