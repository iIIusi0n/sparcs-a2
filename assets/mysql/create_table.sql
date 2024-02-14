CREATE TABLE users
(
    id           INT                 NOT NULL AUTO_INCREMENT,
    username     VARCHAR(255) UNIQUE NOT NULL,
    real_name    VARCHAR(255)        NOT NULL,
    email        VARCHAR(255)        NOT NULL,
    phone_number VARCHAR(255)        NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE hospitals
(
    id               INT            NOT NULL AUTO_INCREMENT,
    name             VARCHAR(255)   NOT NULL,
    latitude         DECIMAL(10, 8) NOT NULL,
    longitude        DECIMAL(11, 8) NOT NULL,
    number_of_doctor INT            NOT NULL,
    address          TEXT           NOT NULL,
    phone_number     VARCHAR(255)   NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE in_hospitals
(
    hospital_id INT NOT NULL,
    user_id     INT NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (hospital_id) REFERENCES hospitals (id)
);

CREATE TABLE waiting_numbers
(
    user_id     INT NOT NULL,
    hospital_id INT NOT NULL,
    number      INT NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (hospital_id) REFERENCES hospitals (id)
);

CREATE TABLE chat_rooms
(
    id         INT          NOT NULL AUTO_INCREMENT,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE chats
(
    id         INT  NOT NULL AUTO_INCREMENT,
    room_id    INT  NOT NULL,
    sender_id  INT  NOT NULL,
    content    TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (room_id) REFERENCES chat_rooms (id),
    FOREIGN KEY (sender_id) REFERENCES users (id)
);

CREATE TABLE chat_read
(
    chat_id    INT NOT NULL,
    user_id    INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chat_id) REFERENCES chats (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
