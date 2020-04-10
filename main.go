package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/goapt/dotenv"
	"github.com/urfave/cli"

	"github.com/fifsky/drone-wechat-work/wechat"
)

func main() {
	log.Println("Start notify")

	if _, err := os.Stat("/run/drone/env"); err == nil {
		_ = dotenv.Overload("/run/drone/env")
		str, _ := ioutil.ReadFile("/run/drone/env")
		log.Println(string(str))
	}

	app := cli.NewApp()
	app.Name = "WeChat work robot plugin"
	app.Usage = "Wechat Work robot plugin"
	app.Action = run
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "url",
			Usage:  "The wechat work robot url",
			EnvVar: "PLUGIN_URL",
		},
		cli.StringFlag{
			Name:   "msgtype",
			Usage:  "The type of message, either text, markdown",
			Value:  "text",
			EnvVar: "PLUGIN_MSGTYPE",
		},
		cli.StringFlag{
			Name:   "touser",
			Usage:  "The users to send the message to, @all for all users",
			Value:  "@all",
			EnvVar: "PLUGIN_TOUSER",
		},
		cli.StringFlag{
			Name:   "content",
			Usage:  "message content",
			EnvVar: "PLUGIN_CONTENT",
		},

		// template
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	robot := wechat.WeChat{
		Build: wechat.Build{
			Owner:   c.String("repo.owner"),
			Name:    c.String("repo.name"),
			Tag:     c.String("build.tag"),
			Number:  c.Int("build.number"),
			Event:   c.String("build.event"),
			Status:  c.String("build.status"),
			Commit:  c.String("commit.sha"),
			Ref:     c.String("commit.ref"),
			Branch:  c.String("commit.branch"),
			Author:  c.String("commit.author"),
			Message: c.String("commit.message"),
			Link:    c.String("build.link"),
			Started: c.Int64("build.started"),
			Created: c.Int64("build.created"),
		},
		Url:     c.String("url"),
		MsgType: c.String("msgtype"),
		ToUser:  c.String("touser"),
		Content: c.String("content"),
	}

	err := robot.Send()

	if err != nil {
		log.Println("notify fail", err)
	} else {
		log.Println("notify success, DRONE_BUILD_STATUS:", os.Getenv("DRONE_BUILD_STATUS"))
	}

	return err
}
