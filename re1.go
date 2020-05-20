package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var (
	rePhone = `1[3456789]\d{9}`
)

func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}
func main() {
	resp, err := http.Get("http://ilhw.cn/?bd_vid=11111897003490782706")
	HandleError(err, `http.Get("http://ilhw.cn/?bd_vid=11111897003490782706")`)
	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	re := regexp.MustCompile(rePhone)
	sult := re.FindAllString(html, -1)
	for _, num := range sult {
		fmt.Println(num)

	}

}
