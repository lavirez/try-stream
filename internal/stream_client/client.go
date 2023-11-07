package streamclient

import (
	"os"

	"github.com/GetStream/stream-go2/v8"
	"github.com/joho/godotenv"
)

func setUp() (*stream.Client, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	API_KEY := os.Getenv("STREAM_API_KEY")
	SECRET_KEY := os.Getenv("STREAM_SECRET_KEY")

	client, err := stream.New(API_KEY, SECRET_KEY)
	if err != nil {
		return nil, err
	}
	return client, nil
}
