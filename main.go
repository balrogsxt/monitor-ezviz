package main

import (
	"flag"
	"github.com/balrogsxt/monitor-ezviz/ezviz"
	"github.com/go-cmd/cmd"
	"log"
	"os"
	"strings"
)

var (
	key    = flag.String("key", "", "秘钥AppKey,留空则使用环境变量 EZVIZ_KEY")
	secret = flag.String("secret", "", "秘钥AppSecret,留空则使用环境变量 EZVIZ_SECRET")
	device = flag.String("device", "", "设备编号,留空则使用 EZVIZ_DEVICE")
	ffplay = flag.String("ffplay_path", "ffplay", "ffplay运行路径,默认调用环境变量ffplay,留空则使用环境变量 FFPLAY_PATH")
)

func main() {
	flag.Parse()
	ffplayPath := strings.Trim(*ffplay, " ")
	appKey := strings.Trim(*key, " ")
	appSecret := strings.Trim(*secret, " ")
	deviceId := strings.Trim(*device, " ")
	if len(appKey) == 0 {
		appKey = os.Getenv("EZVIZ_KEY")
	}
	if len(appSecret) == 0 {
		appSecret = os.Getenv("EZVIZ_SECRET")
	}
	if len(deviceId) == 0 {
		deviceId = os.Getenv("EZVIZ_DEVICE")
	}
	if len(ffplayPath) == 0 {
		ffplayPath = os.Getenv("FFPLAY_PATH")
	}

	client := ezviz.NewClient(appKey, appSecret)
	playAddr, err := client.GetPlayAddress(deviceId)
	if err != nil {
		log.Fatalf("获取播放地址失败: %s", err.Error())
	}
	log.Println("获取播放地址成功,准备播放")
	args := []string{
		"-window_title", "监控", playAddr,
	}
	//检测是否安装ffplay
	c := cmd.NewCmd(ffplayPath, args...)
	status := c.Start()
	for {
		select {
		case v := <-status:
			if v.Complete {
				break
			}
		case v := <-c.Stdout:
			log.Println(v)
		case v := <-c.Stderr:
			log.Println(v)
		}
		if c.Status().Complete {
			break
		}
		if c.Status().Error != nil {
			break
		}
	}
	log.Println("播放断开")
}
