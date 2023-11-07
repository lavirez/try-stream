package main

import ( 
    "os"

    "github.com/joho/godotenv"
    "github.com/GetStream/stream-go2/v8"
)

func setUp() { 
    if err := godotenv.Load(); err != nil { 
        panic("Error Loading .env variable")
        return
    }
    API_KEY := os.Getenv("STREAM_API_KEY")
    SECRET_KEY := os.Getenv("STREAM_SECRET_KEY")

    client, err := stream.New(API_KEY, SECRET_KEY)
    if err != nil { 
        panic(err)
    }
}

