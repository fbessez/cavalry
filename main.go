package main

import (
    "fmt"

    "github.com/davecgh/go-spew/spew"
    "github.com/slack-go/slack"
)

const (
    ApiToken = "YOUR TOKEN GOES HERE"
    UserId   = "YOUR USER ID GOES HERE"
)

func main() {
    client := getSlackClient()

    getUserInfo(client)
    getUserPresence(client)
}

func getSlackClient() (client *slack.Client) {
    return slack.New(ApiToken)
}

func getUserInfo(client *slack.Client) {
    user, err := client.GetUserInfo(UserId)
    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }

    spew.Dump(user)
}

func getUserPresence(client *slack.Client) {
    presence, err := client.GetUserPresence(UserId)
    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }
    spew.Dump(presence)
}

