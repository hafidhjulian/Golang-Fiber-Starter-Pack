package main

import (
	"fmt"
	"golang-fiber-starterpack/config"
	"golang-fiber-starterpack/param"
	"golang-fiber-starterpack/routers"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/unrolled/secure"
)

var AllowedIP []string = []string{
	// "IP YANG AKAN DI ALLOW",

}

func main() {
	config := config.LoadConfig(".")
	app := fiber.New(fiber.Config{
		Immutable: true,
		BodyLimit: 10 * 1024 * 1024,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(param.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:            AllowedIP,
		FrameDeny:               true,
		CustomBrowserXssValue:   "1; mode=block",
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXssFilter:        true,
		ContentSecurityPolicy:   "script-src $NONCE",
		STSSeconds:              31536000,
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},
		STSIncludeSubdomains:    true,
		STSPreload:              true,
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"},
	})

	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
		Done: func(c *fiber.Ctx, logString []byte) {
			logdata := `{"time":"` + time.Now().Add(time.Hour*6).Format(time.RFC3339) + `","remote_ip":"` + c.IP() + `","host":"` + c.Hostname() + `","user_agent":"` + string(c.Request().Header.UserAgent()) + `",` +
				`"method":"` + c.Method() + `","uri":"` + c.BaseURL() + `","status":"` + strconv.Itoa((c.Response().StatusCode())) + `",` +
				`"path":"` + string(c.Request().RequestURI()) + `", "referer":"` + string(c.Request().Header.Referer()) + `"}`
			fmt.Println(logdata)
		},
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	app.Use(adaptor.HTTPMiddleware(secureMiddleware.Handler))

	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})

	// routers
	routers.ExampleRoute(app)

	// Start Engine
	host := fmt.Sprintf(":%d", config.ServerPort)
	log.Fatal(app.Listen(host))
}
