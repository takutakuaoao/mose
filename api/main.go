package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/health-checker", func(c *gin.Context) {
		err := godotenv.Load()

		if err != nil {
			panic(err.Error())
		}

		db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		defer db.Close()

		//構造体
		type Sample struct {
			ID        int    `json:"id"`
			Title     string `json:"title"`
			CreatedAt string `json:"created_at"`
		}
		var sample Sample

		err = db.QueryRow("SELECT * FROM samples").Scan(&sample.ID, &sample.Title, &sample.CreatedAt)

		if err != nil {
			panic(err.Error())
		}

		//jsonを生成
		res, err := json.Marshal(sample)
		if err != nil {
			panic(err.Error())
			return
		}

		c.JSON(200, gin.H{
			"result": "success",
			"sample": string(res),
		})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}
