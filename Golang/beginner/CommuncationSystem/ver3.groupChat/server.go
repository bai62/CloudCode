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

func (this*Server)BroadCast(user *User,msg string){
	this.Message<-user.Name+" "+ msg
}


func (this *Server) Handler(conn net.Conn) {
	name := conn.RemoteAddr().String()
	user := NewUser(conn)

	this.MapLock.Lock()
	this.OnlineMap[name]=user
	this.MapLock.Unlock()

	this.BroadCast(user,"is online")

	go func() {
		buf := make([]byte,4096)
		for{
			n, err := conn.Read(buf)
			if n == 0{
				this.BroadCast(user,"user is offline")
				return
			}
			if err != nil && err != io.EOF{
				fmt.Println("conn read err")
				return
			}

			msg := string(buf[:n-1])
			this.BroadCast(user,msg)
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