package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	version := "11.151.12"
	brkdwn :=strings.Split(version,".")

	var nextVersion string

	for i:=0;i<(len(brkdwn)-1);i++{
		nextVersion += brkdwn[i]+"."
	}
	res,_:= strconv.Atoi( brkdwn[(len(brkdwn)-1)])

	nextVersion+=strconv.Itoa(res+1)
	fmt.Println(nextVersion+"-SNAPSHOT")
}
