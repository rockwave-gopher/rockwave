package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	http.HandleFunc("/form",showData)
	http.HandleFunc("/proposal",proposal)
	http.ListenAndServe(":3000", nil)
}

func proposal(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	body,err := ioutil.ReadAll(r.Body)
	if err !=nil{
		fmt.Println("there has occured errors.")
	}
	fmt.Println(string(body))
}

func showData(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	if(r.Method == "GET"){
		//return json 串
		w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json") //返回数据格式是json
		msg := `{
  "labelPosition": "left",
  "span": 13,
  "formDesc": {
    "insured_type": {
      "options": [
        {
          "text": "个人",
          "value": 1
        },
        {
          "text": "团体",
          "value": 2
        }
      ],
      "label": "投保类型",
      "type": "select",
      "required": false
    },
    "insured_name": {
      "label": "姓名",
      "type": "input"
    },
    "insured_id_type": {
      "options": [
        {
          "text": "身份证",
          "value": 1
        },
        {
          "text": "军官证",
          "value": 2
        },
        {
          "text": "护照",
          "value": 3
        }
      ],
      "label": "证件类型",
      "type": "select"
    },
    "insured_id": {
      "label": "证件编号",
      "type": "input"
    },
    "insured_mail": {
      "label": "邮箱",
      "type": "input"
    },
    "insured_contact": {
      "label": "联系电话",
      "type": "input"
    },
    "proposal_quota": {
      "options": [
        {
          "text": "1",
          "value": 1
        },
        {
          "text": "2",
          "value": 2
        },
        {
          "text": "3",
          "value": 3
        }
      ],
      "label": "投保份数",
      "type": "select"
    },
    "insured_relation": {
      "options": [
        {
          "text": "本人",
          "value": 1
        },
        {
          "text": "配偶",
          "value": 2
        },
        {
          "text": "子女",
          "value": 3
        }
      ],
      "label": "与被保人关系",
      "type": "select",
      "attrs": {
        "automaticDropdown": true
      }
    }
  }
}`
		w.Write([]byte(msg))
		}
	if(r.Method == "POST"){

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// handle error
		}

		fmt.Println(string(body))
	}

}
