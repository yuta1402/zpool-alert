package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/yuta1402/zpool-alert/pkg/slack"
)

func main() {
	var (
		apiURL string
	)

	flag.StringVar(&apiURL, "api-url", "", "API of slack")

	flag.VisitAll(func(f *flag.Flag) {
		n := "ZPOOL_ALERT_" + strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
		if s := os.Getenv(n); s != "" {
			f.Value.Set(s)
		}
	})

	flag.Parse()

	out, err := exec.Command("zpool", "status", "-x").Output()
	if err != nil {
		log.Fatal(err)
	}

	result := string(out)
	if result == "all pools are healthy" {
		return
	}

	res, err := slack.PostAlert("```"+string(out)+"```", apiURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Status)
}
