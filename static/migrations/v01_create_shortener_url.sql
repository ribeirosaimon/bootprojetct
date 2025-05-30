CREATE TABLE IF NOT EXISTS shortener_url (
     id           CHAR(36)      NOT NULL PRIMARY KEY,
    shortcode    VARCHAR(50)   NOT NULL UNIQUE,
    original     TEXT          NOT NULL,
    clicks       BIGINT        DEFAULT 0,
    created_by   VARCHAR(100)  NOT NULL,
    is_active    BOOLEAN       DEFAULT TRUE,
    created_at   DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expired_at   DATETIME      NULL,
    last_access  DATETIME      NULL
    );