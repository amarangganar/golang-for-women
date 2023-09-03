package controllers

import (
	"assignmentproject/database"
	"assignmentproject/helpers"
	"assignmentproject/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(ctx *gin.Context) {
	db := database.Get()

	users := []models.Student{}
	if fetchErr := db.Preload("Scores").Find(&users).Error; fetchErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusBadRequest,
			"mesage": "Failed to retrieve students.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Students has been retrieved.",
		"data":    users,
	})
}

func CreateStudent(ctx *gin.Context) {
	student := models.Student{}

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	// validate request body
	if validateErr := helpers.Validate(student); validateErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Field validation error",
			"error":   validateErr,
		})

		return
	}

	db := database.Get()

	// store to database
	if storeErr := db.Debug().Create(&student).Error; storeErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusInternalServerError,
			"message": storeErr,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "New student has been successfully created.",
		"data":    student,
	})
}

func GetStudent(ctx *gin.Context) {
	// convert param studentID to int
	studentID, err := strconv.Atoi(ctx.Param("studentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid student ID.",
		})

		return
	}

	db := database.Get()

	// find student
	student := models.Student{}
	if findErr := db.Debug().Preload("Scores").First(&student, studentID).Error; findErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Student not found.",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"mesage": fmt.Sprintf("Student with ID %d has been retrieved.", studentID),
		"data":   student,
	})
}

func UpdateStudent(ctx *gin.Context) {
	// convert param studentID to int
	studentID, err := strconv.Atoi(ctx.Param("studentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid student ID.",
		})

		return
	}

	db := database.Get()

	// find student
	student := models.Student{}
	if findErr := db.Debug().Preload("Scores").First(&student, studentID).Error; findErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Student not found.",
		})

		return
	}

	// get request body
	body := models.Student{}
	if bodyErr := ctx.ShouldBindJSON(&body); bodyErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": bodyErr.Error(),
		})

		return
	}

	// validate request body
	if validateErr := helpers.Validate(body); validateErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Field validation error",
			"error":   validateErr,
		})

		return
	}

	// update student
	student.Name = body.Name
	student.Age = body.Age

	// only update Student's scores if student.Scores not nil
	if body.Scores != nil {
		student.Scores = body.Scores
	}

	if updateErr := db.Debug().Save(&student).Error; updateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusInternalServerError,
			"message": updateErr,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"mesage": fmt.Sprintf("Student with ID %d data has been updated.", studentID),
		"data":   student,
	})
}

func DeleteStudent(ctx *gin.Context) {
	// convert param studentID to int
	studentID, err := strconv.Atoi(ctx.Param("studentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid student ID.",
		})

		return
	}

	db := database.Get()

	// find to be deleted student
	student := models.Student{}
	if findErr := db.Debug().First(&student, studentID).Error; findErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Student not found.",
		})

		return
	}

	// delete student
	deleteErr := db.Debug().Delete(student).Error
	if deleteErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": deleteErr.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"mesage": fmt.Sprintf("Student with ID %d has been deleted.", studentID),
		"data":   nil,
	})
}
