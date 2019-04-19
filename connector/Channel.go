package connector

import "log"

type IChannel interface {
	Run()
	Publish(content interface{})
}

type channel struct {
	ctype      string
	register   chan *ISession
	unregister chan *ISession
	sessions   map[*ISession]bool
}

func NewChannel(ctype string) IChannel {
	return &channel{ctype: ctype, register: make(chan *ISession), unregister: make(chan *ISession), sessions: make(map[*ISession]bool)}
}

func (ch channel) Run() {
	for {
		select {
		case s := <-ch.register:
			if ch.sessions[s] {
				log.Fatalf("Duplicated session:%s", ISession(*s).String())
				break
			}
			ch.sessions[s] = true
		case s := <-ch.unregister:
			if !ch.sessions[s] {
				log.Fatalf("Not registered yet,session:%s", ISession(*s).String())
			}
		default:

		}
	}
}

func (ch channel) Publish(content interface{}) {

}
