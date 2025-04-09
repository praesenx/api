-- Insert 10 sample categories
INSERT INTO categories (
    uuid,
    name,
    slug,
    description
    -- id, created_at, updated_at, deleted_at columns will use their defaults
) VALUES
      (gen_random_uuid(), 'Technology', 'technology', 'Articles and news about technology, gadgets, and software.'),
      (gen_random_uuid(), 'Travel', 'travel', 'Posts about travel destinations, tips, and experiences.'),
      (gen_random_uuid(), 'Food', 'food', 'Recipes, restaurant reviews, and discussions about food.'),
      (gen_random_uuid(), 'Programming', 'programming', 'Tutorials, discussions, and news related to software development.'),
      (gen_random_uuid(), 'Science', 'science', 'Exploring the world of science, research, and discoveries.'),
      (gen_random_uuid(), 'Lifestyle', 'lifestyle', 'Covering various aspects of daily life, wellness, and personal growth.'),
      (gen_random_uuid(), 'Business', 'business', 'Insights into the world of business, finance, and entrepreneurship.'),
      (gen_random_uuid(), 'Entertainment', 'entertainment', 'News and reviews about movies, music, games, and pop culture.'),
      (gen_random_uuid(), 'Health', 'health', 'Information and tips related to health, fitness, and well-being.'),
      (gen_random_uuid(), 'Sports', 'sports', 'Updates, analysis, and stories from the world of sports.');
