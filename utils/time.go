package utils

import (
	"errors"
	"time"
)

// # Get Current System Time
func GetCurrentTime() (*time.Time, error) {
	// # Get Current Time
	now := time.Now()

	// # IST Timezone Location
	timeZone := "Asia/Kolkata"

	// # Get Timezone Location
	ISTLocation, err := time.LoadLocation(timeZone)
	if err != nil {
		err = errors.New("error: unable to load location" + "\n" + err.Error())
		return nil, err
	}

	// # Convert Time to IST
	time := now.In(ISTLocation)

	// # Return IST Time
	return &time, nil
}
