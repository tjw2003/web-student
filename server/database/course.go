package database

import "web-student/server/models"

func InsertCourse(course models.Course) (int64, error) {
	return DB.Insert("lessons(name,credit,teacher)", course.Name, course.Credit, course.Teacher)
}

func InsertSelectCourse(userId int, course_id int) (int64, error) {
	return DB.Insert("select_courses(user_id, course_id)", userId, course_id)
}

func GetCourses() ([]models.Course, error) {
	courses := make([]models.Course, 0)

	rows, err := DB.Query("courses")
	if err != nil {
		return courses, nil
	}

	for rows.Next() {
		course := models.Course{}
		if err := rows.Scan(&course.ID, &course.Name, &course.Credit, &course.Teacher); err != nil {
			return courses, nil
		}
		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}

func GetSelectedCourses(id int) ([]models.Course, error) {
	courses := make([]models.Course, 0)

	rows, err := DB.Query("select_courses", "id in (SELECT course_id FROM select_courses WHERE user_id = $1)", id)
	if err != nil {
		return nil, nil
	}

	for rows.Next() {
		course := models.Course{}
		if err := rows.Scan(&course.ID, &course.Name, &course.Credit, &course.Teacher); err != nil {
			return courses, nil
		}
		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return courses, err
	}

	return courses, nil
}