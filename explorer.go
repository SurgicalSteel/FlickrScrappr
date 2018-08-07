package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

func GetPagesLink(userID string, page int) string {
	return fmt.Sprintf("https://www.flickr.com/photos/%s/page%d", userID, page)
}

func GetOriginalImagePageLink(userID, imageID string) string {
	return fmt.Sprintf("https://www.flickr.com/photos/%s/%s/sizes/o/", userID, imageID)
}

func GetMainPageLink(userID string) string {
	return fmt.Sprintf("https://www.flickr.com/photos/%s", userID)
}

func GetAllLinks(respBody []byte) []html.Token {
	//reader, _ := ioutil.ReadAll(respBody)
	//fmt.Println(string(reader))
	links := make([]html.Token, 0)
	tokenizr := html.NewTokenizer(bytes.NewReader(respBody))
	for {
		tt := tokenizr.Next()
		switch {
		case tt == html.ErrorToken:
			return links
		case tt == html.StartTagToken:
			t := tokenizr.Token()
			if t.Data == "a" {
				links = append(links, t)
			}
		}
	}
	return links
}

func GetMaximumPageIndex(allPageLinks []html.Token) int {
	fmt.Println("Started getting maximum available page...")
	var maxi int = 0
	// let the brute force begins
	for _, vl := range allPageLinks {
		for _, va := range vl.Attr {
			if va.Key == "data-track" {
				if strings.Contains(va.Val, "pagination") && strings.Contains(va.Val, "Click") {

					rawIndexString := strings.Trim(strings.Trim(va.Val, "pagination"), "Click")
					rawIndex, _ := strconv.Atoi(rawIndexString)
					maxi = cmaxi(maxi, rawIndex)
				}

			}
		}
	}
	return maxi
}

func GetPageImageIDs(allPageLinks []html.Token) []string {
	fmt.Println("Started getting all image IDs in a page...")
	result := make([]string, 0)
	for _, vl := range allPageLinks {
		for _, va := range vl.Attr {
			if va.Key == "href" {
				pref := fmt.Sprintf("/photos/%s/", UserID)
				suff := "/"
				if strings.Contains(va.Val, pref) {
					currentPhotoID := strings.Trim(strings.Trim(va.Val, pref), suff)
					fmt.Println(va.Val)
					fmt.Println(currentPhotoID)
					result = append(result, currentPhotoID)
				}
			}
		}
	}

	return result
}

func GetOriginalImageLink(originalSizePage []byte) string {
	fmt.Println("Started getting original image link...")
	//found := false
	pageRaw := bytes.NewReader(originalSizePage)
	originalImageLink := ""
	doc, err := goquery.NewDocumentFromReader(pageRaw)
	if err != nil {
		fmt.Println("[GOQUERY] Failed to initialize document. ", err)
		return ""
	}
	val, ex := doc.Find("#allsizes-photo").Find("img").Attr("src")
	if !ex {
		fmt.Println("[GOQUERY] Original Image does not exist")
		return ""
	}
	originalImageLink = val

	return originalImageLink
}

func FindImageIDs(s string) []string {
	res := make([]string, 0)
	it := 0
	currID := ""
	for it < len(s) {
		if it < len(s)-13 {
			if s[it:it+13] == "timingCache['" {
				itx := it + 13
				for s[itx:itx+1] != "'" {
					currID += s[itx : itx+1]
					itx++
				}
				itx--
				res = append(res, currID)
				currID = ""
			}
		} else {
			break
		}
		it++
	}
	return res
}

func cmaxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}
