package fixedwindow

import (
	"fmt"
	"sync"
	"time"
)

type FixedWindowUser struct {
	UserID     string `json:"userid"`
	Last       int64  `json:"last"`
	Count      int    `json:"count"`
	WindowSize int64  `json:"windowsize"`
}

var windowSize int64 = 60
var allowedRequests = 3

var fixedWindowUsers = make(map[string]FixedWindowUser)

var mux = sync.Mutex{}

func GetResource(userId string) bool {
	mux.Lock()
	if v, found := fixedWindowUsers[userId]; found {
		last := v.Last
		current := time.Now().Unix()
		if current-last > windowSize {
			fwu := FixedWindowUser{userId, time.Now().Unix(), 1, windowSize}
			fixedWindowUsers[userId] = fwu
			mux.Unlock()
			return true
		} else if v.Count < allowedRequests {
			fmt.Println(v)
			v.Count = v.Count + 1
			fixedWindowUsers[userId] = v
			mux.Unlock()
			return true
		} else {
			fmt.Println(v)
			mux.Unlock()
			return false
		}
	} else {
		fwu := FixedWindowUser{userId, time.Now().Unix(), 1, windowSize}
		fixedWindowUsers[userId] = fwu
		mux.Unlock()
		return true
	}
}

func GetFixedWindowUsers() map[string]FixedWindowUser {
	return fixedWindowUsers
}
