package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	args := os.Args
	arg, err := strconv.Atoi(args[1])

	studentList := []Student{
		{Nama: "Raditeo", Alamat: "Semarang", Pekerjaan: "Pranata Komputer UIN Walisongo", Alasan: "Peningkatan skill pegawai"},
		{Nama: "Satria", Alamat: "Kendal", Pekerjaan: "Pranta Komputer UIN Walisongo", Alasan: "Peningkatan skill pegawai"},
		{Nama: "Mubin", Alamat: "Demak", Pekerjaan: "Pegawai BLU UIN Walisongo", Alasan: "Peningkatan skill pegawai"},
		{Nama: "Wahyu", Alamat: "Kudus", Pekerjaan: "Pegawai BLU UIN Walisongo", Alasan: "Peningkatan skill pegawai"},
		{Nama: "Akhmad", Alamat: "Semarang", Pekerjaan: "Pranta Komputer UIN Walisongo", Alasan: "Peningkatan skill pegawai"},
	}

	if err == nil {
		fmt.Println("Siswa nomor urut:", arg)
		PrintStudent(studentList[arg-1])
	}
}

func PrintStudent(s Student) {
	fmt.Println("Nama :", s.Nama)
	fmt.Println("Alamat :", s.Alamat)
	fmt.Println("Pekerjaan :", s.Pekerjaan)
	fmt.Println("Alasan :", s.Alasan)
}
