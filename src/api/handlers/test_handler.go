package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type header struct {
	UserId  string
	Browser string
}

type PersonData struct {
	Firstname string
	Lastname  string
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "test working",
	})
}

func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Users List working",
	})
}

func (h *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "User by id working",
		"id":     id,
	})
}

func (h *TestHandler) UserByUsername(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"result":   "User by username working",
		"username": username,
	})
}

func (h *TestHandler) Accounts(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "Account List",
		"id":     id,
	})
}

func (h *TestHandler) AddUser(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"result": "Add User working",
	})
}

func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("userId")
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	})
}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	c.BindHeader(&header)

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"header": header,
	})
}

func (h *TestHandler) QuertBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("id")

	c.JSON(http.StatusOK, gin.H{
		"result": "QuertBinder1",
		"id":     id,
		"name":   name,
	})
}

func (h *TestHandler) QuertBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("id")

	c.JSON(http.StatusOK, gin.H{
		"result": "QuertBinder1",
		"ids":    ids,
		"name":   name,
	})
}

func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	})
}

func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := PersonData{}
	c.ShouldBindJSON(&p)

	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"person": p,
	})
}

func (h *TestHandler) FormBinder(c *gin.Context) {
	p := PersonData{}
	c.ShouldBind(&p)

	c.JSON(http.StatusOK, gin.H{
		"result": "FormBinder",
		"data":   p,
	})
}

func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	})
}
