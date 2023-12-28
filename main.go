package main

import (
	"log"
	"os"
)

func main() {
	//delayTime, err := util.CheckServer(config.TestAddr)
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Println(delayTime)
	//service.PGinfo()
	if err := cmd.NewApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
