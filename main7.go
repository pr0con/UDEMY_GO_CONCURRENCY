package main


import(
	"fmt"
	"time"
	"sync"
)

type data struct {
	Mutex *sync.RWMutex
	round map[string]int
}

func newData() *data {
	d := make(map[string]int)
	return &data{ &sync.RWMutex{}, d }
}
	

func (d *data) update(wid string) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	count, ok := d.round[wid]
	if !ok {
		fmt.Println(" some error occured... ");
		return
	}
	d.round[wid] = count + 1	
}

func doWork(wid string, d *data, wg *sync.WaitGroup) {
	for range time.Tick(time.Second * 2) {
		d.update(wid)
	}
	
	wg.Done()
}

func getData(d *data) {
	for range time.Tick(time.Second * 2) {
		d.Mutex.RLock()
			fmt.Println(d);
		d.Mutex.RUnlock()
	}
}

func main() {
	var wg sync.WaitGroup
	
	
	d := newData()
	d.round["one"] = 0
	d.round["two"] = 0
	d.round["three"] = 0	

	go doWork("one",d, &wg); wg.Add(1);
	go doWork("two", d, &wg); wg.Add(1);
	go doWork("three",d, &wg); wg.Add(1);
	
	go getData(d)
	
	wg.Wait()
	fmt.Println("we wont get here")
}


