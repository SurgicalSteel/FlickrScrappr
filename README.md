# FlickrScrappr
Simple Flickr Scrapping Tool

This FlickrScrappr is a tool that you can use to download all specific user's original public images by scrapping on each page of user's photostream.


### You have been warned
This tool relies heavily on Flickr's markup structure. If they change the structure, this code may not be working (and it is my job to ensure it is working). The contributors are not responsible for anything (including any abusing act) that you do by using the code in this repository.

### Usage
What you need to do :
1. Set the `const UserID` in main.go, assign it with the target's userID[1].
2. Ensure that the result folder is exist. So DO NOT delete that folder.
3. Do `dep ensure`
4. Do `go build`
5. Run it (`./FlickrScrappr`)
6. Wait......
7. Voila, all of your desired images are on `result` folder.

[1] For Example :
If the target url page is https://www.flickr.com/photos/666129401173@N02/
then the user ID is 666129401173@N02

### This repository has been archived as there are no new improvements.
