/*
HTTP 下载
*/

package download

import (
	"context"
	"fmt"
	"github.com/virusdefender/goutils/errors"
	"io"
	"net/http"
)

func HTTPDownload(ctx context.Context, url string, toFile io.Writer) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return errors.Wrapf(err, "new request, url: %s", url)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "http.Do, url: %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("http status %d, url: %s", resp.StatusCode, url)
	}
	n, err := io.Copy(toFile, resp.Body)
	if err != nil {
		contentLength := resp.Header.Get("Content-Length")
		if contentLength == "" {
			contentLength = "<empty>"
		}
		return errors.Wrapf(err, "io.Copy failed, content length: %s, copied: %d, url: %s", contentLength, n, url)
	}
	return nil
}
