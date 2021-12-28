package ucclient

import (
	"web/log"

	"google.golang.org/grpc"
)

const PORT = "9001"

func main() {
	log.Init("logs/ucClient", "uc", "[uc] 🎄 ", "info")
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatal(`Fail to dial grpc.Client: %v`, err)
	}
	defer conn.Close()
}
