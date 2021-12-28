package pkg

import (
	"time"
)

func super() {
	go time.Sleep(1 * time.Second)
}
