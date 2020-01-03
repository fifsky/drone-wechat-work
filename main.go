package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/fifsky/drone-wechat-work/wechat"
)

func main() {
	app := cli.NewApp()
	app.Name = "WeChat work robot plugin"
	app.Usage = "Wechat Work robot plugin"
	app.Action = run
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url",
			Usage: "The wechat work robot url",
		},
		cli.StringFlag{
			Name:  "msgtype",
			Usage: "The type of message, either text, textcard",
			Value: "textcard",
		},
		cli.StringFlag{
			Name:  "touser",
			Usage: "The users to send the message to, @all for all users",
			Value: "@all",
		},
		cli.StringFlag{
			Name:  "content",
			Usage: "message content",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	robot := wechat.WeChat{
		Url:     c.String("url"),
		MsgType: c.String("msgtype"),
		ToUser:  c.String("touser"),
		Content: c.String("content"),
	}
	return robot.Send()
}
