package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	firstGoroutine()
	secondGoroutine()
	start := time.Now()
	thirdGoroutine()
	fmt.Printf("Execution time: %s\n", time.Since(start))
	fmt.Println("---")
	start2 := time.Now()
	fourthGoroutine()
	fmt.Printf("Execution time 2: %s\n", time.Since(start2))
	peerGoroutineDemo()

}

// pakai time.Sleep() untuk menunggu goroutine selesai
// ini bukan cara yang baik karena kita tidak tahu kapan goroutine selesai
func firstGoroutine() {
	go helloworld(nil)
	time.Sleep(1 * time.Second)
	goodbye(nil)
}

// pakai sync.WaitGroup buat memastikan semua goroutine selesai sebelum main goroutine selesai
func secondGoroutine() {
	var wg sync.WaitGroup
	wg.Add(2)
	go helloworld(&wg)
	go goodbye(&wg)
	wg.Wait()
}

// pakai channel untuk mengirim pesan dari goroutine ke main goroutine
func thirdGoroutine() {
	msg := make(chan string)
	go greet(msg)

	greeting := <-msg // receiving message from channel is blocking

	fmt.Println("Greeting received")
	fmt.Println(greeting)

	_, ok := <-msg
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}

// menggunakan kombinasi channel dan waitgroup untuk menunggu goroutine selesai
func fourthGoroutine() {
	var wg sync.WaitGroup
	var msg = make(chan string)
	wg.Add(1)
	go greet2(msg, &wg)
	greeting := <-msg
	wg.Wait()
	fmt.Println("Greeting received")
	fmt.Println(greeting)

	_, ok := <-msg
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

}

func helloworld(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	println("Hello, World!")
}

func goodbye(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	println("Goodbye!")
}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!")

	ch <- "Hello Rwitesh"
	close(ch) // channel perlu ditutup setelah selesai digunakan
	// untuk member tahu receiver bahwa tidak ada lagi data yang akan dikirim
	// jika tidak ditutup, receiver akan menunggu selamanya (blocking) terutama saat looping
	// tidak menutup channel tidak masalah jika resiver tidak menunggu data terus menerus (misal dengan loop)

	fmt.Println("Greeter completed")
}

func greet2(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("Greeter waiting to send greeting!")

	ch <- "Hello Rwitesh"
	wg.Done()
	close(ch)

	fmt.Println("Greeter completed")
}

// Peer goroutine: dua goroutine saling berkomunikasi (selevel)
func peerGoroutineDemo() {
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1: pengirim
	go func() {
		defer wg.Done()
		ch <- "Pesan dari goroutine 1"
		fmt.Println("Goroutine 1 selesai mengirim")
	}()

	// Goroutine 2: penerima
	go func() {
		defer wg.Done()
		msg := <-ch
		fmt.Printf("Goroutine 2 menerima: %s\n", msg)
	}()

	wg.Wait()
	fmt.Println("Peer goroutine selesai")
}
