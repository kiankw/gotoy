package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func cron1() {
	fmt.Println("cron1")
	time.Sleep(3 * time.Second)
}

func cron2() {
	fmt.Println("cron2")
	time.Sleep(3 * time.Second)
}

func main() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(timezone)

	// 每秒执行一次
	s.Every(1).Seconds().Do(func() {
		cron1()
	})

	// 每秒执行一次
	s.Every(1).Second().Do(func() {
		cron2()
	})

	s.StartBlocking()
}

func main1() {

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	defer fmt.Println("happy end")

	for range ticker.C {
		fmt.Println("test")
	}
}
