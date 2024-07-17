CREATE TABLE recycling_submissions (
    id UUID PRIMARY KEY,
    center_id UUID REFERENCES recycling_centers(id),
    user_id UUID ,
    items JSONB NOT NULL,
    eco_points_earned INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);