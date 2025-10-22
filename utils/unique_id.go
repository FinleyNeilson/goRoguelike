package utils

import (
	"math/rand"
	"sync"
	"time"
)

var (
	usedIDs = make(map[int]struct{})
	randMu  sync.Mutex
	randSrc = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// UniqueID returns a pseudo-random integer that hasn't been used yet
func UniqueID() int {
	randMu.Lock()
	defer randMu.Unlock()

	for {
		id := randSrc.Int()
		if _, exists := usedIDs[id]; !exists {
			usedIDs[id] = struct{}{}
			return id
		}
	}
}

//func ResetIDs() {
//	randMu.Lock()
//	defer randMu.Unlock()
//	usedIDs = make(map[int]struct{})
//}
