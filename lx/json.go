package main

import (
    "fmt"
    "encoding/json"
)

type User struct {

    //`json:"-"` // 直接忽略字段
    //`json:"msg_name"` // 对应json字段名
    //`json:"msg_name,omitempty"` // 如果为空置则忽略字段
    UserName string `json:"username"`
    PassWord string `json:"password"`
}

func main(){

    user := User{
        UserName: "tyming",
        PassWord: "1234567890",//这里逗号不能少
    }
    //json.Marshal() 编写为json格式
    str, _ := json.Marshal(user)
    fmt.Printf("%s\n",str)


}
