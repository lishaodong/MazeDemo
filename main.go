// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// This sample is about using long polling and WebSocket to build a web-based chat room based on beego.
package main

import (
	"github.com/lishaodong/MazeDemo/internal"
	"github.com/lishaodong/Taipei-Torrent/torrent"
	"strings"
	"fmt"
	"net"
	"github.com/astaxie/beego"
	"github.com/lishaodong/MazeDemo/models"
	"os/exec"
)

var(
	launcher *torrent.Launcher
	local string
	l *internal.Launcher
)


func main() {
	l=internal.NewLauncher()
	go l.Run()

	local = GetAddr()


	go LaunchServer()
	OpenPage()
	<-make(chan bool)
}

func LaunchServer(){
	ss := make([]string, 10)
	ss[0] = "./configosst.torrent"
	launcher = torrent.NewLauncher(ss[0:1])

	go launcher.Launch()

	go func(){
		var peer string
		for i:=0;i<10;{
			peer=<-launcher.AddPeerChan
			models.IP = strings.Split(peer,":")[0]
			if(strings.Split(peer,":")[0]==local){
				beego.Trace("same ip, rejected,%v",i)
				continue
			}


			models.IP = strings.Split(peer,":")[0]
			beego.Trace("get Peer:",models.IP)
			i++
		}
	}()
}


func OpenPage(){
	cmd := exec.Command("open","http://localhost:8081")
	cmd.Run()
}
func GetAddr() string { //Get ip
	conn, err := net.Dial("udp", "baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return "Erorr"
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
