package denco_test

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	"github.com/naoina/denco"
)

func BenchmarkRouter_Lookup_100(b *testing.B) {
	benchmarkRouter_Lookup(b, 100)
}

func BenchmarkRouter_Lookup_300(b *testing.B) {
	benchmarkRouter_Lookup(b, 300)
}

func BenchmarkRouter_Lookup_700(b *testing.B) {
	benchmarkRouter_Lookup(b, 700)
}

func BenchmarkRouter_Build_100(b *testing.B) {
	benchmarkRouter_Build(b, 100)
}

func BenchmarkRouter_Build_300(b *testing.B) {
	benchmarkRouter_Build(b, 300)
}

func BenchmarkRouter_Build_700(b *testing.B) {
	benchmarkRouter_Build(b, 700)
}

func benchmarkRouter_Lookup(b *testing.B, n int) {
	b.StopTimer()
	router := denco.New()
	records := makeTestRecords(n)
	if err := router.Build(records); err != nil {
		b.Fatal(err)
	}
	record := pickTestRecord(records)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if r, _, _ := router.Lookup(record.Key); r != record.Value {
			b.Fail()
		}
	}
}

func benchmarkRouter_Build(b *testing.B, n int) {
	b.StopTimer()
	records := makeTestRecords(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		router := denco.New()
		if err := router.Build(records); err != nil {
			b.Fatal(err)
		}
	}
}

func makeTestRecords(n int) []denco.Record {
	records := make([]denco.Record, n)
	for i := 0; i < n; i++ {
		records[i] = denco.NewRecord("/"+randomString(50), fmt.Sprintf("testroute%d", i))
	}
	return records
}

func pickTestRecord(records []denco.Record) denco.Record {
	return records[len(records)/2]
}

func randomString(n int) string {
	const srcStrings = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789/"
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(srcStrings)-1)))
		if err != nil {
			panic(err)
		}
		buf.WriteByte(srcStrings[num.Int64()])
	}
	return buf.String()
}
