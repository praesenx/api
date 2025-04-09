WITH all_posts AS (
    SELECT id FROM posts WHERE deleted_at IS NULL
),
     all_categories AS (
         SELECT id FROM categories WHERE deleted_at IS NULL
     ),
     post_coverage AS (
         SELECT
             p.id AS post_id,
             (SELECT id FROM all_categories ORDER BY random() LIMIT 1) AS category_id
         FROM all_posts p
     ),
     category_coverage AS (
         SELECT
             (SELECT id FROM all_posts ORDER BY random() LIMIT 1) AS post_id,
             c.id AS category_id
         FROM all_categories c
     ),
     additional_random_links AS (
         SELECT p.id AS post_id, c.category_id
         FROM all_posts p
                  CROSS JOIN LATERAL (
             SELECT id AS category_id
             FROM all_categories
             ORDER BY random()
             LIMIT 2
             ) c
     ),
     combined_assignments AS (
         SELECT post_id, category_id FROM post_coverage
         UNION
         SELECT post_id, category_id FROM category_coverage
         UNION
         SELECT post_id, category_id FROM additional_random_links
     )

INSERT INTO post_categories (post_id, category_id)
SELECT post_id, category_id
FROM combined_assignments
ON CONFLICT (post_id, category_id) DO NOTHING;
