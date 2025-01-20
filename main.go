package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AyanDas-99/blog/db"
	"github.com/AyanDas-99/blog/post"
	"github.com/AyanDas-99/blog/user"
	"github.com/gin-gonic/gin"
)

type Blog struct {
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
	Metadata  Meta   `json:"metadata"`
}

type Meta struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Length int    `json:"length"`
}

func NewBlog(id string, author string, content string, timestamp string) Blog {
	return Blog{content, timestamp, Meta{id, author, len(content)}}
}

func blogList(count int) []Blog {
	blogs := []Blog{}

	for i := range count {
		blog := NewBlog(strconv.Itoa(i), "ayan", "This is the content", "2024-01-01 15:05:05")
		blogs = append(blogs, blog)
	}
	return blogs
}

var DB sql.DB

func main() {
	router := gin.Default()
	router.GET("/user", getUsers)
	router.GET("/user/:id", getUserById)
	router.PUT("/user", addUser)
	router.DELETE("/user/:id", deleteUser)

	router.GET("/post", getPosts)
	router.GET("/post/:id", getPostById)
	router.PUT("/post", addPost)
	router.DELETE("/post/:id", deletePost)
	db.InitDb()
	router.Run("localhost:8080")

}
func getUserById(c *gin.Context) {
	id := c.Param("id")
	res, err := user.GetUserById(id)
	if err != nil {
		fmt.Println("Error fetching user of id", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not query results"})
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}
func getUsers(c *gin.Context) {
	res, err := user.GetUsers()
	if err != nil {
		fmt.Println("Error fetching users:", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not query results"})
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}

func addUser(c *gin.Context) {
	var reqUser user.User

	if er := c.BindJSON(&reqUser); er != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "could not resolve body"})
		return
	}

	res, err := user.PutUsers(reqUser)
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Could not add user"})
		return
	}
	fmt.Println(res)

	c.IndentedJSON(http.StatusOK, res)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	res, er := user.DeleteUser(id)

	if er != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not delete id"})
		return
	}

	fmt.Println(res)
	c.IndentedJSON(http.StatusOK, res)
}

func getPostById(c *gin.Context) {
	id := c.Param("id")
	res, err := post.GetPostById(id)
	if err != nil {
		fmt.Println("Error fetching post of id", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not query results"})
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}
func getPosts(c *gin.Context) {
	res, err := post.GetPosts()
	if err != nil {
		fmt.Println("Error fetching posts:", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not query results"})
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}

func addPost(c *gin.Context) {
	var reqPost post.Post

	if er := c.BindJSON(&reqPost); er != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "could not resolve body"})
		return
	}

	res, err := post.PutPost(reqPost)
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Could not add user"})
		return
	}
	fmt.Println(res)

	c.IndentedJSON(http.StatusOK, res)
}

func deletePost(c *gin.Context) {
	id := c.Param("id")
	res, er := post.DeletePost(id)

	if er != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not delete id"})
		return
	}

	fmt.Println(res)
	c.IndentedJSON(http.StatusOK, res)
}
