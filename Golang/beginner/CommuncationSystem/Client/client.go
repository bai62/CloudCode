package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

var serverIp string
var serverPort int

func init() {
	//定义命令行参数变量，在init函数中进行初始化
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "默认ip为127.0.0.1")
	flag.IntVar(&serverPort, "port", 8888, "默认端口为8888")
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}
	//创建链接
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("dial err")
		return nil
	}
	client.conn = conn

	//返回客户端
	return client
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println(">>>1.group chat")
	fmt.Println(">>>2.private chat")
	fmt.Println(">>>3.rename user name")
	fmt.Println(">>>0.exit")
	fmt.Println(">>>please select a mode")
	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>Please enter a valid number<<<")
		return false
	}
}

func (client *Client) run() {
	for client.flag != 0 {
		for !client.menu() {
		}
		switch client.flag {
		//group chat
		case 1:
			client.publicChat()
			break
		//private chat
		case 2:
			client.privateChat()
			break
		//rename
		case 3:
			client.updateName()
			break
		}

	}
}

func (client *Client) publicChat() {
	fmt.Println(">>>enter chat content and exit when exit")

	var msg string
	fmt.Scanln(&msg)

	for msg != "exit" {
		if len(msg) != 0 {
			_, err := client.conn.Write([]byte(msg + "\n"))
			if err != nil {
				fmt.Println("conn write err")
				break
			}
		}

		fmt.Scanln(&msg)
	}
}

func (client *Client) queryUser() {
	_, err := client.conn.Write([]byte("who\n"))
	if err != nil {
		fmt.Println("query user err")
	}
}

func (client *Client) privateChat() {
	client.queryUser()
	var username string

	var msg string
	fmt.Println(">>>enter chat content and exit when exit")
	fmt.Scanln(&username)

	fmt.Scanln(&msg)
	for username != "exit" {
		for msg != "exit" {
			sengMsg := "to|" + username + "|" + msg + "\n"
			_, err := client.conn.Write([]byte(sengMsg))
			if err != nil {
				fmt.Println("send msg err")
				break
			}
			fmt.Scanln(&msg)
		}
	}

}

func (client *Client) updateName() {
	fmt.Println(">>>enter your new name:")
	fmt.Scanln(&client.Name)
	msg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("rename false")
	}
}

//处理server回应的消息， 直接显示到标准输出即可
func (client *Client) DealResponse() {
	//一旦client.conn有数据，就直接copy到stdout标准输出上, 永久阻塞监听
	io.Copy(os.Stdout, client.conn)
	//for{ 等价于
	//	buf := make()
	//	client.conn.Read(buf)
	//	fmt.Println(buf)
	//}
}

func main() {
	//把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	client := NewClient(serverIp, serverPort)

	go client.DealResponse()

	if client.conn == nil {
		fmt.Println(">>>>>创建客户端失败")
		return
	}

	fmt.Println(">>>>>创建客户端成功")

	client.run()
}
