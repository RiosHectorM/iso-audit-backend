-- name: GetAuditsByUserID :many
SELECT a.* FROM audits a
JOIN assignments asg ON a.id = asg.audit_id
WHERE asg.user_id = $1;

-- name: CreateAudit :one
INSERT INTO audits (id, title, norm, status, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: AssignUserToAudit :exec
INSERT INTO assignments (user_id, audit_id, sector_id)
VALUES ($1, $2, $3);

-- name: GetAuditByID :one
SELECT * FROM audits
WHERE id = $1;