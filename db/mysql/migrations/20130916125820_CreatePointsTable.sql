-- +goose Up
CREATE TABLE points (lat float, lng float);

-- +goose Down
DROP TABLE points;
