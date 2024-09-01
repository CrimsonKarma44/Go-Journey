package main

import (
	"github/setting"
	"github/urls"
)

//var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	setting.Init()
	urls.UrlHandler()
}
