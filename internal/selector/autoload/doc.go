package autoload

// v4のブロック毎の各ParserではregistryパッケージにSelector情報が登録しており
// それらを一括でimportするためのパッケージです。
// 下記の run go generate をする事で
// 各パッケージ内に起動時にSelectorの情報を登録するための register_gen.go と
// 一括インポート用の　packages_gen.go　が自動生成されます。

//go:generate go run ../../cmd/gencode
