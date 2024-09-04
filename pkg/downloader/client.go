package downloader

import (
	"io"
	"net/http"
)

type client struct {
	finished   bool         // ダウンロードが完了したかどうか
	httpClient *http.Client // HTTPクライアント
}

// newClient は新しいClientを作成します。
func newClient() *client {
	return &client{
		finished:   true,
		httpClient: &http.Client{},
	}
}

// downloadTile はタイルをダウンロードします。
func (c *client) downloadTile(url string) ([]byte, error) {
	c.finished = false
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		c.finished = true
		err := resp.Body.Close()
		if err != nil {
			panic("failed to close response body")
		}
	}()
	// レスポンスをバイト列に変換します。
	return io.ReadAll(resp.Body)

}
