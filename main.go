package main

import (
	"log"
	"onez19/config" // เปลี่ยนเป็นชื่อโมดูลของคุณ
	"onez19/routes" // เปลี่ยนเป็นชื่อโมดูลของคุณ
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// โหลดไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// สร้างแอพ Fiber
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // ระบุ origin ที่อนุญาต
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// เชื่อมต่อฐานข้อมูล
	config.ConnectDatabase()

	// กำหนดเส้นทาง
	routes.SetupRoutes(app)

	// เริ่มเซิร์ฟเวอร์
	port := ":" + os.Getenv("PORT")
	log.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(app.Listen(port))
}
