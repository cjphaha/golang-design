package singleton

import "sync"

var (
	once = &sync.Once{}
)

func GetInstanceDc() *Singleton {
	if singleton == nil {
		once.Do(func() {
			singleton = &Singleton{}
		})
	}
	return singleton
}
