package config

import (
	"app/cron"
	robfig "github.com/robfig/cron"
)

func SetupCron() {
	c := robfig.New()

	// 添加定时任务
	c.AddFunc("* * * * *", cron.Task1)
	c.AddFunc("0 0 12 * * *", cron.Task2)

	// 启动定时任务
	c.Start()
}
