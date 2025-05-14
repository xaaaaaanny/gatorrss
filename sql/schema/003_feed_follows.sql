-- +goose Up
CREATE TABLE feed_follows(
                        id UUID primary key,
                        created_at timestamp not null,
                        updated_at timestamp not null,
                        user_id UUID references users(id) on delete cascade not null,
                        feed_id UUID references feeds(id) on delete cascade not null,
                        UNIQUE (user_id, feed_id));

-- +goose Down
DROP TABLE feed_follows;