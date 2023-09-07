-- name: SelectAllBuildings :many
SELECT * FROM buildings WHERE is_active = true AND deleted_at IS NULL;