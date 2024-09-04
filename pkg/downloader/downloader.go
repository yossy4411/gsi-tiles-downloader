package downloader

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Downloader struct {
	tileID     string    // タイルのID
	tileCount  int       // タイルの数
	tileURLs   []*Tile   // タイルのURL
	maxThreads int       // 最大スレッド数
	clients    []*client // クライアント
	ZoomMin    int       // 最小ズームレベル
	ZoomMax    int       // 最大ズームレベル
	baseURL    string    // ベースURL
	Downloaded int       // ダウンロード済みのタイル数
	targetDir  string    // ターゲットディレクトリ
}

// NewDownloader は新しいDownloaderを作成します。
func NewDownloader(tileID string, maxThreads int) *Downloader {
	clients := make([]*client, maxThreads)
	for i := 0; i < maxThreads; i++ {
		clients[i] = newClient()
	}

	return &Downloader{
		tileID:     tileID,
		tileCount:  0,
		tileURLs:   nil,
		maxThreads: maxThreads,
		clients:    clients,
		ZoomMin:    5,
		ZoomMax:    17,
		baseURL:    "https://cyberjapandata.gsi.go.jp/xyz/" + tileID + "/",
		Downloaded: 0,
	}
}

// ensureDirExists checks if a directory exists, and creates it if it doesn't.
func ensureDirExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetTargetDir はターゲットディレクトリを設定します。
func (d *Downloader) SetTargetDir(dir string) error {
	d.targetDir = dir
	// ディレクトリが存在しない場合は作成します。
	return ensureDirExists(dir)
}

// GetTileCount はタイルのIDを取得します。
func (d *Downloader) GetTileCount() int {
	return d.tileCount
}
func (d *Downloader) GetMokurokuURL() string {
	return "https://cyberjapandata.gsi.go.jp/xyz/" + d.tileID + "/mokuroku.csv.gz"
}

// GetURLs はタイルのmokuroku.csvファイルを取得し、タイルのURLを取得します。
func (d *Downloader) GetURLs(path string) error {
	resp, err := downloadFile(path)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("failed to close response body")
		}
	}(resp.Body)
	gzReader, err := parseGzip(resp)
	if err != nil {
		return err
	}
	defer func(gzReader *gzip.Reader) {
		err := gzReader.Close()
		if err != nil {
			fmt.Println("failed to close gzip reader")
		}
	}(gzReader)
	reader := csv.NewReader(gzReader)
	var minZoom = d.ZoomMax
	var maxZoom = d.ZoomMin
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		tile := NewTile(record)
		if tile.Z < d.ZoomMin || tile.Z > d.ZoomMax {
			continue
		}
		d.tileURLs = append(d.tileURLs, tile)

		if tile.Z < minZoom {
			minZoom = tile.Z
		}
		if tile.Z > maxZoom {
			maxZoom = tile.Z
		}
	}

	d.ZoomMax = maxZoom
	d.ZoomMin = minZoom
	d.tileCount = len(d.tileURLs)

	return nil
}

func downloadFile(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file: %s", resp.Status)
	}
	return resp, nil
}

func parseGzip(resp *http.Response) (*gzip.Reader, error) {

	gzReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return gzReader, nil
}

// DownloadTile はタイルをダウンロードします。
func (d *Downloader) DownloadTile() error {
	var index = 0
	for {
		index = (index + 1) % d.maxThreads
		if d.clients[index].finished {
			err := d.saveToFile(d.Downloaded, index)
			if err != nil {
				return err
			}
			d.Downloaded++
			break
		}
	}
	return nil
}

func (d *Downloader) saveToFile(i int, index int) error {
	data, err := d.clients[index].downloadTile(d.baseURL + d.tileURLs[i].URL)
	if err != nil {
		return err
	}
	url := d.targetDir + "/" + d.tileURLs[i].URL
	dirs := strings.Split(url, "/")
	err = ensureDirExists(strings.Join(dirs[:len(dirs)-1], "/"))
	if err != nil {
		return err
	}
	file, err := os.Create(url)

	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("failed to close file")
		}
	}(file)
	_, err = file.Write(data)
	return err
}

// IsFinished はダウンロードが完了したかどうかを返します。
func (d *Downloader) IsFinished() bool {
	return d.Downloaded == d.tileCount
}
