package httpControllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

var (
	sourceUrl = "https://httpbin.org"

	username = "abed"
)

func POSTrequest(clip string) {

	data := url.Values{
		"username":   {username},
		"occupation": {clip},
	}

	resp, err := http.PostForm(sourceUrl+"/post", data)

	if err != nil {
		log.Fatal(err)
	}
	//TODO: Need to work on printing better messages.
	if resp.StatusCode == 200 {
		fmt.Println("Clip Successfully Posted")
	} else {
		log.Fatal(resp.StatusCode)
	}
}

func GETrequest() string {
	//TODO: Modify print message according to github.com/JammUtkarsh/cshare-server GET method
	resp, err := http.Get(sourceUrl + "/get")

	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	fmt.Println(resp)

	return ""
}
