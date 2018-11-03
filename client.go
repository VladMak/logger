package main

import(
	"fmt"
	"./logger"
	"time"
)

func main(){
	var input int
	logger_this := logger.Logger{"Logger"}
	fmt.Println("Что вы хотите записать?")
	fmt.Println("0 - Ошибку, 1 - Предупреждение, 2 - Сообщение")
	fmt.Scanln(&input)
	log := logger.Info

	log.Date = time.Now().Format("2006-01-02")
	log.Time = time.Now().Format("15:04:05")
	fmt.Println("Введите ваше сообщение:")
	var mes string
	fmt.Scanln(&mes)
	log.Info = mes
	fmt.Println(logger_this)
	fmt.Println(log)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
