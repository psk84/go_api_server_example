package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type User struct {
	idx  int    `json:"idx"`
	name string `json:"name"`
	pwd  string `json:"pwd"`
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// log.Println(r.PostForm)

	user := &User{}
	user.name = r.PostFormValue("name")
	user.pwd = r.PostFormValue("pwd")

	// log.Println("Request UserInfo [name : %s, pwd: %s", name, pwd)
	//DB 연동 유저 저장
	userIdx, success := insertUser(user.name, user.pwd)

	log.Println(userIdx)

	if userIdx != 0 && success {
		result, _ := json.Marshal(DefaultResult{RESULT_SUCCESS, RESULT_SUCCESS_CODE})
		io.WriteString(w, string(result))
	} else {
		result, _ := json.Marshal(DefaultResult{RESULT_FAIL, RESULT_FAIL_SERVER_ERROR_CODE})
		io.WriteString(w, string(result))
	}
}

func HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// log.Println(r.PostForm)

	user := &User{}
	user.name = r.PostFormValue("name")
	user.pwd = r.PostFormValue("pwd")

	// log.Println("Request UserInfo [name : %s, pwd: %s", name, pwd)
	//DB 연동 유저 저장
	isExist := selectUser(user.name, user.pwd)

	if isExist {
		result, _ := json.Marshal(DefaultResult{RESULT_SUCCESS, RESULT_SUCCESS_CODE})
		io.WriteString(w, string(result))
	} else {
		result, _ := json.Marshal(DefaultResult{RESULT_FAIL, RESULT_FAIL_WRONG_ID_OR_PWD})
		io.WriteString(w, string(result))
	}
}

func createUser(h handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}

func loginUser(h handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}
