package config

import "fmt"

type ABC struct {
	Int_val int `하하하하하`
}

func (abc *ABC) Print() {
	fmt.Println("%d", abc.Int_val)
}

//Test function do first character of function name to capital letter to use external
func Test() {
	fmt.Println("hahahah")
}
