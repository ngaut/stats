package stats

import (
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

const (
	testCount = 100
)

func TestMarshal(t *testing.T) {
	go func() {
		http.ListenAndServe(":5555", nil)
	}()

	time.Sleep(time.Second)

	for i := 0; i < testCount; i++ {
		PubInt("name", 100)
		Publish("stringkey", "string")
	}

	resp, err := http.Get("http://localhost:5555/debug/stats")
	if err != nil {
		t.Fatal("error getting stats")
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("error read stats body")
	}

	js, err := simplejson.NewJson(b)
	if err != nil {
		println(string(b))
		t.Fatal("not a valid json " + err.Error())
	}

	arr, err := js.Get("name").Array()
	if err != nil {
		t.Fatal("json not match")
	}

	if len(arr) != testCount {
		t.Fatal("count not match")
	}

	arr, err = js.Get("stringkey").Array()
	if err != nil {
		t.Fatal("json not match")
	}

	if len(arr) != testCount {
		t.Fatal("count not match")
	}

}

func TestMarshalMax100(t *testing.T) {
	go func() {
		http.ListenAndServe(":5555", nil)
	}()

	time.Sleep(time.Second)

	for i := 0; i < testCount; i++ {
		PubInt("TestMarshalMax100", 100)
		Publish("TestMarshalMax100stringkey", "string")
	}

	resp, err := http.Get("http://localhost:5555/debug/stats")
	if err != nil {
		t.Fatal("error getting stats")
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("error read stats body")
	}

	js, err := simplejson.NewJson(b)
	if err != nil {
		println(string(b))
		t.Fatal("not a valid json " + err.Error())
	}

	arr, err := js.Get("TestMarshalMax100").Array()
	if err != nil {
		t.Fatal("json not match")
	}

	if len(arr) != testCount {
		t.Fatal("count not match", testCount)
	}

	arr, err = js.Get("TestMarshalMax100stringkey").Array()
	if err != nil {
		t.Fatal("json not match")
	}

	if len(arr) != testCount {
		t.Fatal("count not match", testCount)
	}
}

func TestInc(t *testing.T) {
	go func() {
		http.ListenAndServe(":5555", nil)
	}()

	time.Sleep(time.Second)

	for i := 0; i < testCount; i++ {
		Inc("xxx")
	}

	resp, err := http.Get("http://localhost:5555/debug/stats")
	if err != nil {
		t.Fatal("error getting stats")
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("error read stats body")
	}

	js, err := simplejson.NewJson(b)
	if err != nil {
		println(string(b))
		t.Fatal("not a valid json " + err.Error())
	}

	arr, err := js.Get("xxx").Array()
	if err != nil {
		t.Fatal("json not match")
	}

	if cnt, _ := arr[0].(json.Number).Int64(); cnt != testCount {
		fmt.Printf("%v-%v\n", arr[0], testCount)
		t.Fatal("count not match")
	}
}
