package main

import (
	"github.com/ayuxsys/ddg/log"
	"github.com/ayuxsys/ddg/scraper"

	"github.com/spf13/pflag"
)

type Cmdi struct {
	out outputOpts
	//	timeout     int
	query       string
	queryList   string
	verbose     bool
	delaySec    int
	randomDelay string
	// headless    bool
	hideBrowser bool
}

type outputOpts struct {
	file     string
	jsonl    bool
	markdown bool
}

func NewCmdi() *Cmdi {
	cmdi := &Cmdi{}

	//pflag.IntVarP(&cmdi.timeout, "timeout", "t", 5, "timeout in seconds")
	pflag.StringVarP(&cmdi.query, "query", "q", "", "query to search")
	pflag.IntVarP(&cmdi.delaySec, "delay", "d", 5, "delay timeout in seconds to wait for while multiple scraping queries")
	pflag.StringVarP(&cmdi.queryList, "list", "l", "", "file containing a list of queries to search")
	pflag.StringVarP(&cmdi.randomDelay, "random", "r", "", "use random delay between requests in seconds (min,max)")
	pflag.BoolVarP(&cmdi.verbose, "verbose", "v", false, "enable verbose logging")
	//pflag.BoolVarP(&cmdi.headless, "headless", "b", false, "use an headless browser while scraping")
	pflag.BoolVarP(&cmdi.hideBrowser, "hide-browser", "i", false, "hide browser GUI while scraping (leads to captcha err better use 'xvfb')")
	pflag.StringVarP(&cmdi.out.file, "output", "o", "", "path to the output file (note: can only use one output mode at a time)")
	pflag.BoolVarP(&cmdi.out.jsonl, "json", "j", false, "enable jsonl output")
	pflag.BoolVarP(&cmdi.out.markdown, "md", "m", false, "enable markdown output")

	pflag.Usage = func() {
		println("Usage: ddg [flags]")
		pflag.PrintDefaults()
	}
	pflag.Parse()

	return cmdi
}

func (c *Cmdi) Validate() {
	if c.query == "" && c.queryList == "" && !isPipedStdin() {
		log.Error.Fatal("target query or query list is required")
	}
}

func (c *Cmdi) mapOutputMode() scraper.OutputMode {
	if c.out.jsonl {
		return scraper.JsonLOp
	}
	if c.out.markdown {
		return scraper.MarkDownOp
	}
	return scraper.MarkDownOp // default (todo: add pretty output)
}
