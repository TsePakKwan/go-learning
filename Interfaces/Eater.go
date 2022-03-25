package main

import (
	"fmt"
	"strconv"
)

// 参考サイト
// https://qiita.com/rtok/items/46eadbf7b0b7a1b0eb08

// 食べるためのインターフェース
type Eater interface {
	PutIn()       // 口に入れる
	Chew()        // 噛む
	Swallow()     // 飲み込む
	Information() // 情報
}

// 人間の構造体
type Human struct {
	Height int // 身長
	Weight int // 体重
}

// カメの構造体
type Turtle struct {
	Kind string // 種類
}

// インターフェースが引数になる、食べるメソッド
func EatAll(e Eater) {
	e.PutIn() // インターフェースから呼び出す
	e.Chew()
	e.Swallow()
	e.Information()
}

// 人間用のインターフェースの実装
func (h Human) PutIn() {
	fmt.Println("道具を使って丁寧に口に運ぶ")
}
func (h Human) Chew() {
	fmt.Println("歯でしっかり噛む")
}
func (h Human) Swallow() {
	fmt.Println("よく噛んだら飲み込む")
}

func (h Human) Information() {
	fmt.Println("身長:" + strconv.Itoa(h.Height) + "cm,体重:" + strconv.Itoa(h.Weight) + "kg")
}

// カメ用のインターフェースの実装
func (h Turtle) PutIn() {
	fmt.Println("獲物を見つけたら首をすばやく伸ばして噛む")
}
func (h Turtle) Chew() {
	fmt.Println("クチバシで噛み砕く")
}
func (h Turtle) Swallow() {
	fmt.Println("小さく砕いたら飲み込む")
}
func (h Turtle) Information() {
	fmt.Println("種類:" + h.Kind)
}

// インターフェースが引数になる、飲むメソッド
func DrinkAll(e Eater) {
	e.PutIn() // インターフェースから呼び出す
	e.Swallow()
}

func main() {
	var eat Eater // インターフェースEater型の変数を用意

	var man Human = Human{Height: 300, Weight: 20} // 人間用の構造体を作成
	fmt.Println("＜人間が食べる＞")
	eat = man   // Human型がインターフェースであるEater型に変換される
	EatAll(eat) // インターフェースを引数に関数を呼ぶ

	var cheloniaMydas = Turtle{Kind: "アオウミガメ"} // カメ用の構造体を作成
	fmt.Println("＜カメが食べる＞")
	eat = cheloniaMydas // Turtle型がインターフェースであるEater型に変換される
	EatAll(eat)
}
