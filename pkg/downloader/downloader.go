package downloader

type Downloader struct {
	tileID     string   // タイルのID
	tileCount  int      // タイルの数
	tileURLs   []string // タイルのURL
	maxThreads int      // 最大スレッド数
}

// NewDownloader は新しいDownloaderを作成します。
func NewDownloader(tileID string, maxThreads int) *Downloader {
	return &Downloader{
		tileID:     tileID,
		tileCount:  0,
		tileURLs:   nil,
		maxThreads: maxThreads,
	}
}

// GetTileCount はタイルのIDを取得します。
func (d *Downloader) GetTileCount() int {
	return d.tileCount
}
