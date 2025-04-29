package http

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	// Users endpoints
	// Handles user registration, authentication, profile management, and social features
	users := api.Group("/users")
	{
		// Register a new user account
		users.POST("/register", nil)
		// Authenticate user and return JWT token
		users.POST("/login", nil)
		// Get user's public profile information
		users.GET("/profile", nil)
		// Update user's profile information
		users.PATCH("/profile", nil)
		// Follow another user
		users.POST("/follow/:user_id", nil)
		// Unfollow a user
		users.DELETE("/unfollow/:user_id", nil)
		// Get list of users following the current user
		users.GET("/followers", nil)
		// Get list of users the current user is following
		users.GET("/following", nil)
		// Get current authenticated user's information
		users.GET("/me", nil)
	}

	// Posts (Product Advertisements) endpoints
	// Handles product advertisements, their interactions, and reviews
	// Note: All GET endpoints that return lists (product ads, comments, etc.) support cursor-based pagination
	// Query parameters: ?cursor=xxx&limit=20
	// Product advertisements are filtered based on user's following accounts when authenticated
	// Posts represent product advertisements that sellers create to showcase their products
	posts := api.Group("/posts")
	{
		// Create a new product advertisement
		posts.POST("", nil)
		// Returns product advertisements from followed accounts, paginated
		posts.GET("", nil)
		// Get a specific product advertisement
		posts.GET("/:post_id", nil)
		// Update a product advertisement
		posts.PATCH("/:post_id", nil)
		// Delete a product advertisement
		posts.DELETE("/:post_id", nil)
		// Like a product advertisement
		posts.POST("/:post_id/like", nil)
		// Remove like from a product advertisement
		posts.DELETE("/:post_id/unlike", nil)
		// Record a view of a product advertisement
		posts.POST("/:post_id/view", nil)
		// Add a comment to a product advertisement
		posts.POST("/:post_id/comments", nil)
		// Get comments on a product advertisement, paginated
		posts.GET("/:post_id/comments", nil)
		// Report a product advertisement
		posts.POST("/:post_id/reports", nil)
		// Add a review to a product advertisement
		posts.POST("/:post_id/reviews", nil)
		// Get reviews for a product advertisement, paginated
		posts.GET("/:post_id/reviews", nil)
	}

	// Blogs endpoints
	// Handles blog posts, their interactions, and reviews
	// Note: All GET endpoints that return lists (blogs, comments, etc.) support cursor-based pagination
	// Query parameters: ?cursor=xxx&limit=20
	// Blogs are filtered based on user's following accounts when authenticated
	blogs := api.Group("/blogs")
	{
		// Create a new blog post
		blogs.POST("", nil)
		// Returns blogs from followed accounts, paginated
		blogs.GET("", nil)
		// Get a specific blog post
		blogs.GET("/:blog_id", nil)
		// Update a blog post
		blogs.PATCH("/:blog_id", nil)
		// Delete a blog post
		blogs.DELETE("/:blog_id", nil)
		// Like a blog post
		blogs.POST("/:blog_id/like", nil)
		// Remove like from a blog post
		blogs.DELETE("/:blog_id/unlike", nil)
		// Add a comment to a blog post
		blogs.POST("/:blog_id/comments", nil)
		// Get comments on a blog post, paginated
		blogs.GET("/:blog_id/comments", nil)
		// Report a blog post
		blogs.POST("/:blog_id/reports", nil)
		// Add a review to a blog post
		blogs.POST("/:blog_id/reviews", nil)
		// Get reviews for a blog post, paginated
		blogs.GET("/:blog_id/reviews", nil)
	}

	// Notifications endpoints
	// Handles user notifications and their status
	// Note: GET endpoint supports cursor-based pagination
	// Query parameters: ?cursor=xxx&limit=20
	notifications := api.Group("/notifications")
	{
		// Get user's notifications, paginated
		notifications.GET("", nil)
		// Mark a notification as read
		notifications.PATCH("/:notification_id/read", nil)
	}

	// Reports endpoints
	// Handles content reports and their resolution
	// Note: GET endpoint supports cursor-based pagination
	// Query parameters: ?cursor=xxx&limit=20
	reports := api.Group("/reports")
	{
		// Get list of reports, paginated
		reports.GET("", nil)
		// Mark a report as resolved
		reports.PATCH("/:report_id/resolve", nil)
		// Reject a report
		reports.PATCH("/:report_id/reject", nil)
	}

	// Payments & Subscriptions endpoints
	// Handles payment processing, subscription management, and payment history
	// Note: GET /history endpoint supports cursor-based pagination
	// Query parameters: ?cursor=xxx&limit=20
	payments := api.Group("/payments")
	{
		// Subscribe to a premium plan
		payments.POST("/subscribe", nil)
		// Get payment history, paginated
		payments.GET("/history", nil)
		// Process a payment checkout
		payments.POST("/checkout", nil)
		// Webhook endpoint for payment status updates
		payments.POST("/webhook", nil)
	}

	// Categories endpoints
	// Handles product categories and category-based product listings
	// Note: GET /:category_id/posts endpoint supports cursor-based pagination
	// Query parameters: ?cursor=xxx&limit=20
	categories := api.Group("/categories")
	{
		// Get all available categories
		categories.GET("", nil)
		// Get product advertisements in a specific category, paginated
		categories.GET("/:category_id/posts", nil)
	}

	// Admin endpoints
	// Handles administrative functions and content moderation
	// Note: All GET endpoints that return lists support cursor-based pagination
	// Query parameters: ?cursor=xxx&limit=20
	admin := api.Group("/admin")
	{
		// Get list of seller applications, paginated
		admin.GET("/seller-applications", nil)
		// Approve a seller application
		admin.PATCH("/seller-applications/:application_id/approve", nil)
		// Reject a seller application
		admin.PATCH("/seller-applications/:application_id/reject", nil)
		// Get list of reports, paginated
		admin.GET("/reports", nil)
		// Take action on reported content
		admin.PATCH("/take-action", nil)
		// Get list of users, paginated
		admin.GET("/users", nil)
		// Deactivate a user account
		admin.PATCH("/users/:user_id/deactivate", nil)
		// Take down a product advertisement
		admin.PATCH("/posts/:post_id/take-down", nil)
		// Take down a blog post
		admin.PATCH("/blogs/:blog_id/take-down", nil)
	}

	// Seller endpoints
	// Handles seller-specific product management and analytics
	// Note: All GET endpoints that return lists support cursor-based pagination
	// Query parameters: ?cursor=xxx&limit=20
	seller := api.Group("/seller")
	{
		// Create a new product
		seller.POST("/products", nil)
		// Get list of seller's products, paginated
		seller.GET("/products", nil)
		// Get details of a specific product
		seller.GET("/products/:seller_product_id", nil)
		// Update a product
		seller.PATCH("/products/:seller_product_id", nil)
		// Delete a product
		seller.DELETE("/products/:seller_product_id", nil)
		// Get list of orders, paginated
		seller.GET("/orders", nil)
		// Get seller's analytics and statistics
		seller.GET("/analytics", nil)
		// Get payout history, paginated
		seller.GET("/payouts", nil)
	}

	// Root endpoint
	// Health check and service status
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Ethiopian Market of Digital is running!",
		})
	})

	return router
}
