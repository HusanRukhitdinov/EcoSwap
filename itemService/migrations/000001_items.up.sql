CREATE TABLE items (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    category_id UUID REFERENCES item_categories(id),
    condition VARCHAR(20) NOT NULL,
    swap_preference JSONB,
    owner_id UUID,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
