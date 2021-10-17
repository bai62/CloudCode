package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func (this *User)ListenMessage(){
	for{
		msg:= <-this.C
		this.conn.Write([]byte(msg+"\n"))
	}
}

func NewUser(conn net.Conn) *User {

	user := User{
		Name: conn.RemoteAddr().String(),
		Addr: conn.RemoteAddr().String(),
		conn: conn,
		C: make(chan string),
	}

	go user.ListenMessage()

	return &user
}

