package streamclient

import (
	"fmt"
	"os"

	"github.com/GetStream/stream-go2/v8"
	"github.com/joho/godotenv"
)

func SetUp() *stream.Client {
	if err := godotenv.Load(); err != nil {
		panic("Steam has no env")
	}
	API_KEY := os.Getenv("STREAM_API_KEY")
	SECRET_KEY := os.Getenv("STREAM_SECRET_KEY")
	fmt.Println(API_KEY)
	fmt.Println(SECRET_KEY)

	client, err := stream.New(API_KEY, SECRET_KEY)
	if err != nil {
		panic("Steam is not connected")
	}
	return client
}
