-- Insert 100 sample tags using generate_series for unique names/slugs
INSERT INTO tags (
    uuid,
    name,
    slug
    -- description is nullable, id/timestamps use defaults
)
SELECT
    gen_random_uuid(),             -- Generate a unique UUID for each tag
    'Sample Tag ' || gs.i,       -- Create a unique name like 'Sample Tag 1', 'Sample Tag 2', ...
    'sample-tag-' || gs.i        -- Create a unique slug like 'sample-tag-1', 'sample-tag-2', ...
FROM
    generate_series(1, 100) AS gs(i); -- Generate numbers from 1 to 100, aliased as gs(i)
