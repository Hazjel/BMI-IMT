package main

import "fmt"

func main() {
	var h, w, imt float64
	var gender int
	var repeat string

	fmt.Println("Program Menghitung IMT")
	fmt.Println("==========================================")
	fmt.Println("Disusun oleh:")
	fmt.Println("1. Izzra Hilal Adityo (103012330168)")
	fmt.Println("2. Jian Hazel Sitorus (103012300187)")
	fmt.Println("")

	for {
		fmt.Println("Masukkan jenis kelamin anda (1/2)")
		fmt.Println("1. Laki-laki")
		fmt.Println("2. Perempuan")
		fmt.Scan(&gender)

		if gender == 1 { //laki-laki
			fmt.Println("")
			fmt.Print("Masukkan tinggi badan anda (m): ")
			fmt.Scan(&h) //input tinggi
			fmt.Print("Masukkan berat badan anda (kg): ")
			fmt.Scan(&w)      //input berat
			imt = w / (h * h) //rumus IMT
			fmt.Println("")
			fmt.Println("BMI anda adalah", imt) //hasil imt
			if imt < 18.5 {                     //kategori kurus
				fmt.Println("Anda kurus")
			} else if imt >= 18.5 && imt < 25 { //kategori normal
				fmt.Println("Anda normal")
			} else if imt >= 25 && imt < 30 { //kategori kelebihan berat badan
				fmt.Println("Anda kelebihan berat badan")
			} else if imt >= 30 && imt < 35 { //kategori obesitas I
				fmt.Println("Anda mengalami obesitas I")
			} else if imt >= 35 && imt < 40 { //kategori obesitas II
				fmt.Println("Anda mengalami obesitas II")
			} else if imt >= 40 { //kategori obesitas III
				fmt.Println("Anda mengalami obesitas III")
			}

		} else { //perempuan
			fmt.Println("")
			fmt.Print("Masukkan tinggi badan anda (m): ")
			fmt.Scan(&h) //input tinggi
			fmt.Print("Masukkan berat badan anda (kg): ")
			fmt.Scan(&w)      //input berat
			imt = w / (h * h) //rumus IMT
			fmt.Println("")
			fmt.Println("BMI anda adalah", imt)
			if imt < 17 { //kategori kurus
				fmt.Println("Anda kurus")
			} else if imt >= 17 && imt < 23 { //katorgi normal
				fmt.Println("Anda normal")
			} else if imt >= 23 && imt < 28 { //kategori kelebihan berat badan
				fmt.Println("Anda kelebihan berat badan")
			} else if imt >= 28 && imt < 33 { //kategori obesitas I
				fmt.Println("Anda mengalami obesitas I")
			} else if imt >= 33 && imt < 38 { //kategori obesitas II
				fmt.Println("Anda mengalami obesitas II")
			} else if imt >= 38 { //kategori obesitas III
				fmt.Println("Anda mengalami obesitas III")
			}
		}

		fmt.Println("")
		fmt.Println("Apakah anda ingin menghitung IMT lagi? (ya/tidak)")
		fmt.Scan(&repeat)

		if repeat != "ya" {
			break
		}
		fmt.Println("Terimakasih karena sudah menggunakan program ini.")
	}
}
