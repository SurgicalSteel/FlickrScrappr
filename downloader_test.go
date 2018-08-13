package main

import (
	"fmt"
	"testing"
)

func TestGetURLFileNameSuffix(t *testing.T) {
	type nameSuffixTestPlan struct {
		url            string
		expectedResult string
	}
	testCases := make(map[string]nameSuffixTestPlan)
	testCases["empty Suffix"] = nameSuffixTestPlan{
		url:            "https://www.flickr.com/photos/666129401173@N02/",
		expectedResult: "",
	}
	testCases["exist Suffix"] = nameSuffixTestPlan{
		url:            "http://www.apaajaboleh.com/photo/anak-smp-lagi-main-ular-tangga.png",
		expectedResult: "anak-smp-lagi-main-ular-tangga.png",
	}
	for k, v := range testCases {
		fmt.Println("Executing test case", k, "....")
		actual := getURLFileNameSuffix(v.url)
		if actual != v.expectedResult {
			t.Error("Mismatched result. Expecting", v.expectedResult, "but got", actual)
		}
	}
}
