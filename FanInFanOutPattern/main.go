package main

import (
	"fmt"
	"sync"
	"time"
)

type TChannel struct {
	Channel      chan string
	ChannelDones chan struct{}
	Channels     map[string]chan string
	sync.RWMutex
}

func (c *TChannel) AddChannel(key string, ch chan string) {
	c.Lock()
	c.Channels[key] = ch
	c.Unlock()
}

func (c *TChannel) DeleteChannel(key string) {
	c.Lock()
	delete(c.Channels, key)
	c.Unlock()
}

var workerId int = 0

func (c *TChannel) Publisher() TChannel {
	c = new(TChannel)
	c.Channel = make(chan string)
	c.ChannelDone = make(chan struct{})
	c.AddChannel()
	return c
}

func (c *TChannel) SendMessage(pub chan string, oneMessage string) {
	go func() {
		pub <- oneMessage
	}()
}

func Subscriber(pub chan string, done chan struct{}) {
	workerId++
	go func(id int) {
		fmt.Println("Subscribe start. ", id)
	Loop:
		for {
			select {
			case ev := <-pub:
				fmt.Println(ev, id)
			case <-done:
				break Loop
			}
		}
		fmt.Println("Subscribe exit. ", id)
	}(workerId)
}

func main() {
	channelOne, done := Publisher()
	Subscriber(channelOne, done)
	Subscriber(channelOne, done)
	Subscriber(channelOne, done)
	time.Sleep(1 * time.Second)
	SendMessage(channelOne, "First Message")
	SendMessage(channelOne, "First Message")
	SendMessage(channelOne, "First Message")
	time.Sleep(2 * time.Second)
	fmt.Println("Complete")
	done <- struct{}{}
	done <- struct{}{}
	done <- struct{}{}
	time.Sleep(2 * time.Second)
	fmt.Println("Exit Program")
}
