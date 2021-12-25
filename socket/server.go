package socket

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var nicknames = make(map[string]bool)
var online int = 0
var lock sync.Locker = &sync.Mutex{}
var IP string = "0.0.0.0"

var UserInfos = make(map[string]*Userinfo)
var DbPth string

type client chan<- string //传出消息通道

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // 所有来自客户端的消息
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients { //将所有客户端的通道轮训发送消息
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {

	ch := make(chan string)   //创建一个通道 无缓存
	go clientWriter(conn, ch) //输出客户端发来的消息

	who := conn.RemoteAddr().String() // 远端地址
	ch <- "You are " + who
	account := Login(conn, ch)
	if account == "" {
		return
	}
	nickname := SetNickname(conn, ch)
	if nickname != "" {
		messages <- who + " has arrived, his nickname is " + nickname
		who = nickname
	} else {
		messages <- who + " has arrived, his nickname is " + who
	}
	entering <- ch
	lock.Lock()
	online += 1
	lock.Unlock()
	ch <- "当前在线人数: " + strconv.Itoa(online)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	} // 忽略input.Err错误
	leaving <- ch
	if UserInfos[account].Online {
		UserInfos[account].Online = false
	}
	messages <- who + " has left"
	lock.Lock()
	online -= 1
	if nickname != "" {
		delete(nicknames, nickname)
	}
	lock.Unlock()
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	// start := time.Now()
	// sesc := time.Since(start).Seconds()
	// if sesc >= 300 {
	// 	conn.Close()
	// }
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE : 忽略网络错误
	}
}

// 1.监听一个地址 端口
// 2.等待连接
// 3.连接成功后处理
func StartServer() {
	var err error
	DbPth = GetAbsPath() + "/db/db.txt"
	fmt.Println(DbPth)
	UserInfos, err = LoadUserInfo(DbPth, UserInfos)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range UserInfos {
		fmt.Println(k, v)
	}
	listener, err := net.Listen("tcp", IP+":9009") //监听一个地址
	if err != nil {
		log.Fatal(err)
		return
	}
	go broadcaster()

	for {
		conn, err := listener.Accept() // 等待client的connect
		if err != nil {
			log.Print(err)
			time.Sleep(50 * time.Microsecond)

			continue
		}
		go handleConn(conn) //给每个连接启动一个go程处理
	}
}
