package main

import (
	"math/rand"
	"testing"
	"time"
)

const (
	numItems = 100 // change this to see how number of items affects speed
	keyLen   = 10
	valLen   = 20
)

func gen(length int) string {
	var s string
	for i := 0; i < length; i++ {
		s += string(rand.Int31n(126-48) + int32(48))
	}
	return s
}

type Item struct {
	Key string
	Val string
}

/********************
 * A typical key/value storage, once as map[string]string, once
 * as []*Item{string,string}.
 */

func populateKV(theArr []*Item, theMap map[string]string) string {
	var query string
	var k, v string

	rand.Seed(time.Now().UnixNano())

	// pick one of the items at random to be the lookup key
	queryI := rand.Int31n(numItems)

	for i := 0; i < numItems; i++ {
		k = gen(keyLen)
		v = gen(valLen)
		theMap[k] = v
		theArr[i] = &Item{Key: k, Val: v}

		if i == int(queryI) {
			query = k
		}
	}
	return query
}

func BenchmarkKVItemSlice(b *testing.B) {
	var found bool
	theMap := make(map[string]string)
	theArr := make([]*Item, numItems)
	q := populateKV(theArr, theMap)

	b.ResetTimer()
	var j int
	for i := 0; i < b.N; i++ {
		for j = 0; j < numItems; j++ {
			if theArr[j].Key == q {
				found = true
				continue
			}
		}
	}
	if !found {
		b.Fail()
	}
}

func BenchmarkKVStringMap(b *testing.B) {
	var found bool
	theMap := make(map[string]string)
	theArr := make([]*Item, numItems)
	q := populateKV(theArr, theMap)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if theMap[q] != "" {
			found = true
		}
	}
	if !found {
		b.Fail()
	}
}

/********************
 * A set of integers, once as []int, once as map[int]struct{}
 */

func populateInts(theArr []int, theMap map[int]struct{}) int {
	var query int
	rand.Seed(time.Now().UnixNano())

	// pick one of the items at random to be the lookup key
	queryI := rand.Int31n(numItems)

	for i := 0; i < numItems; i++ {
		num := rand.Int()
		theArr[i] = num
		theMap[num] = struct{}{}

		if i == int(queryI) {
			query = num
		}
	}
	return query
}

func BenchmarkSetIntSlice(b *testing.B) {
	var found bool
	theMap := make(map[int]struct{})
	theArr := make([]int, numItems)
	q := populateInts(theArr, theMap)

	b.ResetTimer()
	var j int
	for i := 0; i < b.N; i++ {
		for j = 0; j < numItems; j++ {
			if theArr[j] == q {
				found = true
				continue
			}
		}
	}
	if !found {
		b.Fail()
	}
}

func BenchmarkSetIntMap(b *testing.B) {
	var found bool
	theMap := make(map[int]struct{})
	theArr := make([]int, numItems)
	q := populateInts(theArr, theMap)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, ok := theMap[q]; ok {
			found = true
		}
	}
	if !found {
		b.Fail()
	}
}
