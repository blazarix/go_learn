package singleton

import "sync"

type Singleton struct{}

var singleton *Singleton
var once sync.Once

func GetLazyInstance() *Singleton {
	once.Do(
		func() {
			singleton = &Singleton{}
		},
	)
	return singleton
}
