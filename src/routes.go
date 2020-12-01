// routes.go

package main

func initializeRoutes(basePath string) {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handle the index route
	router.GET(basePath + "/", showIndexPage)

	// Group user related routes together
	userRoutes := router.Group(basePath + "/u")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", ensureLoggedIn(), logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	// Group api relates routes together
	apiRoutes := router.Group(basePath + "/api")
	{
		// Handle GET requests at /api/getCompassURLByNameAndVersion/some_compass_name/some_compass_version
		apiRoutes.GET("/getCompassURLByNameAndVersion/:name/:version", getCompassURLByNameAndVersion)

		// Handle GET requests at /api/getCompassObjectsByName/some_compass_name
		apiRoutes.GET("/getCompassObjectsByName/:name", getCompassObjectsByName)

		// Handle GET requests at /api/getCompassObjectsByFuzzyName/some_compass_name
		apiRoutes.GET("/getCompassObjectsByFuzzyName/:name", getCompassObjectsByFuzzyName)
	}

	uiRoutes := router.Group(basePath + "/ui")
	{
		uiRoutes.GET("/logs", showLogsPage)
	}
}
