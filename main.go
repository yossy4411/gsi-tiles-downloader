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

	fmt.Println("タイルのズームレベルの範囲を指定してください。")
	fmt.Print("最小ズームレベル [5]: ")
	var minZoomLevel, maxZoomLevel int

	_, err = fmt.Scan(&minZoomLevel)
	if err != nil {
		fmt.Println("入力エラーが発生しました。")
		return
	}

	fmt.Print("最大ズームレベル [16]: ")
	_, err = fmt.Scan(&maxZoomLevel)
	if err != nil {
		fmt.Println("入力エラーが発生しました。")
		return
	}

	fmt.Println("指定されたズームレベルの範囲は" + fmt.Sprint(minZoomLevel) + "から" + fmt.Sprint(maxZoomLevel) + "です。")

	client := downloader.NewDownloader(tileID, 10)
	err = client.SetTargetDir("tiles")
	if err != nil {
		fmt.Println("ターゲットディレクトリの作成に失敗しました。")
		fmt.Println(err)
		return
	}
	client.ZoomMin = minZoomLevel
	client.ZoomMax = maxZoomLevel
	fmt.Println("mokuroku.csvファイルを取得し、タイルのURLを取得します。")
	mokuroku := client.GetMokurokuURL()
	fmt.Println("mokuroku.csv: " + mokuroku)
	err = client.GetURLs(mokuroku)
	if err != nil {
		fmt.Println("mokuroku.csvファイルの取得に失敗しました。")
		fmt.Println(err)
		return
	}
	fmt.Println("指定されたズームレベルの範囲で見つかったタイルは" + fmt.Sprint(client.GetTileCount()) + "個です。")
	fmt.Println("また、タイルの範囲は" + fmt.Sprint(client.ZoomMin) + "から" + fmt.Sprint(client.ZoomMax) + "です。")
	fmt.Println("タイルのダウンロードを開始します。")

	for {
		err = client.DownloadTile()
		if err != nil {
			fmt.Println("タイルのダウンロードに失敗しました。")
			fmt.Println(err)
			return
		}
		if client.IsFinished() {
			break
		}
		// プログレスバーを表示します。
		fmt.Printf("\rダウンロード中... %d/%d", client.Downloaded, client.GetTileCount())
	}

}
