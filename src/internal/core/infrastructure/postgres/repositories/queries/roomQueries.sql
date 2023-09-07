-- name: SelectAllRooms :many
SELECT * FROM rooms WHERE is_active = true AND deleted_at IS NULL;