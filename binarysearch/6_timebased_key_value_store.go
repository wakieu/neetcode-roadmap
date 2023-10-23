package binarysearch

// https://leetcode.com/problems/time-based-key-value-store/

import "sort"

type Rec struct {
	value string
	timestamp int
}

type TimeMap struct {
    Records map[string][]Rec
}

func Constructor() TimeMap {
    return TimeMap{make(map[string][]Rec)}
}

func (tm *TimeMap) Set(key string, value string, timestamp int)  {
    if _, ok := tm.Records[key]; !ok {
        tm.Records[key] = make([]Rec, 0)
    }
    tm.Records[key] = append(tm.Records[key] , Rec{value, timestamp})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
    i := sort.Search(len(tm.Records[key]), func (i int) bool {
        return tm.Records[key][i].timestamp > timestamp
    })

    if i == 0 {
        return ""
    }

    return tm.Records[key][i-1].value
}