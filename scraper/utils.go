package scraper

import (
	"errors"
	"fmt"
	"net/http"
)

func validateResp(resp *http.Response) error {
	if resp == nil {
		return errors.New("response is nil")
	}

	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		return errors.New("rate limit error: 429 Too Many Requests")

	case http.StatusUnauthorized:
		return errors.New("authentication error: 401 Unauthorized")

	case http.StatusForbidden:
		return errors.New("authorization error: 403 Forbidden")

	case http.StatusNotFound:
		return errors.New("resource not found: 404 Not Found")

	case http.StatusBadRequest:
		return errors.New("bad request: 400 Bad Request")

	default:
		if resp.StatusCode >= http.StatusInternalServerError {
			return fmt.Errorf("server error: HTTP %d", resp.StatusCode)
		}

		if resp.StatusCode >= http.StatusBadRequest {
			return fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
		}
	}

	return nil
}
