// map vs struct
// リファレンス型かタイプ型か -> mapはポインターを考える必要がない
// structはいろいろなタイプを持てるがmapは固定のタイプ
// mapはindexがあるのでイテレート可能

package main

import "fmt"

func main() {

	// 基本の宣言 key: string, value: string
	// var colors map[string]string

	// makeを使った宣言
	// colors := make(map[string]string)

	// 値を入れつつ宣言
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#f4bf745",
		"white": "#ffffff",
	}

	// 値を入れる
	// colors["white"] = "#ffffff"

	// 値を削除
	// delete(colors, "white")

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
