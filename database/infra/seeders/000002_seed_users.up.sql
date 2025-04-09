-- Insert 10 dummy users into the users table
INSERT INTO users (
    uuid,
    first_name,
    last_name,
    username,
    display_name,
    email,
    password_hash,
    public_token
    -- Columns with DEFAULT values (is_admin, verified_at, created_at, updated_at, deleted_at)
    -- and nullable columns (bio, picture_file_name, profile_picture_url)
    -- are omitted and will use their defaults or NULL.
) VALUES
      (gen_random_uuid(), 'Alice', 'Smith', 'alice', 'Alice S', 'alice@example.com', '$2a$10$PlaceholderHashForTestingAlice', gen_random_uuid()::varchar),
      (gen_random_uuid(), 'Bob', 'Johnson', 'bob', 'Bob J', 'bob@example.com', '$2a$10$PlaceholderHashForTestingBob', gen_random_uuid()::varchar),
      (gen_random_uuid(), 'Charlie', 'Williams', 'charlie', 'Charlie W', 'charlie@example.com', '$2a$10$PlaceholderHashForTestingCharlie', gen_random_uuid()::varchar),
      (gen_random_uuid(), 'Diana', 'Brown', 'diana', 'Diana B', 'diana@example.com', '$2a$10$PlaceholderHashForTestingDiana', gen_random_uuid()::varchar),
      (gen_random_uuid(), 'Ethan', 'Jones', 'ethan', 'Ethan J', 'ethan@example.com', '$2a$10$PlaceholderHashForTestingEthan', gen_random_uuid()::varchar),
      (gen_random_uuid(), 'Fiona', 'Garcia', 'fiona', 'Fiona G', 'fiona@example.com', '$2a$10$PlaceholderHashForTestingFiona', gen_random_uuid()::varchar),
      (gen_random_uuid(), 'George', 'Miller', 'george', 'George M', 'george@example.com', '$2a$10$PlaceholderHashForTestingGeorge', gen_random_uuid()::varchar),
      (gen_random_uuid(), 'Hannah', 'Davis', 'hannah', 'Hannah D', 'hannah@example.com', '$2a$10$PlaceholderHashForTestingHannah', gen_random_uuid()::varchar),
      (gen_random_uuid(), 'Ian', 'Rodriguez', 'ian', 'Ian R', 'ian@example.com', '$2a$10$PlaceholderHashForTestingIan', gen_random_uuid()::varchar),
      (gen_random_uuid(), 'Jane', 'Martinez', 'jane', 'Jane M', 'jane@example.com', '$2a$10$PlaceholderHashForTestingJane', gen_random_uuid()::varchar);
