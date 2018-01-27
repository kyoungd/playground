package FanInPattern

func fanin(chs ...<-chan in) <- chan in {
	out := make(chan in)
	for _, c := range chs {
		go registerChannel(c, out)
	}
	return out
}

func registerChannel(c <- chan int, out chan<- int) {
	for n := range c {
		out <- n
	}
}
