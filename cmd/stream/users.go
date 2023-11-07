package main

import "github.com/GetStream/stream-go2/v8"

type User struct {
    userName string
    firstName string
    lastName string
}

func (user User) GenerateToken(client *stream.Client) error { 
    return nil
}

