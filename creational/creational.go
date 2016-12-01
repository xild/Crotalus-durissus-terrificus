package creational

import "sync"

var (
	once     sync.Once
	instance singleton
)

type singleton map[string]string

//Using the sync package!
//once.Do only call a method one time
//in a life period.
//So if you call NewSigleton() a multiple of time
//only first time will execute de Do function.
func Singleton() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}
