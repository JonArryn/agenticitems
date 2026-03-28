CREATE TABLE IF NOT EXISTS items (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    code text UNIQUE NOT NULL,
    name text, 
    description text,
    purchase_cost_cents bigint,
    sell_price_cents bigint,
    deleted_at timestamptz
);