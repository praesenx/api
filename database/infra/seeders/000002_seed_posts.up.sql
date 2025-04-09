-- Use CTEs to get the IDs of the target users first
WITH target_users AS (
    SELECT id
    FROM users
    WHERE username IN (
        'alice', 'bob', 'charlie', 'diana', 'ethan',
        'fiona', 'george', 'hannah', 'ian', 'jane'
    )
    ORDER BY id
),
-- Assign a row number to each user ID to facilitate cycling
 numbered_users AS (
     SELECT id, ROW_NUMBER() OVER () as rn
     FROM target_users
 ),
-- Get the count of target users found
 user_count AS (
     SELECT COUNT(*) as total FROM target_users
 )

-- Insert 50 posts using generated data
INSERT INTO posts (
    uuid,
    author_id,
    slug,
    title,
    excerpt,
    content,
    published_at,
    cover_image_url
)
SELECT
    gen_random_uuid(),
    (SELECT id FROM numbered_users WHERE rn = (mod(gs.i - 1, uc.total) + 1)),
    'post-title-' || gs.i || '-' || substr(md5(random()::text), 1, 8),
    'Sample Post Title ' || gs.i,
    'This is the excerpt for sample post number ' || gs.i || '. It gives a brief overview of the content.',
    'This is the main content body for sample post number ' || gs.i || E'.\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.\n\nDuis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
    CASE WHEN random() < 0.8 THEN CURRENT_TIMESTAMP ELSE NULL END,
    'https://via.placeholder.com/800x400.png/00' || to_hex(mod(gs.i, 256)::int)::text || 'aa/FFFFFF?text=Post+' || gs.i
FROM
    generate_series(1, 50) AS gs(i), user_count uc; -- Make sure this semicolon is included!
