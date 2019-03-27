package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	domain := os.Args[1]
	// url_str := "http://localhost:8080/add_cookie/index.php"
	// values := url.Values{}
	// values.Set("domain", domain)
	//
	// resp, _ := http.Get(url_str + "?" + values.Encode())
	// defer resp.Body.Close()
	//
	// byteArray, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(byteArray)) // htmlをstringで取得

	cookie1 := &http.Cookie{Name: "sample", Value: "sample", Domain: "", HttpOnly: true}
	http.SetCookie(w, cookie1)
}
