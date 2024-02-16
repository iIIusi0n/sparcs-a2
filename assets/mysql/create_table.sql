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

INSERT INTO hospitals (name, latitude, longitude, number_of_doctor, address, phone_number) VALUES ('아이사랑마음병원', 36.352178, 127.382715, 4, '대전광역시 서구 대전시청 현대빌딩', '010-1234-5678');
INSERT INTO hospitals (name, latitude, longitude, number_of_doctor, address, phone_number) VALUES ('우리아이소아청소년의과', 36.351157, 127.384953, 6, '대전광역시 서구 둔산동 마음아파트상가', '010-1234-5678');
INSERT INTO hospitals (name, latitude, longitude, number_of_doctor, address, phone_number) VALUES ('대전아이병원', 36.356418, 127.390434, 7, '대전광역시 서구 갈마동 신가아파트', '010-1234-5678');
INSERT INTO hospitals (name, latitude, longitude, number_of_doctor, address, phone_number) VALUES ('전부다줌청소년의과', 36.357860, 127.388796, 5, '대전광역시 서구 둔산동 마음아파트상가', '010-1234-5678');

INSERT INTO users (username, real_name, email, phone_number) VALUES ('user1', '홍길동', 'x', '010-1234-5678');

INSERT INTO waiting_numbers (user_id, hospital_id, number) VALUES (1, 1, 17);
INSERT INTO waiting_numbers (user_id, hospital_id, number) VALUES (1, 2, 9);
INSERT INTO waiting_numbers (user_id, hospital_id, number) VALUES (1, 3, 25);
INSERT INTO waiting_numbers (user_id, hospital_id, number) VALUES (1, 4, 13);

INSERT INTO chat_rooms (name) VALUES ('서구 7세 어린이 보호자들 모임');
INSERT INTO chat_rooms (name) VALUES ('유성구 소아과 정보 공유방');

INSERT INTO chats (room_id, sender_id, content) VALUES (1, 1, '안녕하세요');
INSERT INTO chats (room_id, sender_id, content) VALUES (1, 1, '어디서 대기중이세요?');
INSERT INTO chats (room_id, sender_id, content) VALUES (1, 1, '아이가 아프네요');

INSERT INTO chats (room_id, sender_id, content) VALUES (2, 1, '안녕하세요');
INSERT INTO chats (room_id, sender_id, content) VALUES (2, 1, '오늘 사람이 많네요..');
INSERT INTO chats (room_id, sender_id, content) VALUES (2, 1, '여기 마음병원 의사분들이 엄청 친절하시네요!!');
