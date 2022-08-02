package internal

import (
	"fmt"
	"github.com/Ravior/gserver/database/gdb"
	"github.com/Ravior/gserver/os/glog"
	"github.com/Ravior/gserver/util/gconfig"
	"os"
	"runtime"
)

var (
	logo = `                                        
  ____ ____                          
 / ___/ ___|  ___ _ ____   _____ _ __ 
| |  _\___ \ / _ \ '__\ \ / / _ \ '__|
| |_| |___) |  __/ |   \ V /  __/ |   
 \____|____/ \___|_|    \_/ \___|_|     
                                        `
	topLine    = `┌───────────────────────────────────────────────────┐`
	borderLine = `│`
	bottomLine = `└───────────────────────────────────────────────────┘`
)

// InitBaseModule 基础模块
func InitBaseModule(serverIds ...string) {
	fmt.Println("[1]初始化日志模块")
	if len(serverIds) > 0 {
		glog.Init(serverIds[0])
	} else {
		glog.Init("")
	}

	fmt.Println("[2]初始化数据库配置模块")
	gdb.Init()
}

// PrintSystemInfo 打印系统信息
func PrintSystemInfo() {
	fmt.Println(logo)
	fmt.Println(topLine)
	fmt.Printf("%s Author           %s %s\r\n", borderLine, borderLine, "https://gitlib.com")
	fmt.Printf("%s Version          %s %s\r\n", borderLine, borderLine, gconfig.Global.Version)
	fmt.Printf("%s ServerName       %s %s\r\n", borderLine, borderLine, gconfig.Global.ServerName)
	fmt.Printf("%s ServerId         %s %s\r\n", borderLine, borderLine, gconfig.Global.ServerId)
	fmt.Printf("%s Pid              %s %d\r\n", borderLine, borderLine, os.Getpid())
	fmt.Printf("%s GOMAXPROCS       %s %d\r\n", borderLine, borderLine, runtime.GOMAXPROCS(-1))
	fmt.Printf("%s NumCPU           %s %d\r\n", borderLine, borderLine, runtime.NumCPU())
	// 如果配置了TCP端口则打印TCP服务器配置信息
	if gconfig.Global.TcpServer.Port > 0 {
		fmt.Println(bottomLine)
		fmt.Println(topLine)
		fmt.Printf("%s Protocol         %s %s\r\n", borderLine, borderLine, "TCP")
		fmt.Printf("%s IP               %s %s\r\n", borderLine, borderLine, gconfig.Global.TcpServer.IP)
		fmt.Printf("%s Port             %s %d\r\n", borderLine, borderLine, gconfig.Global.TcpServer.Port)

		if gconfig.Global.TcpServer.WorkerPoolSize > 0 {
			fmt.Printf("%s WorkerPollSize   %s %d\r\n", borderLine, borderLine, gconfig.Global.TcpServer.WorkerPoolSize)
		}
		if gconfig.Global.TcpServer.WorkerTaskLen > 0 {
			fmt.Printf("%s WorkerSize       %s %d\r\n", borderLine, borderLine, gconfig.Global.TcpServer.WorkerTaskLen)
		}
		if gconfig.Global.TcpServer.MaxConn > 0 {
			fmt.Printf("%s MaxConn          %s %d\r\n", borderLine, borderLine, gconfig.Global.TcpServer.MaxConn)
		}
	}

	//  如果配置了TCP端口则打印TCP服务器配置信息
	if gconfig.Global.WsServer.Port > 0 {
		fmt.Println(bottomLine)
		fmt.Println(topLine)
		fmt.Printf("%s Protocol         %s %s\r\n", borderLine, borderLine, "WebSocket")
		fmt.Printf("%s IP               %s %s\r\n", borderLine, borderLine, gconfig.Global.WsServer.IP)
		fmt.Printf("%s Port             %s %d\r\n", borderLine, borderLine, gconfig.Global.WsServer.Port)

		if gconfig.Global.WsServer.WorkerPoolSize > 0 {
			fmt.Printf("%s WorkerPollSize   %s %d\r\n", borderLine, borderLine, gconfig.Global.WsServer.WorkerPoolSize)
		}
		if gconfig.Global.WsServer.WorkerTaskLen > 0 {
			fmt.Printf("%s WorkerSize       %s %d\r\n", borderLine, borderLine, gconfig.Global.WsServer.WorkerTaskLen)
		}
		if gconfig.Global.WsServer.MaxConn > 0 {
			fmt.Printf("%s MaxConn          %s %d\r\n", borderLine, borderLine, gconfig.Global.WsServer.MaxConn)
		}
	}
	// 如果配置了Http端口则打印Http服务器配置
	if gconfig.Global.HttpServer.Port > 0 {
		fmt.Println(bottomLine)
		fmt.Println(topLine)
		fmt.Printf("%s Protocol         %s %s\r\n", borderLine, borderLine, "HTTP")
		fmt.Printf("%s HttpIp           %s %s\r\n", borderLine, borderLine, gconfig.Global.HttpServer.IP)
		fmt.Printf("%s HttpPort         %s %d\r\n", borderLine, borderLine, gconfig.Global.HttpServer.Port)
	}

	if len(gconfig.Global.DataBase) > 0 {
		fmt.Println(bottomLine)
		fmt.Println(topLine)
		for pool, database := range gconfig.Global.DataBase {
			fmt.Printf("%s MySQL            %s %s\r\n", borderLine, borderLine, pool)
			fmt.Printf("%s MySQL IP         %s %s\r\n", borderLine, borderLine, database.Host)
			fmt.Printf("%s MySQL Port       %s %d\r\n", borderLine, borderLine, database.Port)
			fmt.Printf("%s MySQL DbName     %s %s\r\n", borderLine, borderLine, database.DbName)
			fmt.Printf("%s MySQL Prefix     %s %s\r\n", borderLine, borderLine, database.Prefix)
			fmt.Printf("%s MySQL ShowLog    %s %v\r\n", borderLine, borderLine, database.ShowLog)
		}
	}

	fmt.Println(bottomLine)
}
