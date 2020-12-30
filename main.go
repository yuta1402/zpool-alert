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

	ilog := log.New(os.Stdout, "[info] ", log.LstdFlags|log.LUTC)
	elog := log.New(os.Stderr, "[error] ", log.LstdFlags|log.LUTC)

	out, err := exec.Command("zpool", "status", "-x").Output()
	if err != nil {
		elog.Fatal(err)
	}

	result := string(out)

	if strings.Contains(result, "all pools are healthy") {
		ilog.Println(result)
		return
	}

	res, err := slack.PostAlert("```"+string(out)+"```", apiURL)
	if err != nil {
		elog.Fatal(err)
	}

	ilog.Println(res.Status)
}
