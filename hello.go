package hello_webview

import (
	"fmt"
	"io"
	"net/http"
)

var Name string = "World"

func init() {
	go func() {
		http.HandleFunc("/", hello)
		http.ListenAndServe(":8080", nil)
	}()
}

func SetName(name string) {
	Name = name
}

func Greetings() string {
	return fmt.Sprintf("Hello, %s!", Name)
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
	<head>
		<title>Hello World</title>
	</head>
	<body>
		` + Greetings() + `
	</body>
</html>`,
	)
}


