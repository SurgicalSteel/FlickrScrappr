# FlickrScrappr
Simple Flickr Scrapping Tool

This FlickrScrappr will download all specific user's original public images by scrapping on each page of user's photostream.
Original images usually have a large size. So I set default client timeout to 20s.
This tool relies heavily on Flickr's markup structure. If they change the structure, this code may not be working (and it is my job to ensure it is working).


What you need to do :
1. Set the `const UserID` in main.go, assign it with the target's userID
2. Ensure that the result folder is exist. So DO NOT delete that folder.
3. Do `dep ensure`
4. Do `go build`
5. Run it (`./FlickrScrappr`)
6. Wait......
7. Voila, all of your desired images are on `result` folder.
