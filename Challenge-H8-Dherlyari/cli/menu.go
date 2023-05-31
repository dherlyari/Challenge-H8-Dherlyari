package cli

import (
	"fmt"
	"os"
)

func MainMenu() {

	fmt.Println("Selamat Datang di Challenge Hactive8!!!")
	fmt.Println("--------------------------------")

	var input string
	fmt.Println("Tekan (1) untuk mendaftar")
	fmt.Println("Tekan (2) untuk Login")
	fmt.Println("Tekan (3) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch input {
	case "1":
		Register()
	case "2":
		Login()
	case "3":
		fmt.Println("Terima kasi telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		MainMenu()
	}
}
