// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: departmentQueries.sql

package repositories

import (
	"context"
)

const getDepartmentById = `-- name: GetDepartmentById :one
SELECT department_id, name, code, description, number_of_semesters, department_email, department_phone_number, tution_per_semester, created_at, updated_at, deleted_at, head_of_department_id, faculty_id FROM departments WHERE department_id = $1
`

func (q *Queries) GetDepartmentById(ctx context.Context, departmentID int32) (Department, error) {
	row := q.db.QueryRow(ctx, getDepartmentById, departmentID)
	var i Department
	err := row.Scan(
		&i.DepartmentID,
		&i.Name,
		&i.Code,
		&i.Description,
		&i.NumberOfSemesters,
		&i.DepartmentEmail,
		&i.DepartmentPhoneNumber,
		&i.TutionPerSemester,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.HeadOfDepartmentID,
		&i.FacultyID,
	)
	return i, err
}
