package main

import{
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
}

const validCID = 'QmP8jTG1m9GSDJLCbeWhVSVgEzCPPwXRdCRuJtQ5Tz9Kc9'
ipfs.files.get(validCID, function (err, files) {
	files.forEach((file) => {
		console.log(file.path)
		console.log("File content >> ",file.content.toString('utf8'))
	})
})
