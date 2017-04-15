package main

import (
	"fmt"
	"time"
)

const (
	sourceWechat     = "wechat"
	sourceWeibo      = "weibo"
	sourceTextMessge = "textmessage"
)

type msg struct {
	content string
	source  string
}

func wechatReceiver() <-chan *msg {
	c := make(chan *msg)
	go func() {
		c <- &msg{"wechat1", sourceWechat}
		c <- &msg{"wechat2", sourceWechat}
		c <- &msg{"wechat3", sourceWechat}
	}()

	return c
}

func weiboReceiver() <-chan *msg {
	c := make(chan *msg)
	go func() {
		c <- &msg{"weibo1", sourceWeibo}
		c <- &msg{"weibo2", sourceWeibo}
		c <- &msg{"weibo3", sourceWeibo}
	}()

	return c
}

func textmessageReceiver() <-chan *msg {
	c := make(chan *msg)
	go func() {
		c <- &msg{"textmessage1", sourceTextMessge}
		c <- &msg{"textmessage2", sourceTextMessge}
		c <- &msg{"textmessage3", sourceTextMessge}
	}()

	return c
}

func serviceAggregation(ins ...<-chan *msg) <-chan *msg {
	out := make(chan *msg)
	for _, c := range ins {
		go func(c <-chan *msg) {
			for v := range c {
				out <- v
			}
		}(c)
	}

	return out
}

func main() {
	c := serviceAggregation(weiboReceiver(), wechatReceiver(), textmessageReceiver())

	timeout := time.After(time.Second * 5)
	for {
		select {
		case m := <-c:
			fmt.Println(m.content, m.source)

		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
}
