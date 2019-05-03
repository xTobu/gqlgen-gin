
package main

import (
	"fmt"

	"github.com/99designs/gqlgen/handler"
	"github.com/xTobu/gqlgen-todos"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

// Handler initializes the graphql middleware.
func Handler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.GraphQL(gqlgen_todos.NewExecutableSchema(gqlgen_todos.Config{Resolvers: &gqlgen_todos.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// HandlerPlayground initializes the graphql middleware.
func HandlerPlayground() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.Playground("GraphQL playground", "/gql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// GinContextToContextMiddleware
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	r := gin.New()
		r.GET("/gql", HandlerPlayground())
		r.POST("/gql", Handler())
		r.Use(GinContextToContextMiddleware())
	fmt.Println("Now server is running on port 9999")
	r.Run(":9999")
}