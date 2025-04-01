CREATE TABLE vendors (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NOT NULL            -- Name of the rank
);

CREATE TABLE types (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NOT NULL            -- Tea type
);

CREATE TABLE subtypes (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,           -- Tea subtype
    region TEXT NOT NULL,         -- Region
    parent TEXT NOT NULL          -- parent of the type
);

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