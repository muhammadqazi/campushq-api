package validators

import (
	"context"
	"fmt"
	"strconv"
)

func (v *studentValidators) StudentIdSchema(id string) error {
	ctx := context.TODO()

	studentId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid student ID")
	}

	_, err = v.repository.SelectStudentById(ctx, int32(studentId))
	if err != nil {
		return fmt.Errorf("student with ID %s does not exist", id)
	}

	return nil
}
