WITH
    random_users AS (
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
    )

INSERT INTO post_views (post_id, user_id, ip_address, user_agent, viewed_at)
SELECT
    rp.post_id,
    ru.user_id,
    ('192.168.' || floor(random() * 255)::int || '.' || floor(random() * 255)::int)::inet AS ip_address,
    'User-Agent ' || floor(random() * 1000)::int AS user_agent,
    now() - (random() * interval '30 days') AS viewed_at
FROM
    random_users ru
    JOIN random_posts rp ON ru.rn = rp.rn;
