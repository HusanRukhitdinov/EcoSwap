CREATE TABLE challenge_participations (
    id UUID PRIMARY KEY,
    challenge_id UUID REFERENCES eco_challenges(id),
    user_id UUID,
    status VARCHAR(20) NOT NULL,
    recycled_items_count INTEGER DEFAULT 0,
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
