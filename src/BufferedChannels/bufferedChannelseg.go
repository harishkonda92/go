package BufferedChannels

import (
	"fmt"
	"sync"
	"time"
)

// channels and waitgroups
func process(i int, wg *sync.WaitGroup)  {
	fmt.Println("started wait group go routine", i)
	time.Sleep(2* time.Second)
	fmt.Println("Wait group is ended", i)
	wg.Done()
}

func Mainn()  {
	count := 3
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(i)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("the wait group has beeen ended")
}

