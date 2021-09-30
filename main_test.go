package main

import (
	crand "crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Test_RandomSize(t *testing.T) {
	D := denominator(32)
	n := 6
	fmt.Println("Num of nodes:", n)
	segments := prepareSegments(D, n)
	iter := 100000
	fmt.Println("Num of attempts:", iter)

	t.Run("random sioze", func(t *testing.T) {
		maxValSize := 10000
		fmt.Println("Random volume size from 1 to", maxValSize)

		// histogram
		hist := fillRandomValues(D, hashFn_murmur3, segments, iter, maxValSize)
		reportHist(fmt.Sprintf("random size of %d attempts", iter), hist, segments)
	})

	t.Run("sha256 from random value", func(t *testing.T) {
		maxValSize := 10000
		fmt.Println("Random volume size from 1 to", maxValSize)
		// histogram
		hist := fillRandomValuesSHA256(D, hashFn_murmur3, segments, iter, maxValSize)
		reportHist(fmt.Sprintf("random size sha256 of %d attempts", iter), hist, segments)
	})

	t.Run("uuidv4", func(t *testing.T) {
		// histogram
		hist := fillUUIDv4(D, hashFn_murmur3, segments, iter)
		reportHist(fmt.Sprintf("uuidv4 of %d attempts", iter), hist, segments)
	})

}

func reportHist(title string, hist map[int]int64, segments uint32Slice) {
	var max, min float64 = 0, math.MaxFloat64
	for _, v := range hist {
		max = math.Max(float64(max), float64(v))
		min = math.Min(float64(min), float64(v))
	}

	fmt.Printf("Histogram of %q (hits min %v, hits max %v, num segments %d):\n", title, min, max, len(segments))
	for nodeNum := range segments {
		v := hist[nodeNum]
		fmt.Printf("\t node\t'%d' hits\t%d\n", nodeNum, v)
	}
}

func fillRandomValues(d uint32, h func([]byte) uint32, segments uint32Slice, iter int, maxValSize int) map[int]int64 {
	hist := map[int]int64{}
	for i := 0; i < iter; i++ {
		val := randBytes(rand.Intn(maxValSize) + 1)
		valh := h(val)

		pos := segments.numSegment(valh % d)

		hist[pos]++
	}
	return hist
}

func fillRandomValuesSHA256(d uint32, h func([]byte) uint32, segments uint32Slice, iter int, maxValSize int) map[int]int64 {
	hist := map[int]int64{}
	for i := 0; i < iter; i++ {
		val := randBytes(rand.Intn(maxValSize) + 1)
		valSha256 := sha256.Sum256(val)
		valh := h(valSha256[:])

		pos := segments.numSegment(valh % d)

		hist[pos]++
	}
	return hist
}

func fillUUIDv4(d uint32, h func([]byte) uint32, segments uint32Slice, iter int) map[int]int64 {
	hist := map[int]int64{}
	for i := 0; i < iter; i++ {
		val := uuid.New()
		valh := h(val[:])

		pos := segments.numSegment(valh % d)

		hist[pos]++
	}
	return hist
}

func randBytes(len int) []byte {
	b := make([]byte, len)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		panic(err)
	}
	return b
}
