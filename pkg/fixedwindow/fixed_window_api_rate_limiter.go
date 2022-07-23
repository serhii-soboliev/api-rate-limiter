package fixedwindow

import (
	"fmt"
	"time"
)

type FixedWindowUser struct {
	UserID     string `json:"userid"`
	Last       int64 `json:"last"`
	Count      int `json:"count"`
	WindowSize int64 `json:"windowsize"`
}

var windowSize int64 = 60

var fixedWindowUsers = make(map[string]FixedWindowUser)

func GetResource(userId string) bool {
	if v, found := fixedWindowUsers[userId]; found {
		last := v.Last
		current := time.Now().Unix()
		if current - last > windowSize {
			fwu := FixedWindowUser{userId, time.Now().Unix(), 1, windowSize}
			fixedWindowUsers[userId] = fwu	
			return true
		} else if v.Count < 3 {
			fmt.Println(v)
			v.Count = v.Count + 1
			fixedWindowUsers[userId] = v
			return true
		} else {
			return false
		}	
	} else {
		fwu := FixedWindowUser{userId, time.Now().Unix(), 1, windowSize}
		fixedWindowUsers[userId] = fwu
		return true
	}
}

func GetFixedWindowUsers() map[string]FixedWindowUser {
	return fixedWindowUsers
}
