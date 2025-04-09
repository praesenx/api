-- Populate the post_tags table with 50 random, unique post-tag pairings.

-- Use CTEs to get active post and tag IDs
WITH all_posts AS (
    SELECT id FROM posts WHERE deleted_at IS NULL
),
     all_tags AS (
         SELECT id FROM tags WHERE deleted_at IS NULL -- Assuming 100 tags exist
     ),
-- Generate all possible unique combinations of posts and tags
     all_possible_pairs AS (
         SELECT p.id AS post_id, t.id AS tag_id
         FROM all_posts p
                  CROSS JOIN all_tags t -- Creates the Cartesian product (e.g., 50 posts * 100 tags = 5000 possible unique pairs)
     )
-- Insert exactly 50 random pairs from the possible combinations
INSERT INTO post_tags (post_id, tag_id)
SELECT post_id, tag_id
FROM all_possible_pairs
ORDER BY random() -- Randomize the order of all potential pairs
LIMIT 50          -- Take the first 50 distinct random pairs
ON CONFLICT (post_id, tag_id) DO NOTHING; -- Safety net: If a selected pair somehow already exists (e.g., if table wasn't empty), skip inserting it.
