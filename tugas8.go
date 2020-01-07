package main

import "fmt"
import "runtime"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var pesan = []string{"Apa Kabar Teman-Teman"}
	fmt.Println("\nMengirimkan Pesan...\n")
	var waktu = make(chan int)

	go kirim_data(pesan, waktu)
	terima_data(pesan, waktu)
}

func kirim_data(pesan []string, channel_1 chan<- int) {
	for i := 0; true; i++ {
		channel_1 <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}
func terima_data(pesan []string, channel_1 <-chan int) {
loop:
	for {
		select {
		case data := <-channel_1:
			fmt.Print("Pesan Diterima -> ", pesan, " | Waktu Diterima -> ", data, " Detik", "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("\nTimeout, Tidak ada aktivitas dalam 5 detik\n")
			break loop
		}
	}
}
