CREATE TABLE IF NOT EXISTS users (
    id              integer         primary key autoincrement,
    email           varchar(50)     not null unique,
    username        varchar(50)     not null unique,
    password_hash   varchar(255)    not null,
    token           varchar(50)     not null unique,
    expire_at       date            not null
);

CREATE TABLE IF NOT EXISTS posts (
    id			integer         primary key autoincrement,
    user_id	integer         not null,
    text		varchar(255),
    FOREIGN KEY (user_id)REFERENCES users(id)

);

CREATE TABLE IF NOT EXISTS categories (
    id			integer        primary key autoincrement,
    name		varchar(255)
);

CREATE TABLE IF NOT EXISTS posts_categories (
    id			integer     primary key autoincrement,
    post_id     integer     not null,
    category_id integer     not null,
    FOREIGN KEY (post_id)REFERENCES posts(id)
    FOREIGN KEY (category_id)REFERENCES categories(id)


);

CREATE TABLE IF NOT EXISTS comments (
    id			integer        primary key autoincrement,
    user_id	integer,
    post_id		integer,
    text		varchar(255),
    FOREIGN KEY (user_id)REFERENCES users(id),
    FOREIGN key (post_id)REFERENCES posts(id)
);

CREATE TABLE IF NOT EXISTS posts_likes (
    id			integer        primary key autoincrement,
    user_id	    integer,
    post_id		integer,
    type        BOOLEAN NOT NULL,
    UNIQUE      (post_id, user_id),
    FOREIGN KEY (user_id)   REFERENCES users(id),
    FOREIGN KEY (post_id)   REFERENCES posts(id)
);

CREATE TABLE IF NOT EXISTS comments_likes (
    id			integer        primary key autoincrement,
    user_id	    integer,
    comment_id	integer,
    type        BOOLEAN NOT NULL,
    UNIQUE      (comment_id, user_id),
    FOREIGN KEY (user_id)       REFERENCES users(id),
    FOREIGN KEY (comment_id)    REFERENCES comments(id)

);