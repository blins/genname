package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const Alpabet = `qwertyupadfghkzxvbnm23456789`

func generate(count int, alpabet string) string {
	rand.Seed(time.Now().UnixNano())
	res := make([]byte, count)
	for i := 0; i < count; i++ {
		res[i] = alpabet[rand.Intn(len(alpabet))]
	}
	return string(res)
}

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		var count int = 10
		count_str := request.Form.Get("count")
		if count_str != "" {
			if c, err := strconv.Atoi(count_str); err == nil {
				count = c
			}
		}
		alpabet := Alpabet
		if a := request.Form.Get("alpabet"); a != "" {
			alpabet = a
		}
		writer.Write([]byte(generate(count, alpabet)))
	})
	http.ListenAndServe(":8080", nil)
}
