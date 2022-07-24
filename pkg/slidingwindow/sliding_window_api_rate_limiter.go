package slidingwindow

import (
	"sync"
	"time"
)

var windowSize int64 = 60
var allowedRequests = 3

var slidingWindowUsers = make(map[string][]int64)

var mux = sync.Mutex{}

func removeOutDated (arr []int64, current int64, wS int64) []int64 {
	idx := 0
	for i:=0; i<len(arr); i++ {
		if current - arr[i] <= wS {
			idx = i
			break
		}
	}
	return arr[idx:]
}

//Try to access a resource with implemented Sliding Window API rate limiting
func GetResource(userId string) bool {
	mux.Lock()
	current := time.Now().Unix()
	if v, found := slidingWindowUsers[userId]; found {
		v = removeOutDated(v, current, windowSize)	
		v = append(v, []int64 {current} ...)
		slidingWindowUsers[userId] = v
		mux.Unlock()
		return len(v) < allowedRequests
	} else {
		slidingWindowUsers[userId] = []int64 {current}
		mux.Unlock()
		return true
	}
}

func GetSlidingWindowUsers() map[string][]int64 {
	return slidingWindowUsers
}
