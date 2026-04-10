package workers

import "time"

func StartScanner() {
	ticker := time.NewTicker(time.Second * 30)
	for range ticker.C {
		scan()
	}
}

func scan() {

}
