# GSI Tiles Downloader

[地理院タイル](https://maps.gsi.go.jp/development/ichiran.html)をダウンロードするためのツール

## 使い方

### ダウンロード

リリースまでしばらくお待ちください。

### コマンドライン

ディレクトリにダウンロード
```sh
gsi-tiles-downloader std --zoom 5-16 --output-dir ./tiles
```

MBTilesファイルにダウンロード
```sh
gsi-tiles-downloader std --zoom 5-16 --output-file ./tiles.mbtiles
```

PMTilesファイルにダウンロード
```sh
gsi-tiles-downloader std --zoom 5-16 --output-file ./tiles.pmtiles
```

ここでの`std`は標準地図などの地図種別を指定します。地図種別は[地理院タイル](https://maps.gsi.go.jp/development/ichiran.html)のページを参照してください。
以下に一部の地図種別を示します。

- `std`: 標準地図
- `pale`: 淡色地図
- `blank`: 白地図
- `english`: 英語地図
- `relief`: 色別標高図
- `experimental_bvmap`: 地理院地図Vector (試験的)
