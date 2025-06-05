INSERT INTO shortener_url (
    id, short_code, hash_size, original, clicks, created_by, is_active, created_at, expired_at, last_access
)
SELECT
    UUID(),
    SUBSTRING(MD5(RAND() * RAND()), 1, 8),
    8,
    CONCAT('https://site.com/page/', FLOOR(RAND() * 100000)),
    0,
    CONCAT('user_', FLOOR(RAND() * 10000)),
    TRUE,
    NOW() - INTERVAL FLOOR(RAND() * 60 * 24 * 60 * 60) SECOND,
    DATE_ADD(
            (NOW() - INTERVAL FLOOR(RAND() * 60 * 24 * 60 * 60) SECOND),
            INTERVAL FLOOR(RAND() * 30 * 24 * 60 * 60) SECOND
    ),
    NOW() - INTERVAL FLOOR(RAND() * 60 * 24 * 60 * 60) SECOND
FROM information_schema.COLUMNS
WHERE (SELECT COUNT(*) FROM shortener_url) < 100