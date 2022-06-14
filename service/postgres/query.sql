-- name: ListRegistration :many
SELECT * FROM registration
ORDER BY number_plate;

-- name: CreateRegistration :one
INSERT INTO registration (number_plate, vehicle, insurer)
VALUES ($1, $2, $3)
RETURNING *;


