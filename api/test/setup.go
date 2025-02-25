package test

import "github.com/joho/godotenv"

func Init() {
	err := godotenv.Load("../.env.test")

	if err != nil {
		panic(err.Error())
	}
}
