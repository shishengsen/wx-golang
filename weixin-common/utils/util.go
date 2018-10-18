package utils

import (
	"encoding/json"
	"math/rand"
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
