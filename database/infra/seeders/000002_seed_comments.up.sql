-- Step 1: Generate 100 random users, posts, and optionally existing parent comments
WITH random_users AS (
    SELECT id AS user_id, row_number() OVER () AS rn
    FROM users
    ORDER BY random()
    LIMIT 100
),
     random_posts AS (
         SELECT id AS post_id, row_number() OVER () AS rn
         FROM posts
         ORDER BY random()
         LIMIT 100
     ),
     random_parents AS (
         SELECT id AS parent_id, row_number() OVER () AS rn
         FROM comments
         ORDER BY random()
         LIMIT 100
     ),
     comment_data AS (
         SELECT
             ru.user_id,
             rp.post_id,
             CASE
                 WHEN random() < 0.5 THEN NULL
                 ELSE rcom.parent_id
                 END AS parent_id,
             md5(random()::text || clock_timestamp()::text)::uuid AS uuid,
             'Random comment #' || ru.rn AS content,
             now() - (random() * interval '30 days') AS created_at,
             now() AS updated_at
         FROM random_users ru
                  JOIN random_posts rp ON ru.rn = rp.rn
                  LEFT JOIN random_parents rcom ON ru.rn = rcom.rn
     )

-- Step 2: Insert randomized comments
INSERT INTO comments (uuid, post_id, author_id, parent_id, content, created_at, updated_at)
SELECT uuid, post_id, user_id, parent_id, content, created_at, updated_at
FROM comment_data;
