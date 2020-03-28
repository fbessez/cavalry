package main

import (
    "fmt"
    "time"

    "github.com/fbessez/cavalry/config"
    "github.com/fbessez/cavalry/client"
)

const (
    hoursInDay = 24 
    refreshPeriod = 30 // minutes
)

func main() {
    for {
        if isWorkingHours() {
            client.SetActive()
            fmt.Printf("Sleeping for %d minutes...", refreshPeriod)
            time.Sleep(refreshPeriod * time.Minute)
        } else {
            hoursTilStart := hoursTilStart()
            fmt.Printf("Sleeping for %d hours...", hoursTilStart)

            time.Sleep(time.Duration(hoursTilStart) * time.Hour)
        }
    }
}

func isWorkingHours() bool {
    currentHour := time.Now().Hour()

    return currentHour >= config.CONSTANTS.StartHour && currentHour <= config.CONSTANTS.EndHour
}

func hoursTilStart() int {
    currentHour := time.Now().Hour()

    if currentHour < hoursInDay {
        return hoursInDay - currentHour + config.CONSTANTS.StartHour
    }

    return config.CONSTANTS.StartHour - currentHour
}
