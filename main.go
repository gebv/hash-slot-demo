package main

import (
	"bufio"
	"flag"
	"fmt"
	"hash"
	"log"
	"os"
	"sort"

	"github.com/twmb/murmur3"
)

var fNumNodes = flag.Int("num", 6, "Number of nodes.")
var fBitSize = flag.Int("bit", 32, "Bit size for hasher.")
var fValue = flag.String("val", "payload123", "Incoming value.")

func main() {
	flag.Parse()

	bitSize := *fBitSize
	fmt.Println("bit size:", bitSize)
	// denominator
	D := denominator(bitSize)
	fmt.Println("denominator:", D)

	fmt.Println("hasher: murmur3")
	hashFn := hashFn_murmur3

	numNodes := *fNumNodes
	fmt.Println("num of nodes:", numNodes)

	fmt.Println("node boundaries ( length =", uint32(D/uint32(numNodes)), "):")
	nodeSegments := prepareSegments(D, numNodes)

	if isInputFromPipe() {
		scanner := bufio.NewScanner(os.Stdin)

		h := map[int]int64{}
		for scanner.Scan() {
			v := scanner.Text()
			vh := hashFn([]byte(v))
			h[nodeSegments.numSegment(vh%D)]++
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("distribution histogram:", h)
		os.Exit(0)
	} else {
		val := *fValue
		fmt.Println("input value:", val)
		valh := hashFn([]byte(val))
		fmt.Println("hash orig:", valh)
		fmt.Println("hash mod:", valh%D)
		fmt.Println("assigned to a node:", nodeSegments.numSegment(valh%D))
		os.Exit(0)
	}
}

func denominator(bitSize int) uint32 {
	bitMax := uint32(1<<bitSize) - 1
	// denominator
	return uint32(bitMax >> (bitSize / 6))
}

func prepareSegments(denominator uint32, numNodes int) uint32Slice {
	nodeSegments := uint32Slice{}

	// segment length
	L := uint32(denominator / uint32(numNodes))
	for i := 0; i < numNodes; i++ {
		right := L*uint32(i+1) + 1

		nodeSegments = append(nodeSegments, right)
	}
	sort.Sort(nodeSegments)
	return nodeSegments
}

func hashFn_murmur3(in []byte) uint32 {
	var h32 hash.Hash32 = murmur3.New32()
	h32.Write([]byte(in))
	return h32.Sum32()
}

type uint32Slice []uint32

func (x uint32Slice) Len() int           { return len(x) }
func (x uint32Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x uint32Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x uint32Slice) numSegment(in uint32) int {
	for i, v := range x {
		if v == in {
			return i
		}
		if i == 0 && in <= v {
			return i
		}
		if i > 0 && in > x[i-1] && in < v {
			return i
		}
		if i == len(x)-1 && in > v {
			return i
		}
	}
	panic(fmt.Sprintf("must not happen: in %d, arr %v", in, x))
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}
