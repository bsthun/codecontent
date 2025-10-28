package endpoint

import (
	"backend/common/config"
	"backend/common/fiber/middleware"
	"backend/endpoint/course"
	"backend/endpoint/public"
	"backend/endpoint/state"
	"backend/type/common"
	"mime"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func Bind(
	app *fiber.App,
	config *config.Config,
	frontend common.FrontendFS,
	publicEndpoint *publicEndpoint.Handler,
	stateEndpoint *stateEndpoint.Handler,
	courseEndpoint *courseEndpoint.Handler,
	middleware *middleware.Middleware,
) {
	api := app.Group("/api")
	api.Use(middleware.Id())

	// * public endpoints
	public := api.Group("/public")
	public.Get("/login/redirect", publicEndpoint.HandleLoginRedirect)
	public.Post("/login/callback", publicEndpoint.HandleLoginCallback)

	// * state endpoints
	state := api.Group("/state", middleware.Jwt(true))
	state.Post("/state", stateEndpoint.HandleState)

	// * course endpoints
	courses := api.Group("/courses", middleware.Jwt(true))
	courses.Post("/create", courseEndpoint.HandleCourseCreate)
	courses.Post("/edit", courseEndpoint.HandleCourseEdit)
	courses.Post("/delete", courseEndpoint.HandleCourseDelete)
	courses.Post("/list/manager", courseEndpoint.HandleCourseListManager)

	// * frontend
	app.Get("*", func(c *fiber.Ctx) error {
		filePath := filepath.Join(".local/dist", c.Path())
		file, err := frontend.ReadFile(filePath)
		if err != nil {
			file, _ = frontend.ReadFile(".local/dist/index.html")
			c.Set("Content-Type", "text/html")
			return c.Send(file)
		}

		contentType := mime.TypeByExtension(filepath.Ext(filePath))
		c.Set("Content-Type", contentType)

		return c.Send(file)
	})
}
