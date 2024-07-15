-- name: GetQuickCodesByType :many
SELECT quick_code_id, type, name, code, description, created_at, updated_at
FROM bk_quick_codes
WHERE type = $1;


-- name: GetAllStates :many
SELECT DISTINCT name, code, description
FROM bk_quick_codes
WHERE type = 'STATE';

-- name: GetAllDomains :many
SELECT DISTINCT name, code, description
FROM bk_quick_codes
WHERE type = 'DOMAIN';
