package main

func channelTest(ch chan<- *int) {
	var a = 2
	ch <- &a
	//<-ch 报错
}
