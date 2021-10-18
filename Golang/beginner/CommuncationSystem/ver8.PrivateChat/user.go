package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

func (this *User) Online() {

	this.server.MapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.MapLock.Unlock()

	this.server.BroadCast(this, "is online")
}

func (this *User) Offline() {
	this.server.MapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.MapLock.Unlock()

	this.server.BroadCast(this, "user is offline")
}

//给当前User对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//处理信息
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		this.server.MapLock.Lock()
		for _, user := range this.server.OnlineMap {
			oUser := user.Name + " is online\n"
			this.SendMsg(oUser)
		}
		this.server.MapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		name := strings.Split(msg, "|")[1]
		_, ok := this.server.OnlineMap[name]
		if ok {
			this.SendMsg("this user name is already used")
		} else {
			this.server.MapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[name] = this
			this.server.MapLock.Unlock()
			this.Name = name
			this.SendMsg("You have modified you user name successfully\n")
		}

	} else if len(msg) > 4 && msg[:3] == "to|" {
		strs :=strings.Split(msg,"|")
		name := strs[1]
		user,ok := this.server.OnlineMap[name]
		if !ok{
			this.SendMsg("User doesn't exist")
		}else{
			user.SendMsg(this.Name+"to you:"+strs[2]+"\n")
		}

	} else {
		this.server.BroadCast(this, msg)
	}
}

func NewUser(conn net.Conn, server *Server) *User {

	user := User{
		Name:   conn.RemoteAddr().String(),
		Addr:   conn.RemoteAddr().String(),
		conn:   conn,
		C:      make(chan string),
		server: server,
	}

	go user.ListenMessage()

	return &user
}
