package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	chanImageUrls chan string
	waitGroup     sync.WaitGroup

	chanTask chan string
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func handleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func downloadFile(url string, filename string) (ok bool) {
	fmt.Println(url, filename)
	return true
	resp, err := http.Get(url)
	handleError(err, "Failed to get url")
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	handleError(err, "Failed to read response body")

	filename = "/Users/edward/workspace/go-gin/scraper-result/" + filename

	err = os.WriteFile(filename, bytes, 0666)

	if err != nil {
		return false
	} else {
		return true
	}
}

func getFilenameFromUrl(url string) (filename string) {
	lastIndex := strings.LastIndex(url, "/")
	filename = url[lastIndex+1:]
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "-" + filename

	return
}

func downloadImg() {
	for url := range chanImageUrls {
		filename := getFilenameFromUrl(url)
		ok := downloadFile(url, filename)
		if ok {
			fmt.Printf("%s download success!\n", filename)
		} else {
			fmt.Printf("%s download failed!\n", filename)
		}
	}

	waitGroup.Done()
}

func checkOk() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s download finish!\n", url)
		count++

		if count == 27 {
			close(chanImageUrls)
			break
		}
	}

	waitGroup.Done()
}

func getPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	handleError(err, "http.get get page str")
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	handleError(err, "read all get page str")

	pageStr = string(pageBytes)

	return
}

func getImgs(url string) (urls []string) {
	pageStr := getPageStr(url)
	re := regexp.MustCompile(reImg)
	result := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("found %d imgs\n", len(result))

	for _, result := range result {
		url := result[0]
		urls = append(urls, url)
	}

	return
}

func getImgUrls(url string) {
	urls := getImgs(url)

	for _, url := range urls {
		chanImageUrls <- url
	}

	time.Sleep(2 * time.Second)

	chanTask <- url
	waitGroup.Done()
}

func main() {
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan string, 27)

	for i := 1; i <= 27; i++ {
		waitGroup.Add(1)
		go getImgUrls("https://meirentu.cc/pic/484174867058-" + strconv.Itoa(i) + ".html")
	}

	waitGroup.Add(1)
	go checkOk()

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go downloadImg()
	}

	waitGroup.Wait()
}
