// Description: Generic utils functions
// Author: Pixie79
// ============================================================================
// package utils

package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// UrlToLines retrieves the contents of a URL and returns them line by line.
//
// Parameters:
// - url: the URL to retrieve the contents from.
// - username: the username for basic authentication. If not needed, leave it empty.
// - password: the password for basic authentication. If not needed, leave it empty.
//
// Returns:
// - lines: an array of strings containing the lines of the retrieved content.
func UrlToLines(url string, username string, password string) []string {
	client := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	MaybeDie(err, "could not create http request")

	// Add basic auth if username and password are set
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	res, err := client.Do(req)
	MaybeDie(err, "could not authenticate")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		MaybeDie(err, "error closing connection")
	}(res.Body)

	if !InBetween(res.StatusCode, 200, 299) {
		Die(fmt.Sprintf("url access error %s, Status Code: %d", url, res.StatusCode))
	}

	return LinesFromReader(res.Body)
}
