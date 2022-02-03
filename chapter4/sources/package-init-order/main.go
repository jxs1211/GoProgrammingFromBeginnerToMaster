package main

import (
	"database/sql"
	"fmt"
	"log"
	"unicode/utf8"

	_ "github.com/bigwhite/package-init-order/pkg1"
	_ "github.com/bigwhite/package-init-order/pkg3"
	pq "github.com/lib/pq"
)

var (
	d  = pq.Driver{}
	_  = constInitCheck()
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

const (
	c1 = "c1"
	c2 = "c2"
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("main: const c1 init")
	}
	if c1 != "" {
		fmt.Println("main: const c2 init")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("main: var %s init\n", name)
	return name
}

func init() {
	fmt.Println("main: init")
}

var specialBytes [16]byte

// special reports whether byte b needs to be escaped by QuoteMeta.
func special(b byte) bool {
	return b < utf8.RuneSelf && specialBytes[b%16]&(1<<(b/16)) != 0
}

func init2() {
	for _, b := range []byte(`\.+*?()|[]{}^$`) {
		fmt.Printf("%v|%b|%s|%x|%U", b, b, string(b), b, b)
		break
		specialBytes[b%16] |= 1 << (b / 16)
	}
}

func ShowInitUsageInRegexPkg() {
	fmt.Println(specialBytes)
	init2()
	fmt.Println(specialBytes)
	b := 36
	fmt.Println(specialBytes[b%16] & (1 << (b / 16)))

}

func ShowInitRegisterMode() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("err: ", err)
	}
	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	fmt.Println("rows: ", rows)

}

func main() {
	// do nothing
	ShowInitRegisterMode()
}
