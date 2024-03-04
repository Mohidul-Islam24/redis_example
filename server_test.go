package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserDetails(t *testing.T) {
	//wg := sync.WaitGroup{}
	s := NewServer()
	ts := httptest.NewServer(http.HandlerFunc(s.GetUserDetails))
	for i := 1; i <= 1000; i++ {
		//wg.Add(1)
		//go func(i int) {
		id := i%100 + 1
		url := fmt.Sprintf("%s/?id=%d", ts.URL, id)
		resp, e := http.Get(url)
		if e != nil {
			t.Error(e)
		}
		user := &Users{}
		if er := json.NewDecoder(resp.Body).Decode(user); er != nil {
			t.Error(er)
		}
		fmt.Printf("%+v\n", user)
		//wg.Done()
		//}(i)
		//time.Sleep(time.Millisecond)
	}
	//wg.Wait()
	fmt.Println("number of times we hit db: ", s.dbHit)
}
