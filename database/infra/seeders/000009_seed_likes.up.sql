WITH random_likes AS (
    SELECT DISTINCT ON (u.id, p.id)
        u.id AS user_id,
        p.id AS post_id,
        now() - (random() * interval '30 days') AS created_at,
        now() AS updated_at
    FROM users u
             CROSS JOIN posts p
    ORDER BY u.id, p.id, random()
    LIMIT 100
)

INSERT INTO likes (user_id, post_id, created_at, updated_at)
SELECT user_id, post_id, created_at, updated_at
FROM random_likes;
