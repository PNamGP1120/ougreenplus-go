package middleware

import "github.com/gofiber/fiber/v2/middleware/cors"

var CORS = cors.New(cors.Config{
	AllowOrigins: "*",
	AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	AllowHeaders: "*",
})
