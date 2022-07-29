package slidingcounterwindow

import (
	"sync"
	"time"
)

var allowedRequests = 3

var counters = make(map[string]map[int64]int)

var mux = sync.Mutex{}

func getNormalizedTS() int64 {
	current := time.Now().Unix()
	return current - (current % 60)
}

//Try to access a resource with implemented Sliding Counter Window API rate limiting
func GetResource(userId string) bool {
	mux.Lock()
	normalizedTs := getNormalizedTS()	
    allowRequest := true
	if userCounters, found := counters[userId]; found {
		if c, found := userCounters[normalizedTs]; found {
			if c >= allowedRequests {
				allowRequest = false
			} else {
				userCounters[normalizedTs] = c + 1
			}
		}
	} else {
		curUserCounter := make(map[int64]int)
		curUserCounter[normalizedTs] = 1
		counters[userId] = curUserCounter	
	}
	mux.Unlock()
	return allowRequest
}

func GetCounters() map[string]map[int64]int {
	return counters
}
