package _3_Facade

import (
	"log"
)

func TestFacade() {
	token := "token-valido"
	user := "user-blog"
	to := "a@algo.com"
	comment := "muy buen video :)"

	facade := New(to, comment, token, user)
	err := facade.Comment()
	if err != nil {
		log.Fatal(err)
	}

	facade.Notify()
}
