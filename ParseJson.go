package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ObjectType struct {
	Adm_max   int
	Adm_num   int
	Count     int
	Ec        int
	LevelObj  ObjectLevel
	Max_count int
	Mems      []ObjectMember
}
type ObjectLevel struct {
	Level     string
	Levelname string
}
type ObjectMember struct {
	Card            string            `json:"card"`
	Flag            int               `json:"flag"`
	G               int               `json:"g"`
	Join_time       int               `json:"join_time"`
	Last_speak_time int               `json:"last_speak_time"`
	Lv              ObjectMemberLevel `json:"lv"`
	Nick            string            `json:"nick"`
	Qage            int               `json:"qage"`
	Role            int               `json:"role"`
	Tags            string            `json:"tags"`
	Uin             int               `json:"uin"`
}
type ObjectMemberLevel struct {
	Level int `json:"level"`
	Point int `json:"point"`
}

func ParseJson(unDatas []byte) {
	var jsontype ObjectType
	err := json.Unmarshal(unDatas, &jsontype)
	checkErr(err)
	/*fmt.Printf("\r\njson-->%v\r\n", jsontype)
	fmt.Printf("%d---%d\r\n", jsontype.Count, len(jsontype.Mems))
	for i := 0; i < len(jsontype.Mems); i++ {
		fmt.Printf("%d-->%s\r\n", i, jsontype.Mems[i].Nick)
	}
	/**/

	file, _ := os.Create("1.txt")
	defer file.Close()
	//file.Write(unDatas)
	encoder := json.NewEncoder(file)
	encoder.Encode(jsontype.Mems)
	fmt.Printf("encoder end")
}
