package config

import "fmt"

type ABC struct {
	int_val int `하하하하하`
}

func (abc *ABC) print() {
	fmt.Println("%d", abc.int_val)
}

func Test() {
	fmt.Println("hahahah")
}
