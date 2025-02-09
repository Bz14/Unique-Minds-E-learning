package Controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	domain "unique-minds/Domain"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
    courseUsecase domain.CourseUseCaseInterface
}

// NewCourseController creates a new instance of CourseController with the provided CourseUseCaseInterface.
// Parameters:
//   - u: An implementation of the CourseUseCaseInterface which will be used by the CourseController.
// Returns:
//   - A pointer to a newly created CourseController instance.
func NewCourseController(u domain.CourseUseCaseInterface) *CourseController {
    return &CourseController{courseUsecase: u}
}


// UploadFile handles the uploading of a file through a multipart form.
// It expects a file to be uploaded with the key "file".
// If no file is uploaded, it responds with a 400 Bad Request status and an error message.
// If the file is successfully uploaded, it saves the file to the "./uploads/" directory
// and responds with a 200 OK status and the URL of the uploaded file.
// If there is an error saving the file, it responds with a 500 Internal Server Error status and an error message.
func (c *CourseController) UploadFile(ctx *gin.Context) {
    file, err := ctx.FormFile("file")
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
        return
    }

    filename := filepath.Base(file.Filename)
    savePath := fmt.Sprintf("./uploads/%s", filename)

    if err := ctx.SaveUploadedFile(file, savePath); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
        return
    }

    fileURL := fmt.Sprintf("http://localhost:8080/uploads/%s", filename)

    ctx.JSON(http.StatusOK, gin.H{"fileUrl": fileURL})
}


// UploadCourse handles the uploading of a course by a user.
// It expects a JSON payload representing the course details and a user ID in the context.
// If the user ID is missing, it returns a 401 Unauthorized response.
// If the JSON payload is invalid, it returns a 400 Bad Request response.
// If there is an error during the course upload process, it returns a 500 Internal Server Error response.
// On success, it returns a 200 OK response with a success message.
func (c *CourseController) UploadCourse(ctx *gin.Context) {
    var course domain.Course
    user_id := ctx.GetString("user_id")
    if user_id == ""{
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    if err := ctx.ShouldBindJSON(&course); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
        return
    }

    if err := c.courseUsecase.UploadCourse(&course, user_id); err != nil {
        ctx.JSON(http.StatusInternalServerError,
            domain.ErrorResponse{
                Message: "Error uploading course",
                Data:    err.Error(),
                Status:  http.StatusInternalServerError,
            })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Course uploaded successfully"})
}

// GetFeaturedCourses retrieves a list of featured courses.
// It calls the GetRecentCourses method from the course use case to get the data.
// If an error occurs during the data retrieval, it responds with a 500 status code and the error message.
// Otherwise, it responds with a 200 status code and the list of courses.
func (c *CourseController) GetFeaturedCourses(ctx *gin.Context) {
    courses, err := c.courseUsecase.GetRecentCourses()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, courses)
}

// GetCourses handles the HTTP request to retrieve a list of courses.
// It supports pagination and filtering by search term and tag.
// 
// Query Parameters:
// - page: The page number to retrieve (default is "1").
// - limit: The number of items per page (default is "10").
// - search: A search term to filter courses (optional).
// - tag: A tag to filter courses (optional).
//
// Responses:
// - 200 OK: Returns a list of courses along with pagination details.
// - 500 Internal Server Error: Returns an error message if there is an issue retrieving the courses.
//
// Example:
// GET /courses?page=1&limit=10&search=math&tag=science
func (c *CourseController) GetCourses(ctx *gin.Context) {
	pageNo := ctx.Query("page")
	pageSize := ctx.Query("limit")
    search := ctx.Query("search")
    tag := ctx.Query("tag")

	if pageNo == "" {
		pageNo = "1"
	}
	if pageSize == "" {
		pageSize = "10"
	}

	courses, pagination, err := c.courseUsecase.GetCourses(pageNo, pageSize, search, tag)
	if err != nil {
		ctx.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status: 500,

		})
		ctx.Abort()
	} else {
		ctx.JSON(http.StatusOK, domain.SuccessResponse{
			Status: http.StatusOK,
			Data: map[string]interface{}{
				"course" : courses, 
				"pagination" : pagination,
			},
			Message: "courses fetched successfully.",
		})
	}
}

// GetMyCourse handles the HTTP request to retrieve the courses associated with the authenticated user.
// It expects a user ID to be present in the context, which is typically set by an authentication middleware.
// If the user ID is not found, it responds with a 400 status code and an "Unauthorized" error message.
// If the user ID is found, it calls the course use case to fetch the courses for the user.
// If an error occurs while fetching the courses, it responds with a 500 status code and the error message.
// On success, it responds with a 200 status code and the list of courses.
func (c *CourseController) GetMyCourse(ctx *gin.Context) {
    user_id := ctx.GetString("user_id")
    if user_id == ""{
        ctx.JSON(400, gin.H{
            "Error" : "Unauthorized",
        })
    }
    course, err := c.courseUsecase.GetMyCourses(user_id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, course)
}

// GetEducatorCourses handles the request to retrieve courses for a specific educator.
// It expects a user_id to be present in the context, which identifies the educator.
// If the user_id is missing, it responds with an Unauthorized status.
// If the user_id is present, it fetches the courses associated with the educator.
// In case of an error during the fetch operation, it responds with an Internal Server Error status.
// On success, it responds with the list of courses and an OK status.
func (uc *CourseController) GetEducatorCourses(c *gin.Context) {
    user_id := c.GetString("user_id")
    if user_id == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    courses, err := uc.courseUsecase.GetEducatorCourses(user_id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch courses"})
        return
    }
    c.JSON(http.StatusOK, courses)
}

// DeleteCourse handles the HTTP DELETE request to remove a course by its ID.
// It expects the course ID to be provided as a URL parameter and the user ID to be present in the context.
// If the user ID is not found, it responds with an HTTP 401 Unauthorized status.
// If the course deletion fails, it responds with an HTTP 500 Internal Server Error status and the error message.
// On successful deletion, it responds with an HTTP 200 OK status and a success message.
func (uc *CourseController) DeleteCourse(c *gin.Context) {
    courseID := c.Param("id")
	user_id := c.GetString("user_id")
    if user_id == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    err := uc.courseUsecase.DeleteCourse(courseID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}


// GetCourseById handles the HTTP GET request to retrieve a course by its ID.
// It expects the course ID to be provided as a URL parameter and the user ID to be available in the context.
// If the user ID is not present, it responds with an HTTP 401 Unauthorized status.
// If the course retrieval fails, it responds with an HTTP 500 Internal Server Error status and the error message.
// On success, it responds with an HTTP 200 OK status and the course data in JSON format.
func (c *CourseController) GetCourseById(ctx *gin.Context) {
    id := ctx.Param("id")
    user_id := ctx.GetString("user_id")
    if user_id == ""{
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    course, err := c.courseUsecase.GetCourseByID(id, user_id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, course)
}

// UpdateProgress handles the HTTP request to update the progress of a course for a user.
// It expects a JSON body containing the progress data and requires the course ID as a URL parameter
// and the user ID to be set in the context.
func (c *CourseController) UpdateProgress(ctx *gin.Context) {
    var progressRequest []string
    if err := ctx.ShouldBindJSON(&progressRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    courseID := ctx.Param("id")
    userID := ctx.GetString("user_id")

    progress, err := c.courseUsecase.UpdateProgress(courseID, userID, progressRequest)
      
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, progress)
}

// SaveCourse handles the request to save a course for a student.
// It retrieves the student ID from the context and the course ID from the URL parameter.
// If the student ID is not present, it responds with an unauthorized error.
// Otherwise, it calls the SaveCourse method of the course use case to save the course.
// If an error occurs during this process, it responds with an internal server error.
// On success, it responds with a message indicating that the course ID was appended successfully.
//
// @param ctx *gin.Context - the context of the request
func (c *CourseController) SaveCourse(ctx *gin.Context) {
	studentID := ctx.GetString("user_id")
	courseID := ctx.Param("id")

    if studentID == ""{
        ctx.JSON(500, domain.ErrorResponse{
            Message: "Unauthorized",
            Status: 500,
        })
    }
	err := c.courseUsecase.SaveCourse(studentID, courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course ID appended successfully"})
}
