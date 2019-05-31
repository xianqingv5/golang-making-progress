package main

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var key []byte

func Init() {
	rand.Seed(time.Now().Unix())
	var err error
	key, err = hex.DecodeString("6F46756B794C5535777A3534494150326F72503155325177644E447267494849")
	if err != nil {
		panic(err)
	}
	if len(key) != 32 {
		panic("key size error, only 32 bytes key supported (64 bytes hex string)")
	}
}

func Encrypt(text string) string {
	return string(EncryptBytes([]byte(text)))
}

func EncryptBytes(textBytes []byte) []byte {
	keyBlock, _ := aes.NewCipher(key)
	size := fmt.Sprintf("%016d", len(textBytes))

	randPadding := (16 - (len(textBytes) % 16)) % 16
	for i := 0; i != randPadding; i++ {
		textBytes = append(textBytes, byte(rand.Intn(256)))
	}

	textBytes = append(textBytes, []byte(size)...)

	cipherBytes := make([]byte, 0, len(textBytes))
	blockSize := keyBlock.BlockSize()
	enc := make([]byte, blockSize)

	for len(textBytes) > 0 {
		toEnc := textBytes[:blockSize]
		textBytes = textBytes[blockSize:]
		keyBlock.Encrypt(enc, toEnc)
		cipherBytes = append(cipherBytes, enc...)
	}

	rt := make([]byte, 2*len(cipherBytes))
	n := hex.Encode(rt, cipherBytes)
	return rt[:n]
}

func Decrypt(cipher string) string {
	return string(DecryptBytes([]byte(cipher)))
}

func DecryptBytes(cipher []byte) []byte {
	keyBlock, _ := aes.NewCipher(key)

	textBytes := make([]byte, 0, len(cipher))
	cipherBytes := make([]byte, len(cipher)/2)

	if _, err := hex.Decode(cipherBytes, cipher); err != nil {
		return []byte("")
	}

	var lastBlock []byte
	blockSize := keyBlock.BlockSize()
	dec := make([]byte, blockSize)
	for len(cipherBytes) > 0 {
		if len(cipherBytes) < blockSize {
			return []byte("")
		}
		toDec := cipherBytes[:blockSize]
		cipherBytes = cipherBytes[blockSize:]
		keyBlock.Decrypt(dec, toDec)
		textBytes = append(textBytes, dec...)
		lastBlock = dec
	}

	size, _ := strconv.Atoi(string(lastBlock))
	return textBytes[:size]
}

var bakToken []string

func main() {
	f, err := os.Open("taobao_token")
	if err != nil {
		fmt.Println("open taobao token file paste_token/taobao_token error: ", err)
		return
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read taobao token content error: ", err)
		return
	}

	tokenBytes := bytes.Split(bytes.Trim(content, ", "), []byte(","))
	for _, tokenByte := range tokenBytes {
		bakToken = append(bakToken, string(tokenByte))
	}

	// bakToken = append(bakToken, "￥YwkcYWHPKzL￥")
	size := len(bakToken)
	fmt.Println(size)
	Init()
	fmt.Println(Encrypt(bakToken[rand.Intn(size)]))
	fmt.Println(Encrypt(""))
	fmt.Println(rand.Float32())
}
