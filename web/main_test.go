package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12/httptest"
	"io"
	"os"
	"sync"
	"testing"
)

type result struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func TestSetBack(t *testing.T) {
	e := httptest.New(t, newApp().Application)
	f, _ := os.OpenFile("testlog", os.O_CREATE|os.O_WRONLY, 0666)

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			re := e.POST("/player/setback").
				WithFormField("UserId", "moqikaka001").
				WithFormField("PartnerId", 101).
				Expect().Status(httptest.StatusOK).Body().Raw()

			result := result{}
			_ = json.Unmarshal([]byte(re), &result)

			io.WriteString(f, fmt.Sprintf("协程%d:"+result.Msg+"\n", i))
		}(i)
	}
	defer f.Close()
	wg.Wait()
}

func TestBack(t *testing.T) {
	e := httptest.New(t, newApp().Application)
	f, _ := os.OpenFile("testlog", os.O_CREATE|os.O_WRONLY, 0666)

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			re := e.POST("/player/back").
				WithFormField("UserId", "moqikaka001").
				WithFormField("PartnerId", 101).
				Expect().Status(httptest.StatusOK).Body().Raw()

			result := result{}
			_ = json.Unmarshal([]byte(re), &result)

			io.WriteString(f, fmt.Sprintf("协程%d:"+result.Msg+"\n", i))
		}(i)
	}

	defer f.Close()
	wg.Wait()
}
