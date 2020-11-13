package main

import "fmt"

func main() {

	// 基本の宣言
	// alex := person{"Alex", "Anderson"}
	// alex2 := person{firstName: "Alex", lastName: "Anderson"}

	// フィールドの更新
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// フィールドつき出力
	// fmt.Printf("%+v\n", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimという値のメモリ上のアドレスが入る
	// jimPointer := &jim
	// jimだけで、レシーバーがわで*を使うことでショートカットできる
	jim.updateName("jimmy")
	jim.print()

}

// receiverの*person -> これはタイプに対する説明のようなもの personに対してのポインターを設定
func (pointerToPerson *person) updateName(newFirstName string) {
	// こちらはオペレータ *pointerでメモリーアドレスが指し示す値を取得する
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
