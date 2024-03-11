package main

import "belajar-gin/routers"

var PORT = ":8080"

func main() {
	routers.StartServer().Run(PORT)
}
