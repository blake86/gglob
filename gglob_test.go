/*
@Time : 2020-04-15 11:41
@Author : lihao
@File : Range_test
@Software: GoLand
*/
package gglob

import (
	"fmt"
	"testing"
)

func TestExpand(t *testing.T) {
	fmt.Println(Expand("rsync.master[001-5].hadoop"))
}

func TestT(t *testing.T) {
	id := 0
	fmt.Printf("%000*d\n", 0, id)
	fmt.Printf("%2d\n", id)
}

func TestParse2(t *testing.T) {
	str := `prefix[1-3]suffix`
	fmt.Println(Expand(str))
	str = `prefix[1-3,5]suffix`
	fmt.Println(Expand(str))
	str = `prefix[1-3]mid[004-7]suffix`
	ls, err := Expand(str)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range ls {
		fmt.Println(s)
	}
	fmt.Println(Expand(str))
	str = "prefix[1,3-10,12]mid[01-4,3-5]suffix"
	fmt.Println(str)

	ls, err = Expand(str)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range ls {
		fmt.Println(s)
	}

}
