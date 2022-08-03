package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"os/exec"
	"time"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(30).Minutes().Do(func() {
		println("Triggering the job " + time.Now().Format(time.RFC3339Nano))
		cmd := exec.Command("halcommand.sh")
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(stdout))
		println("Finished the job " + time.Now().Format(time.RFC3339Nano))
		println("****************************************************")
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s.StartBlocking()
}
