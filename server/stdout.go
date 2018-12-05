package server

import (
	"fmt"
	"io"
	"log"
	"strings"
)

//等待服务器加载完地图
func (svr *Server) WaitEndLoading() {
	var buffer []byte = make([]byte, 4096)
	var retStr string
	//运行子进程
	svr.run_process()
	fmt.Println("正在加载服务器地图...")
	for {
		n, err := svr.Stdout.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Fatalln("子进程标准输出出错")
			} else {
				log.Fatalln("获取标准输出出错")
			}
			break
		}
		retStr = string(buffer[:n])
		if strings.Contains(retStr, "[Server thread/INFO]: Done") {
			fmt.Println("服务器地图加载完成!")
			break
		}
	}
}

//正式运行MCD
func (svr *Server) Run() {
	var buffer []byte = make([]byte, 4096)
	var retStr string
	// cl := command.GetInstance()
	for {
		n, err := svr.Stdout.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Fatalln("子进程标准输出出错")
			} else {
				log.Fatalln("获取标准输出出错")
			}
			break
		}
		retStr = string(buffer[:n])
		fmt.Println(retStr)
		// for _, val := range ParseMachineList {
		// 	_command, ok := val.Parsing(retStr)
		// 	//如果是命令,加入待执行列表
		// 	if ok {
		// 		if isFall := cl.Push(_command); isFall {
		// 			//待处理命令列表已满
		// 		}
		// 	}
		// }
	}
}
