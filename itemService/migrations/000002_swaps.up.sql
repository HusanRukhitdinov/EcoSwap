CREATE TABLE swaps (
    id UUID PRIMARY KEY,
    offered_item_id UUID REFERENCES items(id),
    requested_item_id UUID REFERENCES items(id),
    requester_id UUID,
    owner_id UUID,
    status VARCHAR(20) NOT NULL,
    message TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
