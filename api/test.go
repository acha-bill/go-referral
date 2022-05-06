package api

import "sync"

var once sync.Once

func startAPI() {
	once.Do(func() {
		Start()
	})
}
