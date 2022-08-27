package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/panjf2000/ants/v2"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Args struct {
	I    int
	MaxI int
}

func includes(arr []string, a string) bool {
	for _, element := range arr {
		if element == a {
			return true
		}
	}

	return false
}

func signatureByPattern(profanityFn string, expectedSignature []string, maxI int) {
	wg := sync.WaitGroup{}

	threads := runtime.NumCPU() * 10

	ar := strings.Split(profanityFn, "(")
	if len(ar) != 2 {
		panic("invalid function pattern")
	}
	profanityFnName := ar[0]
	fnArgVariant := "(" + ar[1]

	VariantsCovered := 0
	p, _ := ants.NewPoolWithFunc(threads, func(i interface{}) {
		args := i.(Args)
		defer wg.Done()

		for i := args.I; i < args.MaxI; i++ {
			VariantsCovered++
			if VariantsCovered%50000000 == 0 {
				fmt.Println(VariantsCovered)
			}
			suffix := fmt.Sprintf("%x", i)
			variant := profanityFnName + "_" + suffix + fnArgVariant
			signature := crypto.Keccak256Hash([]byte(variant)).Hex()[0:10]

			if !includes(expectedSignature, signature) {
				continue
			}

			fmt.Println(variant, signature)
		}

	})

	divider := maxI / threads
	for i := 0; i < threads; i++ {
		wg.Add(1)
		_ = p.Invoke(Args{I: i * divider, MaxI: i*divider + divider})
	}

	wg.Wait()
	p.Release()
}

func main() {
	argLength := len(os.Args[1:])
	fmt.Printf("Arg length is %d\n", argLength)

	if argLength < 1 {
		panic("first arg should be pattern: fn_example_pattern")
	}
	pattern := os.Args[1]

	sigs := "0x00000001,0x00000002,0x00000003,0x00000004,0x00000005,0x00000006,0x00000007,0x00000008,0x00000009,0x00000010"
	if argLength == 2 {
		sigs = os.Args[2]
		fmt.Println("Using arg sigs: ", sigs)
	} else {
		fmt.Println("Using default sigs: ", sigs)
	}
	arr := strings.Split(sigs, ",")

	fmt.Println("CPU core count: ", runtime.NumCPU())

	fmt.Println("Benchmark started")
	t := time.Now()
	signatureByPattern(pattern, arr, 10000)
	fmt.Println("Benchmark ended\n10k variants in ", time.Since(t))

	signatureByPattern(pattern, arr, 300000000000)
}
