package main

import (
	"bufio"
	"errors"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ayuxsys/ddg/log"
	"github.com/ayuxsys/ddg/must"
	"github.com/ayuxsys/ddg/scraper"
)

func main() {
	cmdi := NewCmdi()
	log.Setup(cmdi.verbose)
	scraper.SetupLogger(cmdi.verbose)
	cmdi.Validate()
	cmdi.Execute()
}

func (c *Cmdi) Execute() {
	scrpr := &scraper.Scraper{
		Opts: scraper.Options{
			OutputMode: c.mapOutputMode(),
		},
	}

	if c.out.file != "" {
		defer scrpr.SetupLogTarget(c.out.file)()
	}

	ctx, cancel := scrpr.NewChromeDPContext(!c.hideBrowser)
	defer cancel()
	if c.query != "" {
		_, err := scrpr.SearchDDGHeadless(c.query, ctx)
		if err != nil {
			log.Error.Fatalf("'scraper.SearchDDGHeadless': %v'", err)
		}
	}
	for _, q := range c.readQueryList() {
		_, err := scrpr.SearchDDGHeadless(q, ctx)
		if err != nil {
			if errors.Is(err, scraper.CaptchaErr) {
				log.Warn.Printf("failed to search %q. captcha error", q)
				continue
			}
			log.Error.Fatalf("'scraper.SearchDDGHeadless: %v'", err)
		}
		sleep := c.deferSleepTimeSec()
		log.Info.Printf("sleep %fs", sleep)
		time.Sleep(time.Duration(sleep) * time.Second)
	}
}

func (c *Cmdi) deferSleepTimeSec() float64 {
	if c.randomDelay == "" {
		return float64(c.delaySec)
	}
	itm := strings.Split(c.randomDelay, ",")
	min, max := must.Eval(strconv.Atoi(itm[0])), must.Eval(strconv.Atoi(itm[1]))
	return randFloat(min, max)
}

func randFloat(min, max int) float64 {
	return float64(min) + rand.Float64()*float64(max-min)
}

func (c *Cmdi) readQueryList() []string {
	var file *os.File
	var err error
	if c.queryList != "" {
		file, err = os.Open(c.queryList)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	} else {
		if !isPipedStdin() {
			return nil
		}
		file = os.Stdin
	}

	scanner := bufio.NewScanner(file)

	var out []string
	for scanner.Scan() {
		query := strings.TrimSpace(scanner.Text())
		out = append(out, query)
	}

	if err := scanner.Err(); err != nil {
		log.Error.Fatal("error reading organisation list file:", err)
	}
	return out
}

func isPipedStdin() bool {
	fi := must.Eval(os.Stdin.Stat())
	return (fi.Mode() & os.ModeCharDevice) == 0
}
