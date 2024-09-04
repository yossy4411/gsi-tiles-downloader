package downloader

import (
	"strconv"
	"strings"
)

type Tile struct {
	X    int    // タイルのX座標
	Y    int    // タイルのY座標
	Z    int    // ズームレベル
	URL  string // タイルのURL
	hash string // MD5ハッシュ
}

// NewTile はCSVレコードから新しいTileを作成します。
func NewTile(record []string) *Tile {
	path := record[0] // タイルのパス
	x, y, z := parseTilePath(path)
	hash := record[3] // MD5ハッシュ
	return &Tile{
		X:    x,
		Y:    y,
		Z:    z,
		URL:  path,
		hash: hash,
	}
}

// parseTilePath はタイルのパスからX座標、Y座標、ズームレベルを取得します。
func parseTilePath(path string) (int, int, int) {
	// タイルのパスの例: "15/5241/12661.png"
	// 拡張子を除いた文字列を取得します。
	path = path[:len(path)-4]
	// "/" で分割します。
	parts := strings.Split(path, "/")
	// パーツの数が3でない場合(エラー)はデフォルト値を返します。
	if len(parts) != 3 {
		return 0, 0, 0
	}
	// パーツを数値に変換します。
	x, _ := strconv.Atoi(parts[1])
	y, _ := strconv.Atoi(parts[2])
	z, _ := strconv.Atoi(parts[0])
	return x, y, z
}
