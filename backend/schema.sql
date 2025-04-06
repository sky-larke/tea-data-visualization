CREATE TABLE teas (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    year INTEGER NOT NULL,
    rank INTEGER NOT NULL, 
    vendor TEXT NOT NULL,
    name TEXT NOT NULL,           
    type TEXT NOT NULL,
    subtype TEXT,
    cultivar TEXT,
    cost NUMERIC NOT NULL,
    amount INTEGER NOT NULL
);