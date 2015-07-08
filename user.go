package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, string("Hello World"))
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// log.Println(r.PostForm)

	user := &User{}
	user.UserId = r.PostFormValue("name")
	user.Pwd = r.PostFormValue("pwd")

	// log.Println("Request UserInfo [name : %s, pwd: %s", name, pwd)
	//DB 연동 유저 저장
	var result = insertUser(user.UserId, user.Pwd)

	if result {
		result, _ := json.Marshal(DefaultResult{RESULT_SUCCESS, RESULT_SUCCESS_CODE})
		io.WriteString(w, string(result))
	} else {
		result, _ := json.Marshal(DefaultResult{RESULT_FAIL, RESULT_FAIL_WRONG_ID_OR_PWD})
		io.WriteString(w, string(result))
	}
}

func HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// log.Println(r.PostForm)

	user := &User{}
	user.UserId = r.PostFormValue("name")
	user.Pwd = r.PostFormValue("pwd")

	// log.Println("Request UserInfo [name : %s, pwd: %s", name, pwd)
	//DB 연동 유저 추출
	var result = &User{}
	result = selectUser(user.UserId, user.Pwd)

	if result != nil {
		result, _ := json.Marshal(DefaultResult{RESULT_SUCCESS, RESULT_SUCCESS_CODE})
		io.WriteString(w, string(result))
	} else {
		result, _ := json.Marshal(DefaultResult{RESULT_FAIL, RESULT_FAIL_WRONG_ID_OR_PWD})
		io.WriteString(w, string(result))
	}
}

func HandleFindUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// log.Println(r.PostForm)

	user := &User{}
	user.UserId = r.PostFormValue("name")
	user.Pwd = r.PostFormValue("pwd")

	// log.Println("Request UserInfo [name : %s, pwd: %s", name, pwd)
	//DB 연동 유저 추출
	var result = &User{}
	result = selectUser(user.UserId, user.Pwd)

	if result != nil {
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

func findUser(h handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}
