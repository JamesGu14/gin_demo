package controllers

import (
	"gin_demo/dao"
	"gin_demo/models"
	"gin_demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Load Student Page
func StudentPage(c *gin.Context) {
	pageIdx, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect page number"})
		pageIdx = 0
	}

	students := services.FindStudents(models.Paging{Page: pageIdx})
	c.HTML(http.StatusOK, "student/list.tmpl", gin.H{
		"title":    "Dashboard",
		"students": students,
	})
}

func StudentCreatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "student/create_or_edit.tmpl", gin.H{
		"title": "Create Student",
		"page":  "create",
	})
}

// GET /student
// Find all students
func FindStudents(c *gin.Context) {
	pageIdx, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect page number"})
		return
	}
	students := services.FindStudents(models.Paging{Page: pageIdx})
	c.JSON(http.StatusOK, gin.H{"data": students})
}

// Get /student/:id
// Find a student by id
func FindStudent(c *gin.Context) {
	student, err := services.FindStudent(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}

// POST /student
// Add a new student
func CreateStudent(c *gin.Context) {
	var input dao.CreateStudentIntput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newStudent := dao.Student{
		UserName:  input.UserName,
		Password:  input.Password,
		FullName:  input.FullName,
		Email:     input.Email,
		Birthday:  input.Birthday,
		ClassID:   input.ClassID,
		IsMonitor: input.IsMonitor,
	}
	services.CreateStudent(newStudent)
	c.JSON(http.StatusOK, gin.H{"message": "succeed"})
}

// PUT /student/:id
func UpdateStudent(c *gin.Context) {
	// Get existing student
	existingStudent, err := services.FindStudent(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input dao.UpdateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.UpdateStudent(existingStudent, input)
	c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "updated"})
}

// DELETE /student/:id
func DeleteStudent(c *gin.Context) {
	student_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student_id"})
		return
	}
	services.DeleteStudent(student_id)
	c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "succeed"})
}
