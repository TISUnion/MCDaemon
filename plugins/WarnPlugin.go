/*
 * 转发服务器警告消息，并在接收到fatal时自动重启服务器进程
 * author: Sciroccogti
 */

package plugin

import (
	"MCDaemon-go/command"
	"MCDaemon-go/lib"
	"fmt"
	"strconv"
	//"time"
)

type WarnPlugin struct{}

func (wp *WarnPlugin) Handle(c *command.Command, s lib.Server) {
	if len(c.Argv) != 3 {
		c.Argv[0] = "help"
	}

	switch c.Argv[0] {
	case "warn":
		ticks, _ := strconv.Atoi(c.Argv[2])
		if ticks >= 40 && ticks < 60 {
			s.Say(command.Text{fmt.Sprintf("嗯？服务姬有点忙不过来了，延迟%dticks~", ticks), "gray"})
			s.WriteLog("info", fmt.Sprintf("服务器延迟%dticks", ticks))
		} else if ticks >= 60 && ticks < 80 {
			s.Say(command.Text{fmt.Sprintf("哎呀呀，让服务姬歇一会吧，延迟%dticks！", ticks), "yellow"})
			s.WriteLog("warn", fmt.Sprintf("服务器延迟%dticks", ticks))
		} else if ticks >= 80 && ticks < 100 {
			s.Say(command.Text{fmt.Sprintf("呜呜呜，服务姬受不了了！延迟%dticks！", ticks), "red"})
			s.WriteLog("warn", fmt.Sprintf("服务器延迟%dticks", ticks))
		} else if ticks >= 100 {
			s.Say(command.Text{"服务器负载过高！请立即停止活动并联系服主！", "red"})
			s.WriteLog("fatal", fmt.Sprintf("服务器延迟%dticks", ticks))
			//time.Sleep(time.Second * 10)
			//s.Restart()
		}
	default:
	}
}

func (wp *WarnPlugin) Init(s lib.Server) {
}

func (wp *WarnPlugin) Close() {
}
