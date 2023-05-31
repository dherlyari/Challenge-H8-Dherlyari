package cli

import (
	"Challenge-H8-Dherlyari/config"
	"Challenge-H8-Dherlyari/entity"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Register() {

	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("-------Register-------")

	var username string
	fmt.Println("Masukan username anda : ")
	username, _ = consoleReader.ReadString('\n')
	username = strings.TrimSpace(username)

	var existingUser entity.Users
	err := config.DB.Where("username = ?", username).First(&existingUser).Error
	if err == nil {
		fmt.Println("Username sudah digunakan. Silahkan gunakan username lain.")
		Register()
		return
	}

	fmt.Println("Masukan password anda : ")
	password, _ := consoleReader.ReadString('\n')
	password = strings.TrimSpace(password)

	user := entity.Users{
		Username: username,
		Password: password,
	}

	err = config.DB.Create(&user).Error
	if err != nil {
		fmt.Println("Terjadi kesalahan. Silahkan coba lagi.")
		Register()
		return
	}

	fmt.Println("Register Berhasil!!!")
	MainMenu()
}

func Login() {

	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("------Login------")

	var username string
	fmt.Println("Masukan username anda : ")
	username, _ = consoleReader.ReadString('\n')
	username = strings.TrimSpace(username)

	var password string
	fmt.Println("Masukan password anda : ")
	password, _ = consoleReader.ReadString('\n')
	password = strings.TrimSpace(password)

	var user entity.Users
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		fmt.Println("Username atau password salah. Silahkan login kembali")
		Login()
		return
	}

	err = config.DB.Where("Password = ?", password).First(&user).Error
	if err != nil {
		fmt.Println("Username atau password salah. Silahkan login kembali.")
		Login()
		return
	}

	fmt.Printf("Selamat datang, %s!!!\n", username)
	MainMenu()
}
