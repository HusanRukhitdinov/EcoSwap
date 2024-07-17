CREATE TABLE ratings (
    id UUID PRIMARY KEY,
    user_id UUID,
    rater_id UUID ,
    rating DECIMAL(2, 1) NOT NULL,
    comment TEXT,
    swap_id UUID REFERENCES swaps(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
