CREATE TABLE IF NOT EXISTS items (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    code text UNIQUE NOT NULL,
    name text, 
    description text,
    purchase_cost NUMERIC(16, 4),
    sell_price NUMERIC(16, 4),
    deleted_at timestamptz
);