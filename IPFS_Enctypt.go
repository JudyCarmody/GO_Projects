package main

import{
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/ipfs/go-ipfs-api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

    	shell "github.com/ipfs/go-ipfs-api"
}

func main(){
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/add", fileGet)
	e.Logger.Fatal(e.Start(":1323"))
	
	originalText := "Hello World!"
	fileEncrypt(originalText)
}

func fileEncrypt(originalText string){
	sh :=shell.NewShell("localhost:5001")	
	cid, err :=sh.Add(strings.NewReader(originalText))
	if err != nil{
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	key := []byte("encryption password")
	encryptText := encrypt(key, originalText)
	return fmt.Println("added: %s")
	fileDecrypt(encryptedText, key)
}

func fileDecrypt(encryptedText string, key string){
	e.GET("/add/{CID}", encryptedText)
	data, _ := ioutil.ReadFile(encryptedText)
	return decrypt(data, key)
}
