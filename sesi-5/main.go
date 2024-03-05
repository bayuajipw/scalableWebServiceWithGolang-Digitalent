package main

import (
	"errors"
	"fmt"
)

// // ================= CHANNEL =================
// func main() {
// 	channel := make(chan string)

// 	go introduce("Kurai", channel)
// 	go introduce("Ilham", channel)
// 	go introduce("Samsul", channel)

// msg1 := <-channel
// fmt.Println(msg1)
// msg2 := <-channel
// fmt.Println(msg2)
// msg3 := <-channel
// fmt.Println(msg3)
// msg4 := <-channel
// fmt.Println(msg4) //nilai error karena data goroutine cuma 3

// 	// atau bisa seperti ini
// 	for datum := range channel {
// 		fmt.Println(datum)
// 	}

// 	close(channel)
// }

// func introduce(name string, channel chan string) {
// 	result := fmt.Sprintf("Hi, my name is %s\n", name)

// 	channel <- result
// }

// // ================= CHANNELS (Channel with anonymous function) =================
// func main() {
// 	channel := make(chan string)

// 	students := []string{"Ilham", "Ramli", "Ismail"}

// 	for _, student := range students {
// 		go func(stud string) {
// 			fmt.Println(stud)
// 			result := fmt.Sprintf("Hi, my name is %s", stud)
// 			channel <- result
// 		}(student)
// 	}

// 	for i := 0; i < 3; i++ {
// 		print(channel)
// 	}

// 	close(channel)
// }

// func print(channel chan string) {
// 	fmt.Println(<-channel)
// }

// // ================= CHANNELS (Channel directions) =================
// func main() {
// 	c := make(chan string)

// 	students := []string{"Marsha", "Toya", "Zee"}

// 	for _, student := range students {
// 		go introduce(student, c)
// 	}

// 	for i := 0; i < 3; i++ {
// 		print(c)
// 	}

// 	close(c)
// }

// func print(c chan string) {
// 	fmt.Println(<-c)
// }

// func introduce(student string, c chan string) {
// 	result := fmt.Sprintf("Hi, my name is %s", student)
// 	c <- result
// }

// // ================= CHANNELS (Unbuffered Channel) =================
// func main() {
// 	c := make(chan int)

// 	go func(chan int) {
// 		fmt.Println("func goroutine starts sending data into this channel")
// 		c <- 10
// 		fmt.Println("func goroutine after sending data into this channel")
// 	}(c)

// 	fmt.Println("main func sleep for 1 sec")
// 	time.Sleep(1 * time.Second)

// 	fmt.Println("main func start receiving data")
// 	num := <-c
// 	fmt.Println("main func received data: ", num)

// 	close(c)
// 	time.Sleep(500 * time.Millisecond)
// }

// // ================= CHANNELS (Buffered Channel) =================
// func main() {
// 	c := make(chan int, 3)

// 	go func(c chan int) {
// 		for i := 1; i <= 5; i++ {
// 			fmt.Println("func goroutine starts sending data into this channel", i)
// 			c <- i
// 			fmt.Println("func goroutine finished sending data into this channel", i)
// 		}

// 		close(c)
// 	}(c)

// 	fmt.Println("main func sleeps for 1 sec")
// 	time.Sleep(1 * time.Second)

// 	for v := range c {
// 		fmt.Println("main func received value from channel:", v)
// 	}
// }

// // ================= CHANNELS (Select) =================
// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		c1 <- "Hi!"
// 	}()

// 	go func() {
// 		time.Sleep(1500 * time.Millisecond)
// 		c2 <- "Salute!"
// 	}()

// 	for i := 1; i <= 2; i++ {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println("Received:", msg1)
// 		case msg2 := <-c2:
// 			fmt.Println("Received", msg2)
// 		}
// 	}
// }

// // ================= DEFER & EXIT =================
// func main() {
// 	defer fmt.Println("Hi my name is Marsha")
// 	fmt.Println("I'm 18 years old")
// 	fmt.Println("I'm from Jakarta")
// }

// func main() {
// 	callDeferFunc()
// 	fmt.Println("Hi, everyone!")
// }

// func callDeferFunc() {
// 	defer deferFunc()
// }

// func deferFunc() {
// 	fmt.Println("Defer func start to execute")
// }

// func main() {
// 	defer deferFunc()
// 	fmt.Println("Hi, everyone!")
// 	os.Exit(1)
// }

// func deferFunc() {
// 	fmt.Println("Defer func start to execute")
// }

// // ================= ERROR =================
// func main() {
// 	var number int
// 	var err error

// 	number, err = strconv.Atoi("123a")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(number)
// 	}

// 	number, err = strconv.Atoi("123")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(number)
// 	}

// }

// // ================= ERROR (Custom Error) =================
// func main() {
// 	var password string

// 	fmt.Scanln(&password)

// 	_, err := validPassword(password)
// 	if err != nil {
// 		fmt.Println("Error from validPassword:", err)
// 	} else {
// 		fmt.Println("Password is valid")
// 	}
// }

// func validPassword(password string) (result bool, err error) {
// 	passwordLen := len(password)

// 	if passwordLen < 5 {
// 		return false, errors.New("password should be more than 4 characters")
// 	}

// 	return true, nil
// }

// // ================= PANIC & RECOVER =================
func main() {
	defer catchError()
	var password string
	fmt.Scanln(&password)

	_, err := validPassword(password)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Password is valid")
	}
}

func validPassword(password string) (bool, error) {
	passLen := len(password)

	if passLen < 5 {
		return false, errors.New("password should be more than 4")
	}

	return true, nil
}

func catchError() {
	if r := recover(); r != nil {
		fmt.Println("Error accourde:", r)
	} else {
		fmt.Println("App is running good.")
	}
}
