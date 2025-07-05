create SCHEMA IF NOT EXISTS users;
CREATE TABLE IF NOT EXISTS users.t_author
(
    id      bigserial PRIMARY KEY,
    name    VARCHAR(50) UNIQUE
);

create SCHEMA IF NOT EXISTS articles;
CREATE TABLE IF NOT EXISTS articles.t_articles
(
    id      bigserial PRIMARY KEY,
    author_id int REFERENCES users.t_author (id),
    title TEXT,
    body TEXT,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);