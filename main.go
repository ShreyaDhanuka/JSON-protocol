// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	socketio "github.com/googollee/go-socket.io"
// )

// func main() {
// 	fmt.Println("Hello World")

// 	server, err := socketio.NewServer(nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	server.On("connection", func(so socketio.Socket) {
// 		log.Println("New Connection")
// 	})

// 	http.Handle("/socket.io/", server)
// 	log.Fatal(http.ListenAndServe(":5000", nil))
// }

package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/graarh/golang-socketio/transport"
)

func doSomethingWith(c *gosocketio.Client, wg *sync.WaitGroup) {
	if res, err := c.Ack("join", "This is a client", time.Second*3); err != nil {
		log.Printf("error: %v", err)
	} else {
		log.Printf("result %q", res)
	}
	wg.Done()
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("127.0.0.1", 3003, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected")
	})
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go doSomethingWith(c, wg)
	wg.Wait()
	log.Printf("Done")
}
