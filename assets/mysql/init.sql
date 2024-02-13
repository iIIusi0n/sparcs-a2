CREATE TABLE users
(
    id              INT                 NOT NULL AUTO_INCREMENT,
    username        VARCHAR(255) UNIQUE NOT NULL,
    name            VARCHAR(255) UNIQUE NOT NULL,
    profile_picture VARCHAR(255),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE posts
(
    id          INT          NOT NULL AUTO_INCREMENT,
    user_id     INT          NOT NULL,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    latitude    DECIMAL(10, 8),
    longitude   DECIMAL(11, 8),
    image_url   VARCHAR(255) NOT NULL,
    hashtags    VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE likes
(
    id         INT NOT NULL AUTO_INCREMENT,
    user_id    INT NOT NULL,
    post_id    INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (post_id) REFERENCES posts (id)
);

CREATE TABLE gatherings
(
    id               INT          NOT NULL AUTO_INCREMENT,
    user_id          INT          NOT NULL,
    title            VARCHAR(255) NOT NULL,
    latitude         DECIMAL(10, 8),
    longitude        DECIMAL(11, 8),
    date_time        DATETIME     NOT NULL,
    max_participants INT          NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE gathering_participants
(
    id           INT NOT NULL AUTO_INCREMENT,
    gathering_id INT NOT NULL,
    user_id      INT NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (gathering_id) REFERENCES gatherings (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
