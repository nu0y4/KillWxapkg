package main

import (
	"flag"
	"fmt"
	"github.com/Ackites/KillWxapkg/cmd"
	"log"
	"net/http"
)

var (
	appID      string
	input      string
	wxpath     string
	outputDir  string
	fileExt    string
	restoreDir bool
	pretty     bool
	noClean    bool
	hook       bool
	save       bool
	repack     string
	watch      bool
	sensitive  bool
)

func init() {
	flag.StringVar(&wxpath, "wxp", "", "微信小程序的AppID(必填)")
}

func main() {

	flag.Parse()
	fmt.Println("感谢@Ackites师傅源项目：https://github.com/Ackites/KillWxapkg")
	fmt.Println("感谢@ekkoo师傅提供的思路")
	fmt.Println("二开项目@面包狗")

	if wxpath != "" {
		fs := http.FileServer(http.Dir("./output"))

		http.Handle("/", fs)

		log.Println("Web服务开启：http://localhost:1549")
		err := http.ListenAndServe(":1549", nil)
		if err != nil {
			log.Fatal("Server failed:", err)
		}
		cmd.WatchDirectory(wxpath)
	}

}
