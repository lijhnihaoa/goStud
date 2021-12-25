package example

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func StudGetUrl() {
	for _, url := range os.Args[1:] {
		if url[:7] != "http://" {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v\n", err)
			os.Exit(1)
		}

		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// var b io.Writer
		fmt.Printf("status code: %s\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %s : %v\n", url, err)
			os.Exit(1)
		}
		// fmt.Printf("%s", b)
	}
}
