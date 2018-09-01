package pubsub

import (
	"fmt"
	"playground/PubSub"
)

func main() {
	events := pubsub.New(5)
	go func() {
		events.Sub("ANIMALS")
	}()
	go func() {
		events.Sub("FISH")
	}()
	events.Pub("hello there", "ANIMALS")
	fmt.Scanln()
}
