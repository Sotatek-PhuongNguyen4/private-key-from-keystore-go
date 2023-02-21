package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
    inPath:="./keystore/UTC--2023-02-10T08-10-11.056313000Z--d6c65a4ca1810c2ae675edd6bdc8cf5542d0d32e"
    outPath:="./key.hex"
    password:="1"
    keyjson,e:=ioutil.ReadFile(inPath); if e != nil { panic(e) }
    key,e:=keystore.DecryptKey(keyjson,password); if e != nil { panic(e) }
    e=crypto.SaveECDSA(outPath,key.PrivateKey); if e!=nil { panic(e) }
    fmt.Println("Key saved to:",outPath)
}