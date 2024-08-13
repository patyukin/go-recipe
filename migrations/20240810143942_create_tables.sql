-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE authors
(
    id      UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id UUID NOT NULL,
    bio     TEXT
);

CREATE TABLE recipes
(
    id           UUID      DEFAULT uuid_generate_v4() PRIMARY KEY,
    title        VARCHAR(100) NOT NULL,
    instructions TEXT         NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ingredients
(
    id   UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    unit VARCHAR(50)  NOT NULL
);

CREATE TABLE recipe_ingredients
(
    recipe_id     UUID REFERENCES recipes (id) ON DELETE CASCADE,
    ingredient_id UUID REFERENCES ingredients (id) ON DELETE CASCADE,
    quantity      DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (recipe_id, ingredient_id)
);

CREATE TABLE recipe_authors
(
    recipe_id UUID REFERENCES recipes (id) ON DELETE CASCADE,
    author_id UUID REFERENCES authors (id) ON DELETE CASCADE,
    PRIMARY KEY (recipe_id, author_id)
);


-- +goose Down
DROP TABLE IF EXISTS recipe_authors;
DROP TABLE IF EXISTS recipe_ingredients;
DROP TABLE IF EXISTS ingredients;
DROP TABLE IF EXISTS recipes;
DROP TABLE IF EXISTS authors;
