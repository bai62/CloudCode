package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	ip        string
	port      int
	OnlineMap map[string]*User
	MapLock   sync.RWMutex
	Message   chan string
}

func (this*Server)BroadCast(user *User){
	this.Message<-fmt.Sprintf("%s上线",user.Name)
}

func (this *Server) Handler(conn net.Conn) {
	name := conn.RemoteAddr().String()
	user := NewUser(conn)

	this.MapLock.Lock()
	this.OnlineMap[name]=user
	this.MapLock.Unlock()

	this.BroadCast(user)

	go func() {
		buf := make([]byte,4096)
		for{
			msg,err := conn.Read(buf)
			if err != nil || err != io.EOF{
				fmt.Println("conn read err")
			}
		}
	}()

}

func (this *Server) ServerListen() {

	for {
		msg := <-this.Message
		this.MapLock.Lock()
		for _, user := range this.OnlineMap {
			user.C <- msg
		}
		this.MapLock.Unlock()
	}
}

func NewServer(ip string, port int) *Server {
	server := Server{
		ip:        ip,
		port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return &server
}

func (this *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.ip, this.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	go this.ServerListen()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn err")
			continue
		}

		go this.Handler(conn)

	}
}