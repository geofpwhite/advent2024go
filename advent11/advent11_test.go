package main

import (
	"os"
	"runtime/pprof"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	{
		// Create a file to store the CPU profile
		f, err := os.Create("cpu.prof")
		if err != nil {
			b.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // Start CPU profiling
		if err := pprof.StartCPUProfile(f); err != nil {
			b.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
		for range b.N {
			Main()
		}
	}
}
