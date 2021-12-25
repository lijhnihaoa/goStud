package socket

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// 获取当前执行文件的绝对路径
func GetAbsPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)

	return filepath.Dir(path)
}

// 加载文件中的账号密码
func LoadUserInfo(filePth string, userinfos map[string]*Userinfo) (map[string]*Userinfo, error) {
	f, err := os.OpenFile(filePth, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		arr := strings.Split(line, "@$")
		userinfos[arr[0]] = &Userinfo{Account: arr[0], Password: arr[1], RegisterTime: arr[2]}
	}
	return userinfos, nil
}

// 每注册一个用户，写入文件保存
func SaveUserInfo(filePth string, ui *Userinfo) error {
	str := fmt.Sprintf("%s@$%s@$%s\n", ui.Account /*用户账号*/, ui.Password /*用户密码*/, ui.RegisterTime /*注册时间*/)
	f, err := os.OpenFile(filePth, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := []byte(str)
	f.Write(buf)
	return nil
}
func SetNickname(conn net.Conn, ch chan string) string {
	ch <- "跳过昵称设置输入: 1, 否则输入: 0"
	var nickname string
	var flag bool
	io1 := bufio.NewScanner(conn)
	for io1.Scan() {
		if !flag && io1.Text() == "1" {
			flag = true
			break
		} else if !flag {
			flag = true
			ch <- "请设置你的昵称: "
			continue
		}
		nickname = io1.Text()
		if !nicknames[nickname] {
			ch <- "确认输入: 1, 重新设置输入: 0"
			for io1.Scan() {
				s := io1.Text()
				if s == "1" {
					lock.Lock()
					nicknames[nickname] = true
					lock.Unlock()
					break
				} else if s == "0" {
					ch <- "请重新设置你的昵称:"
					break
				} else {
					ch <- "确认输入: 1, 重新设置输入: 0"
				}
			}
			if nicknames[nickname] {
				break
			}
		} else {
			ch <- "昵称已被占用，请设置其他昵称 !"
		}

	} // 忽略input.Err错误
	return nickname
}

func RegisterdAccount(conn net.Conn, ch chan string) string {
	ch <- "\t注册新用户"
	var account string
	var password string
	var status string = "creating_account"

	ch <- "请输入账号:"
	input := bufio.NewScanner(conn)
	for input.Scan() {
		s := input.Text()
		switch status {
		case "creating_account":
			if len(s) < 4 {
				continue
			}
			lock.Lock()
			if s != "" && UserInfos[s] == nil {
				account = s
				status = "creating_password"
				ch <- "请输入密码:"
			} else {
				ch <- "输入为空或者账号已存在 !"
				ch <- "请重新输入账号:"
			}
			lock.Unlock()
		case "creating_password":
			if len(s) < 4 {
				ch <- "密码太短, 请重新输入..."
				continue
			} else {
				password = s
				status = "second_password"
				ch <- "再次输入密码"
			}
		case "second_password":
			if password != s {
				ch <- "两次密码不一致，请重新设置密码..."
				status = "creating_password"
			} else {
				ch <- "注册成功"
				goto successful
			}
		}
	}
successful:
	lock.Lock()
	UserInfos[account] = &Userinfo{Account: account, Password: password, RegisterTime: time.Now().Format("2006-01-02 15:04:05"), Online: true}
	SaveUserInfo(DbPth, UserInfos[account])
	lock.Unlock()
	ch <- account + " : " + password
	return account
}
func Login(conn net.Conn, ch chan string) string {
	var count int
	var account string
	var status string = "account"

	ch <- "已有账号直接输入账号，注册请输出入: 1"
	input := bufio.NewScanner(conn)
	for input.Scan() {
		s := input.Text()
		if s == "1" {
			return RegisterdAccount(conn, ch)
		}
		switch status {
		case "account":
			account = s
			status = "password"
			ch <- "输入密码:"
		case "password":
			if count > 3 {
				ch <- "输入次数超过限制..."
				return ""
			}
			count += 1
			if UserInfos[account].Password == s {
				//ch <- "\t登录成功..."
				if UserInfos[account].Online {
					ch <- "\t该账户已经被其他人登录..."
					time.Sleep(1 * time.Second)
					conn.Close()
					return ""
				} else {
					UserInfos[account].Online = true
				}
				ch <- "\t登录成功..."
				goto successful
			} else {
				ch <- "密码错误，请再次输入密码:"
			}
		}
	}
successful:
	return account
}
