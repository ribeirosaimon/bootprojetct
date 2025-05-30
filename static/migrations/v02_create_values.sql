INSERT INTO shortener_url (id, shortcode, original, clicks, created_by, is_active, created_at)
SELECT
    UUID(),
    SUBSTRING(MD5(RAND()*RAND()), 1, 8),
    CONCAT('https://site.com/page/', FLOOR(RAND()*100000)),
    0,
    CONCAT('user_', FLOOR(RAND()*10000)),
    TRUE,
    NOW()
FROM information_schema.COLUMNS
         LIMIT 100;