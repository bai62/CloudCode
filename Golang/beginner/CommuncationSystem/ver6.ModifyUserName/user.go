package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
	server *Server
}

func (this *User)ListenMessage(){
	for{
		msg:= <-this.C
		this.conn.Write([]byte(msg+"\n"))
	}
}

func (this *User)Online(){

	this.server.MapLock.Lock()
	this.server.OnlineMap[this.Name]=this
	this.server.MapLock.Unlock()

	this.server.BroadCast(this,"is online")
}

func (this *User)Offline(){
	this.server.BroadCast(this,"user is offline")
}

func (this *User)DoMessage(msg string){
	this.server.BroadCast(this,msg)
}

func NewUser(conn net.Conn,server *Server) *User {

	user := User{
		Name: conn.RemoteAddr().String(),
		Addr: conn.RemoteAddr().String(),
		conn: conn,
		C: make(chan string),
		server: server,
	}

	go user.ListenMessage()

	return &user
}

