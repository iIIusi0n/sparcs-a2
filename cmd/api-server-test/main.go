package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

const ServerURL = "https://localhost"

var Token string

type ApiTest struct {
	Name           string
	Method         string
	Body           string
	URL            string
	ExpectedStatus int
}

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	RealName    string `json:"real_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

type Hospital struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	NumberOfDoctor int     `json:"number_of_doctor"`
	Address        string  `json:"address"`
	PhoneNumber    string  `json:"phone_number"`
	CreatedAt      string  `json:"created_at"`
}

type InHospital struct {
	HospitalID int    `json:"hospital_id"`
	UserID     int    `json:"user_id"`
	CreatedAt  string `json:"created_at"`
}

type WaitingNumber struct {
	UserID     int    `json:"user_id"`
	HospitalID int    `json:"hospital_id"`
	Number     int    `json:"number"`
	CreatedAt  string `json:"created_at"`
}

type ChatRead struct {
	ChatID    int    `json:"chat_id"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type ChatRoom struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type Chat struct {
	ID        int    `json:"id"`
	RoomID    int    `json:"room_id"`
	SenderID  int    `json:"sender_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func RunTest(test ApiTest) (int, string) {
	return SendRequest(test.Method, test.URL, test.Body)
}

var (
	FailedTests = 0
	PassedTests = 0
)

func RunTests(tests []ApiTest) {
	for _, test := range tests {
		status, body := RunTest(test)
		if status != test.ExpectedStatus {
			log.Println("Failed:", test.Name, "status:", status, "body:", body)
			FailedTests++
		} else {
			log.Println("Passed:", test.Name, "status:", status, "body:", body)
			PassedTests++
		}
	}
}

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
	TestHospitalRouter()
	TestChatRouter()

	log.Println("Passed tests:", PassedTests)
	log.Println("Failed tests:", FailedTests)
}

func CreateTokenForDebug() {
	newUser := User{
		Username:    RandomString(10),
		RealName:    "Test User",
		Email:       RandomString(10) + "@test.com",
		PhoneNumber: "01012345678",
	}
	bodyData, _ := json.Marshal(newUser)

	test := ApiTest{
		Name:   "Create token for debug",
		Method: "POST",
		Body:   string(bodyData),
		URL:    ServerURL + "/api/v1/debug/token",
	}
	status, body := RunTest(test)
	if status != 200 {
		panic("Failed to create token for debug: " + body)
	}

	var response struct {
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
	newUser := User{
		Username:    RandomString(10),
		RealName:    "Test User",
		Email:       RandomString(10) + "@test.com",
		PhoneNumber: "01012345678",
	}

	updatedUser := User{
		ID:          1,
		Username:    RandomString(10),
		RealName:    "Test User",
		Email:       RandomString(10) + "@test.com",
		PhoneNumber: "01012345678",
	}

	bodyData, _ := json.Marshal(newUser)
	bodyData2, _ := json.Marshal(updatedUser)

	tests := []ApiTest{
		{
			Name:           "Get logged in user",
			Method:         "GET",
			URL:            ServerURL + "/api/v1/user",
			ExpectedStatus: 200,
		},
		{
			Name:           "Get user by ID",
			Method:         "GET",
			URL:            ServerURL + "/api/v1/user/1",
			ExpectedStatus: 200,
		},
		{
			Name:           "Create user",
			Method:         "POST",
			Body:           string(bodyData),
			URL:            ServerURL + "/api/v1/user",
			ExpectedStatus: 200,
		},
		{
			Name:           "Update user",
			Method:         "PATCH",
			Body:           string(bodyData2),
			URL:            ServerURL + "/api/v1/user",
			ExpectedStatus: 200,
		},
	}

	RunTests(tests)
}

func TestHospitalRouter() {
	newHospital1 := Hospital{
		Name:           "Test Hospital",
		Latitude:       37.123456,
		Longitude:      127.123456,
		NumberOfDoctor: 10,
		Address:        "Test Address",
		PhoneNumber:    "01012345678",
	}

	newHospital2 := Hospital{
		Name:           "Test Hospital 2",
		Latitude:       37.123473,
		Longitude:      127.123461,
		NumberOfDoctor: 10,
		Address:        "Test Address",
		PhoneNumber:    "01012345678",
	}

	updatedHospital := Hospital{
		ID:             1,
		Name:           "Test Hospital",
		Latitude:       37.123456,
		Longitude:      127.123456,
		NumberOfDoctor: 10,
		Address:        "Updated Address",
		PhoneNumber:    "01012345678",
	}

	newWaitingNumber := WaitingNumber{
		UserID:     1,
		HospitalID: 1,
		Number:     13,
	}

	newWaitingNumber2 := WaitingNumber{
		UserID:     2,
		HospitalID: 1,
		Number:     5,
	}

	bodyData, _ := json.Marshal(newHospital1)
	bodyData2, _ := json.Marshal(newHospital2)
	bodyData3, _ := json.Marshal(updatedHospital)

	bodyData4, _ := json.Marshal(newWaitingNumber)
	bodyData5, _ := json.Marshal(newWaitingNumber2)

	tests := []ApiTest{
		{
			Name:           "Create hospital",
			Method:         "POST",
			Body:           string(bodyData),
			URL:            ServerURL + "/api/v1/hospital",
			ExpectedStatus: 200,
		},
		{
			Name:           "Create hospital",
			Method:         "POST",
			Body:           string(bodyData2),
			URL:            ServerURL + "/api/v1/hospital",
			ExpectedStatus: 200,
		},
		{
			Name:           "Get hospitals",
			Method:         "GET",
			URL:            ServerURL + "/api/v1/hospital",
			ExpectedStatus: 200,
		},
		{
			Name:           "Get hospital by ID",
			Method:         "GET",
			URL:            ServerURL + "/api/v1/hospital/1",
			ExpectedStatus: 200,
		},
		{
			Name:           "Update hospital",
			Method:         "PATCH",
			Body:           string(bodyData3),
			URL:            ServerURL + "/api/v1/hospital/1",
			ExpectedStatus: 200,
		},
		{
			Name:           "Get updated hospital",
			Method:         "GET",
			URL:            ServerURL + "/api/v1/hospital/1",
			ExpectedStatus: 200,
		},
		{
			Name:           "Create waiting number",
			Method:         "POST",
			Body:           string(bodyData4),
			URL:            ServerURL + "/api/v1/hospital/1/waiting",
			ExpectedStatus: 200,
		},
		{
			Name:           "Create waiting number",
			Method:         "POST",
			Body:           string(bodyData5),
			URL:            ServerURL + "/api/v1/hospital/1/waiting",
			ExpectedStatus: 200,
		},
		{
			Name:           "Get hospital waiting number",
			Method:         "GET",
			URL:            ServerURL + "/api/v1/hospital/1/waiting",
			ExpectedStatus: 200,
		},
	}

	RunTests(tests)
}

func TestChatRouter() {
	newChatRoom := ChatRoom{
		Name: "Test Room",
	}

	newChat := Chat{
		RoomID:   1,
		SenderID: 1,
		Content:  "Test Content",
	}

	bodyData, _ := json.Marshal(newChatRoom)
	bodyData2, _ := json.Marshal(newChat)

	tests := []ApiTest{
		{
			Name:           "Create chat room",
			Method:         "POST",
			Body:           string(bodyData),
			URL:            ServerURL + "/api/v1/chat/room",
			ExpectedStatus: 200,
		},
		{
			Name:           "Get chat rooms",
			Method:         "GET",
			URL:            ServerURL + "/api/v1/chat/room",
			ExpectedStatus: 200,
		},
		{
			Name:           "Get chat room by ID",
			Method:         "GET",
			URL:            ServerURL + "/api/v1/chat/room/1",
			ExpectedStatus: 200,
		},
		{
			Name:           "Create chat",
			Method:         "POST",
			Body:           string(bodyData2),
			URL:            ServerURL + "/api/v1/chat/1",
			ExpectedStatus: 200,
		},
		{
			Name:           "Get chat by ID",
			Method:         "GET",
			URL:            ServerURL + "/api/v1/chat/1",
			ExpectedStatus: 200,
		},
	}

	RunTests(tests)
}
