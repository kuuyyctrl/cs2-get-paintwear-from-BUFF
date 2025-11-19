package main

import (
	"BUFF/GetLog"
	"BUFF/db"
	"BUFF/get_wear"
	"fmt"
	"time"
)

// index 由 有无纪念品或暗金决定有的话填2无填1
func main() {
	db.Init()
	tpe := "ancient"
	GetLog.Glog(tpe)
	index := 2
	for k, ul := range GetLog.Log {
		if ul == "" {
			break
		}
		fmt.Println(ul)
		get_wear.Gwear(ul, GetLog.Gid[k], "远古"+"("+tpe+")", index)
		time.Sleep(1500 * time.Millisecond)
	}

}
