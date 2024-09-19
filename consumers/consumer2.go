package consumers

/**
 * @Author huang
 * @Date 2024-07-12
 * @File: consumer1.go
 * @Description:
 */

import (
	"log"
)

func HandleMessage2(msg []byte) {
	log.Printf("Received a message from queue2: %s", msg)
}
