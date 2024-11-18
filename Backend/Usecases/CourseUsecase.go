package usecases

import (
	"errors"
	"strconv"
	domain "unique-minds/Domain"
)

type courseUseCase struct {
    courseRepo domain.CourseRepository
}

func NewCourseUseCase(cr domain.CourseRepository) *courseUseCase  {
    return &courseUseCase {
        courseRepo: cr,
    }
}

func (u *courseUseCase) UploadCourse(course *domain.Course, user_id string) error {
    return u.courseRepo.Save(course, user_id)
}

func (u *courseUseCase) GetRecentCourses() ([]domain.Course, error) {
    result, err := u.courseRepo.FetchRecentCourses()
    if err != nil{
        return nil, err
    }
    return result, nil
}

func (u *courseUseCase) GetCourses(pageNo string, pageSize string, search string, filter string) ([]domain.Course, domain.Pagination, error) {
	PageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return []domain.Course{}, domain.Pagination{}, err
	}
	PageSize, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		return []domain.Course{}, domain.Pagination{}, err
	}
	if PageNo <= 0 || PageSize <= 0 {
		return []domain.Course{}, domain.Pagination{}, errors.New("invalid page number or page size")
	}

	blogs, pagination, err := u.courseRepo.GetCourses(PageNo, PageSize, search, filter)
	if err != nil {
		return nil, domain.Pagination{}, err
	} else {
		return blogs, pagination, nil
	}
}

func (u *courseUseCase) GetMyCourses(id string) ([]domain.Course, error) {
	result, err := u.courseRepo.GetMyCourse(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *courseUseCase) GetEducatorCourses(id string) ([]domain.Course, error) {
	result, err := u.courseRepo.GetCoursesByEducator(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *courseUseCase) DeleteCourse(id string) error {
	err := u.courseRepo.DeleteCourse(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *courseUseCase) GetCourseByID(courseID, userID string) (*domain.CourseDetailResponse, error) {
    course, err := u.courseRepo.GerCourseById(courseID)
    if err != nil {
        return nil, err
    }
    progress, err := u.courseRepo.GetCourseProgress(courseID, userID)
	var courseDetail  domain.CourseDetailResponse
	courseDetail.Course = course
    if err == nil {
		courseDetail.Progress.Progress = progress.Progress
        courseDetail.Progress.CompletedParts = progress.CompletedParts
        courseDetail.Progress.IsCompleted = progress.IsCompleted
    }
    return &courseDetail, nil
}

func  (u *courseUseCase) UpdateProgress(courseID, userID string, completedParts []string) (domain.CourseProgress, error) {
    return u.courseRepo.UpdateCourseProgress(courseID, userID, completedParts)
}


func  (u *courseUseCase) SaveCourse(courseID, userID string) error {
    return u.courseRepo.SaveCourse(courseID, userID)
}
