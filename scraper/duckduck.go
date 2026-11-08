package scraper

import (
	"context"
	"fmt"
	"strings"

	"github.com/ayuxsys/ddg/must"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func (s *Scraper) SearchDDGHeadless(query string, ctx context.Context) (Output, error) {
	if query == "" {
		return Output{}, nil
	}

	out := Output{Query: query}

	if err := chromedp.Do(ctx, chromedp.Navigate("https://duckduckgo.com/?q="+query)); err != nil {
		return out, fmt.Errorf("'chromedp.Do': %w", err)
	}

	html, err := chromedp.Run(ctx, chromedp.Evaluate[string](`document.body.innerHTML`))
	if err != nil {
		return out, fmt.Errorf("get innerHTML of the body 'chromedp.Run': %w", err)
	}
	debugLogRes(html)

	if err := validateDDGResponse(html); err != nil {
		return out, fmt.Errorf("'validateDDGResponse': %w", err)
	}

	out.Results = parseDDGHeadless(html)
	s.logOutput(out)

	return out, nil
}

func parseDDGHeadless(html string) []Result {
	doc := must.Eval(
		goquery.NewDocumentFromReader(strings.NewReader(html)),
	)

	var results []Result
	doc.Find(`a[data-testid="result-title-a"]`).Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text())
		link, _ := s.Attr("href")
		results = append(results, Result{Title: title, Href: link})
	})

	return results
}

func validateDDGResponse(resp string) error {
	// doesn't contains the query in response so not reliable ig
	// if strings.Contains(resp, DDGNoResultsPrefix+query) {
	// 	return errors.New(DDGNoResultsPrefix + query)
	// }
	if strings.Contains(resp, "Unfortunately, bots use DuckDuckGo too.") {
		return CaptchaErr
	}
	return nil
}

// todo: maybe useable with rotating proxy but meh!
// func (s *Scraper) searchDDG(query string, reqTimeoutSec time.Duration) (Output, error) {
// 	out := Output{Query: query}
// 	data := url.Values{}
// 	data.Set("q", query)
// 	log.Debug.Printf("searching query %q", query)
//
// 	req := must.Eval(
// 		http.NewRequest(
// 			http.MethodPost,
// 			"https://html.duckduckgo.com/html/",
// 			strings.NewReader(data.Encode()),
// 		),
// 	)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:157.0) Gecko/20100101 Firefox/157.0")
//
// 	client := http.DefaultClient
// 	client.Timeout = reqTimeoutSec
//
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return out, fmt.Errorf("send request to duckduckgo API: %v", err)
// 	}
// 	defer resp.Body.Close()
//
// 	respBody := string(must.Eval(io.ReadAll(resp.Body)))
//
// 	if err := validateDDGResponse(query, respBody); err != nil {
// 		return out, err
// 	}
//
// 	out.Results = parseDDG(respBody)
//
// 	return out, nil
// }
//
//
// func parseDDG(resp string) []Result {
// 	doc := must.Eval(
// 		goquery.NewDocumentFromReader(strings.NewReader(resp)),
// 	)
//
// 	var results []Result
// 	doc.Find("h2.result__title a.result__a").Each(func(i int, s *goquery.Selection) {
// 		title := strings.TrimSpace(s.Text())
// 		link, _ := s.Attr("href")
// 		results = append(results, Result{Title: title, Href: link})
// 	})
//
// 	return results
// }
