package utils

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"math/rand"
	"net/url"
	"os"
	"time"
)

const (
	temp_file_path		=		"/tmp/wx-go/"
)

var RANDOM_STR []string = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u",
"v","w","x","y","z","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y",
"Z","0","1","2","3","4","5","6","7","8","9"}

func RandomStr() string {
	var finalStr string
	for i := 0; i < 16; i +=1 {
		finalStr += RANDOM_STR[rand.Intn(len(RANDOM_STR))]
	}
	return finalStr
}

func Interface2byte(itf interface{}) []byte {
	data, err := json.Marshal(itf)
	if err != nil {
		panic(err)
	}
	return data
}

func Byte2Inteface(data []byte, e *interface{}) interface{} {
	err := json.Unmarshal(data, e)
	if err != nil {
		panic(err)
	}
	return *e
}

func UrlEncode(s string) string {
	finalUrl, err := url.QueryUnescape(s)
	if err != nil {
		panic(err)
	}
	return finalUrl
}

func CreateTempFile(data []byte) *os.File {
	tmpFile, outputError := os.OpenFile(tempFilePath(), os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		panic(outputError)
	}
	_, err := tmpFile.Write(data)
	if err != nil {
		panic(err)
	}
	return tmpFile
}

func tempFilePath() string {
	_u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return temp_file_path + _u.String()
}

func TimeFormatToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}