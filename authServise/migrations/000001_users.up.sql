CREATE TABLE users (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    password VARCHAR(255),
    fullname VARCHAR(100),
    user_id UUID,
    timestamp VARCHAR(255),
    refresh_token VARCHAR(255),
    email VARCHAR(100),
    reason VARCHAR(100),
    bio TEXT,
    ecopoints INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
