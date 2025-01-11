package router

import (
	fiber "github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	controllers "github.com/herumitra/apt-auth/controllers"
	middleware "github.com/herumitra/apt-auth/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Adding logger middleware of fiber
	app.Use(logger.New())

	// Auth Endpoints
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
	app.Get("/profile", middleware.JWTMiddleware, controllers.GetProfile)

	// SetBranch Endpoint
	app.Post("/set_branch", controllers.SetBranch)

	// Grouped API routes
	apiAdm := app.Group("/api", middleware.JWTMiddleware, middleware.RoleMiddleware("administrator"))
	apiAdmOp := app.Group("/api", middleware.JWTMiddleware, middleware.RoleMiddleware("administrator", "operator"))
	apiAdmOpCsFn := app.Group("/api", middleware.JWTMiddleware, middleware.RoleMiddleware("administrator", "operator", "cashier", "finance"))

	// User Endpoints
	apiAdm.Post("/users", controllers.CreateUser)
	apiAdm.Get("/users", controllers.GetAllUsers)
	apiAdm.Get("/users/:id", controllers.GetUser)
	apiAdm.Put("/users/:id", controllers.UpdateUser)
	apiAdm.Delete("/users/:id", controllers.DeleteUser)

	// Branch Endpoints
	apiAdmOp.Post("/branches", controllers.CreateBranch)
	apiAdmOp.Get("/branches", controllers.GetAllBranch)
	apiAdmOp.Get("/branches/:id", controllers.GetBranch)
	apiAdmOp.Put("/branches/:id", controllers.UpdateBranch)
	apiAdm.Delete("/branches/:id", controllers.DeleteBranch)

	// UserBranch Endpoints
	apiAdm.Post("/user_branches", controllers.CreateUserBranch)
	apiAdmOpCsFn.Get("/user_branches", controllers.GetAllUserBranch)
	apiAdmOpCsFn.Get("/user_branches/:userid", controllers.GetUserBranch)
	apiAdm.Put("/user_branches/:userid/:branchid", controllers.UpdateUserBranch)
	apiAdm.Delete("/user_branches/:userid/:branchid", controllers.DeleteUserBranch)

}
