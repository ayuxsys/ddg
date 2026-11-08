package scraper

import (
	"context"

	"github.com/chromedp/chromedp"
)

type Scraper struct {
	Opts Options
}

type Options struct {
	OutputMode OutputMode
}

type HeadlessOpts struct {
	ShowBrowser bool
	Ctx         context.Context
}

type OutputMode int

const (
	JsonLOp OutputMode = iota + 1
	MarkDownOp
)

type Output struct {
	Query   string   `json:"query"`
	Results []Result `json:"results"`
}

type Result struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}

func (s *Scraper) Validate() error {
	return nil
}

func (s *Scraper) NewChromeDPContext(showBrowser bool) (context.Context, context.CancelFunc) {
	var opts []chromedp.ContextOption

	if showBrowser {
		opts = append(opts, chromedp.WithVisibleWindow())
	}

	return chromedp.NewContext(
		context.Background(),
		opts...,
	)
}
