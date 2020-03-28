package client

import (
    "fmt"

    "github.com/fbessez/cavalry/config"
    "github.com/slack-go/slack"
)

var slackClient = getSlackClient()

// SetActive will set the user as active if not already.
func SetActive() {
	isActive := isUserActive()

	if isActive {
		fmt.Println("You are already active.")
		return
	}

	setUserActive()
}

func getSlackClient() (client *slack.Client) {
  return slack.New(config.CONSTANTS.APIToken)
}

func isUserActive() bool {
	presence, err := slackClient.GetUserPresence(config.CONSTANTS.UserID)
  if err != nil {
      fmt.Printf("%s\n", err)
      return false
  }

  fmt.Printf("Online: %t\n", presence.Online)
  return presence.Online
}

func setUserActive() {
	// TODO: Figure out how to automagically set yourself to active
  _,_,_,_ = slackClient.SendMessage("CUX6MEAKV", slack.MsgOptionText("test", false))
}

