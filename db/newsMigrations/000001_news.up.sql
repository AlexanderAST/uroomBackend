CREATE TABLE news(
    id bigserial not null PRIMARY KEY,
    date date not null,
    name varchar not null unique,
    small_description varchar not null,
    full_description varchar not null,
    image_path varchar not null
);