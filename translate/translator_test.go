/*
  @Date:   2021-05-25
  @File:   translator_test.go
  @Author: ls
*/
package translate

import (
	"fmt"
	"testing"
)

func TestTranslate(t *testing.T) {
	langTgt := language.Chinese.String()
	langSrc := language.English.String()
	text := "hello, world. this is google translate"

	translate, s, err := Translate(langTgt, langSrc, text, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(translate, s)
}
