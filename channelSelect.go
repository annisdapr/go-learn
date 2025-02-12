package main

import "fmt"
import "runtime"

func getAverage(numbers []int, ch chan float64) {
    var sum = 0
    for _, e := range numbers {
        sum += e
    }
    ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
    var max = numbers[0]
    for _, e := range numbers {
        if max < e {
            max = e
        }
    }
    ch <- max
}

func main() {
    runtime.GOMAXPROCS(2)

    var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
    fmt.Println("numbers :", numbers)

    var ch1 = make(chan float64)
    go getAverage(numbers, ch1)

    var ch2 = make(chan int)
    go getMax(numbers, ch2)

    for i := 0; i < 2; i++ {
        select {
        case avg := <-ch1:
            fmt.Printf("Avg \t: %.2f \n", avg)
        case max := <-ch2:
            fmt.Printf("Max \t: %d \n", max)
        }
    }
}

/*
Alur Logika Program
Menyiapkan Data

Slice numbers berisi sekumpulan angka.
Dicetak ke layar untuk referensi.
Menjalankan Goroutine

Dua goroutine dijalankan:
getAverage untuk menghitung rata-rata.
getMax untuk mencari nilai maksimum.
Masing-masing hasil dikirim melalui channel (ch1 untuk rata-rata, ch2 untuk maksimum).
Menunggu dan Menampilkan Hasil

Menggunakan select untuk menerima hasil dari salah satu goroutine yang selesai lebih dulu.
Menampilkan hasil perhitungan rata-rata dan nilai maksimum.
Proses ini dilakukan sebanyak dua kali karena ada dua nilai yang dikirim.
*/