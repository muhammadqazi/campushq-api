// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: building-queries.sql

package repositories

import (
	"context"
)

const selectAllBuildings = `-- name: SelectAllBuildings :many
SELECT building_id, name, code, description, number_of_rooms, created_at, updated_at, deleted_at, is_active FROM buildings WHERE is_active = true AND deleted_at IS NULL
`

func (q *Queries) SelectAllBuildings(ctx context.Context) ([]Building, error) {
	rows, err := q.db.Query(ctx, selectAllBuildings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Building{}
	for rows.Next() {
		var i Building
		if err := rows.Scan(
			&i.BuildingID,
			&i.Name,
			&i.Code,
			&i.Description,
			&i.NumberOfRooms,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.IsActive,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
