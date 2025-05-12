-- +goose Up
CREATE TABLE feeds(
                    id UUID primary key,
                    created_at timestamp not null,
                    updated_at timestamp not null,
                    name text not null,
                    url text not null unique,
                    user_id UUID references users(id) on delete cascade not null );

-- +goose Down
DROP TABLE feeds;