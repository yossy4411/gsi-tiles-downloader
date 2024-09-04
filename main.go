package main

import (
	"fmt"
	"github.com/yossy4411/gsi-tiles-downloader/pkg/downloader"
)

func main() {
	fmt.Println("地理院地図タイルダウンローダー")
	fmt.Println("(c) 2024 おかゆグループ (Okayu Group)")
	fmt.Println("")
	fmt.Println("このプログラムは、地理院地図のタイルをダウンロードします。")
	fmt.Println("地理院地図のタイルのIDを指定してください。")
	fmt.Println("例: XYZタイルのURL「https://cyberjapandata.gsi.go.jp/xyz/std/{z}/{x}/{y}.png」の「std」がタイルのIDです。")
	fmt.Println("タイルのIDを入力してください。")
	fmt.Print("タイルのID: ")
	var tileID string
	_, err := fmt.Scan(&tileID)
	if err != nil {
		fmt.Println("入力エラーが発生しました。")
		return
	}
	fmt.Println("指定されたタイルのIDは「" + tileID + "」です。")
	client := downloader.NewDownloader(tileID, 10)
	fmt.Println("mokuroku.csvファイルを取得し、タイルのURLを取得します。")
}
