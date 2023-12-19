package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data  map[string]string
	mutex sync.RWMutex
)

func init() {
	data = make(map[string]string)
	data["key1"] = "value1"
	data["key2"] = "value2"
}

func readData(key string) string {
	mutex.RLock()
	defer mutex.RUnlock()
	return data[key]
}

func writeData(key, value string) {
	mutex.Lock()
	defer mutex.Unlock()
	data[key] = value
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Read Data:", readData("key1"))
		}()
	}
	time.Sleep(time.Second) // Let the readers start first

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			writeData("key1", "newvalue1")
			fmt.Println("Write Data: key1 -> newvalue1")
		}()
	}

	wg.Wait()
}
