package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/tamalsaha/go-oneliners"
	"encoding/json"
	"strings"
	//util "github.com/kubeware/messenger/test/e2e/framework"
	"log"
)

func init() {
	fmt.Println(0)
}
func init() {
	fmt.Println(1)
}

func main() {
	//subpackage.Check()
	//CheckAgain("asdf")


		fmt.Println("Hello, playground")
		chat := "test-msg: Hello world from kubeware/messenger :D"
		client := http.Client{
			Timeout: time.Minute * 5,
		}

		url := "http://api.hipchat.com/v2/room/1214663/history?end-date=2018-06-08T10:29:11.961878722%2B06:00"
		//+util.GetDateString(time.Now())
		fmt.Println(">>>>>>>>>", url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(">>>>>>>>>>> 01", err)
		}
		req.Header.Set("Authorization", "Bearer "+"qOTqXahya2OoG4kfuWA5BZfXIomILcVEsWEzCogN")
		// =============================
		oneliners.PrettyJson(*req, "req to see hist")
		// =============================
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(">>>>>>>>>>> 02", err)
		}
		// =============================
		oneliners.PrettyJson(*resp, "resp to see hist")
		// =============================
		defer resp.Body.Close()
		type msg struct {
			Type string
			Message string
		}
		var msgs struct{
			Items []msg
		}
		err = json.NewDecoder(resp.Body).Decode(&msgs)
		if err != nil {
			log.Fatal(">>>>>>>>>>> 03", err)
		}
		// =============================
		oneliners.PrettyJson(msgs, "req.body to see hist")
		// =============================
		found := false
		for i := 0; !found && i < len(msgs.Items); i++ {
			if msgs.Items[i].Type == "notification" && msgs.Items[i].Message != "" {
				found = found || strings.Contains(msgs.Items[i].Message, chat)
			}
		}
		if err != nil {
			log.Fatal(">>>>>>>>>>> 04", err)
		}

		fmt.Println("Success")
}
