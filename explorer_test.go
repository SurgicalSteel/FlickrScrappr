package main

import (
	"fmt"
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestGetPagesLink(t *testing.T) {
	type pagesLinkTestPlan struct {
		userID       string
		page         int
		expectedLink string
	}
	testCases := make(map[string]pagesLinkTestPlan)
	testCases["normal Case"] = pagesLinkTestPlan{
		userID:       "666",
		page:         6,
		expectedLink: "https://www.flickr.com/photos/666/page6",
	}

	for k, v := range testCases {
		fmt.Println("Executing test case", k, "....")
		actual := GetPagesLink(v.userID, v.page)
		if actual != v.expectedLink {
			t.Error("Mismatched result. Expecting", v.expectedLink, "but got", actual)
		}
	}
}
func TestGetOriginalImagePageLink(t *testing.T) {
	type originalImagePageLinkTestPlan struct {
		userID       string
		imageID      string
		expectedLink string
	}
    testCases := make(map[string]originalImagePageLinkTestPlan)
    testCases["normal Case"] = originalImagePageLinkTestPlan{
        userID : "666",
        imageID: "666IAMNOTGAY666",
        expectedLink : "https://www.flickr.com/photos/666/666IAMNOTGAY666/sizes/o/",
    }
    for k, v := range testCases {
		fmt.Println("Executing test case", k, "....")
		actual := GetOriginalImagePageLink(v.userID, v.imageID)
		if actual != v.expectedLink {
			t.Error("Mismatched result. Expecting", v.expectedLink, "but got", actual)
		}
	}
}
func TestGetMainPageLink(t *testing.T) {
    type mainPageLinkTestPlan struct {
		userID       string
		expectedLink string
	}
    testCases := make (map[string]mainPageLinkTestPlan)
    testCases["normal case"] = mainPageLinkTestPlan{
        userID : "666",
        expectedLink : "https://www.flickr.com/photos/666",
    }
    for k, v := range testCases {
		fmt.Println("Executing test case", k, "....")
		actual := GetMainPageLink(v.userID)
		if actual != v.expectedLink {
			t.Error("Mismatched result. Expecting", v.expectedLink, "but got", actual)
		}
	}
}
func TestFindImageIDs(t *testing.T){
    type findImageIDsTestPlan struct {
		source       string
		expected  []string
	}
    testCases := make(map[string]findImageIDsTestPlan)
    testCases["normal case"] = findImageIDsTestPlan{
        source : "asasdweqweqwetimingCache['12132444']adsadtimingCache['123144']timingCache[]",
        expected :[]string{"12132444","123144"},
    }
    testCases["empty case"] = findImageIDsTestPlan{
        source : "asdasdaqweqwetimingCache[]asdasdahfghtimingCache['']",
        expected : []string{""},
    }
    for k, v := range testCases {
		fmt.Println("Executing test case", k, "....")
		actual := FindImageIDs(v.source)
		assert.Equal(t, v.expected, actual)
	}
}
