package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

const ServerURL = "https://localhost"

var Token string

func RandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func SendRequest(method, url, body string) (int, string) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	if Token != "" {
		req.Header.Set("Cookie", "token="+Token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyStr := ""
	if resp.Body != nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyStr = string(bodyBytes)
	}

	return resp.StatusCode, bodyStr
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	CreateTokenForDebug()

	TestUserRouter()
}

func CreateTokenForDebug() {
	url := ServerURL + "/api/v1/debug/token"

	user := `{
		"username": "` + RandomString(10) + `",
		"real_name": "Test User",
		"email": "",
		"phone_number": "01012345678"
}`

	status, body := SendRequest("POST", url, user)
	if status != 200 {
		panic("Failed to create token: " + body)
	}

	var response struct {
		UID   int    `json:"uid"`
		Token string `json:"token"`
	}
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		panic(err)
	}

	Token = response.Token

	log.Println("Token:", Token)
}

func TestUserRouter() {
	url := ServerURL + "/api/v1/user/"

	status, body := SendRequest("GET", url, "")
	if status != 200 {
		panic("Failed to get user: " + body)
	}

	log.Println("User:", body)

	var response struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		RealName    string `json:"real_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		CreatedAt   string `json:"created_at"`
	}
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		panic(err)
	}

	url = ServerURL + "/api/v1/user/" + strconv.Itoa(response.ID)
	status, body = SendRequest("GET", url, "")
	if status != 200 {
		panic("Failed to get user: " + body)
	}

	log.Println("User:", body)

	url = ServerURL + "/api/v1/user/"
	response.Email = RandomString(10) + "@updated.com"
	user, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	status, body = SendRequest("PATCH", url, string(user))
	if status != 200 {
		panic("Failed to update user: " + body)
	}

	log.Println("User updated:", body)

	url = ServerURL + "/api/v1/user/"
	status, body = SendRequest("GET", url, "")
	if status != 200 {
		panic("Failed to get user: " + body)
	}

	log.Println("User:", body)

	url = ServerURL + "/api/v1/user/"
	status, body = SendRequest("POST", url, `{
		"username": "`+RandomString(10)+`",
		"real_name": "Test User",
		"email": "",
		"phone_number": "01012345678"
	}`)
	if status != 200 {
		panic("Failed to create user: " + body)
	}

	log.Println("User created:", body)
}
