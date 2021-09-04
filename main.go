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
	bitMax := uint32(1<<bitSize) - 1
	fmt.Println("max value by bit:", bitMax)
	// denominator
	D := uint32(bitMax >> (bitSize / 6))
	fmt.Println("denominator:", D)

	fmt.Println("hasher: murmur3")
	hashFn := func(in []byte) uint32 {
		var h32 hash.Hash32 = murmur3.New32()
		h32.Write([]byte(in))
		return h32.Sum32()
	}

	numNodes := *fNumNodes
	fmt.Println("num of nodes:", numNodes)

	nodeSegments := uint32Slice{}

	left := uint32(0)
	// segment length
	L := uint32(D / uint32(numNodes))
	fmt.Println("node boundaries ( length =", L, "):")
	for i := 0; i < numNodes; i++ {
		right := L*uint32(i+1) + 1
		fmt.Println("\t#%d:", i+1, left, right)

		nodeSegments = append(nodeSegments, right)
		left = right
	}

	sort.Sort(nodeSegments)

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
