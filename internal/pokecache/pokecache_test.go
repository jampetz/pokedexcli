package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Minute)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddCache(t *testing.T) {
	cache := NewCache(time.Minute)
	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key1",
			inputVal: []byte("val123213213"),
		},
		{
			inputKey: "214",
			inputVal: []byte("123get"),
		},
		{
			inputKey: "key-123",
			inputVal: []byte("123-val"),
		},
		{
			inputKey: "",
			inputVal: []byte("val21321"),
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, cs.inputVal)

		actual, ok := cache.Get(cs.inputKey)
		if !ok {
			t.Errorf("%s not found", cs.inputVal)
			continue
		}
		if string(actual) != string(cs.inputVal) {
			t.Errorf("%s doesn't match %s", actual, cs.inputVal)
			continue
		}
	}
}

func TestReapSuccess(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("value123"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}

func TestReapErr(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("value123"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s should not have been reaped", keyOne)
	}
}
