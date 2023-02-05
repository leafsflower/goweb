package main

import "awesomeProject1/routers"

func main() {
	router := routers.InitRouter()
	router.Static("/assets", "./assets")

	router.Run()
}
