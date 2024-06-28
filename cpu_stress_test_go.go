package main

import (
    "runtime"
    "sync"
    "time"
)

func cpuWorkload(wg *sync.WaitGroup) {
    defer wg.Done()
    for {
    }
}

func cpuStressTest(duration time.Duration) {
    var wg sync.WaitGroup
    numCPU := runtime.NumCPU()
    wg.Add(numCPU)

    for i := 0; i < numCPU; i++ {
        go cpuWorkload(&wg)
    }

    time.Sleep(duration)
    wg.Wait() // Wait for all goroutines to complete
}

func main() {
    testDuration := 10 * time.Second // duration in seconds
    println("Starting CPU stress test for", testDuration.Seconds(), "seconds...")
    cpuStressTest(testDuration)
    println("CPU stress test completed.")
}
