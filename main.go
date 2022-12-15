package main
#hi
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type operation struct {
	Operands []int `json:"operands"`
}

type result struct {
	Result int `json:"result"`
}

func main() {
	op := operation{Operands: []int{1, 4}}

	data, _ := json.Marshal(op)
	postBody := bytes.NewBuffer(data)
	var resp *http.Response
	var err error
	for {
		resp, err = http.Post("http://calculator/add",
			"application/json", postBody)
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(1 * time.Second)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var res result
	json.Unmarshal(body, &res)
	fmt.Println(op.Operands[0], "+", op.Operands[1], "=", res.Result)
}
