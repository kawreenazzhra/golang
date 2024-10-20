package main

import "fmt"

func main(){
	var nilai, jam, mnt, dtk  int

	fmt.Scan(&nilai)
	jam = nilai / 3600
	mnt = nilai % 3600 / 60
	dtk = nilai % 60

	fmt.Println(jam,mnt,dtk)


}