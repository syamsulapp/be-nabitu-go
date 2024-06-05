CREATE TABLE users
(
    id              INT             NOT NULL AUTO_INCREMENT,
    username        VARCHAR(128)    NOT NULL,
    email           VARCHAR(128)    NOT NULL,
    password        VARCHAR(128)    NOT NULL,
    created_at      TIMESTAMP       NULL,
    updated_at      TIMESTAMP       NULL,
    deleted_at      TIMESTAMP       NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

CREATE TABLE profiles
(
    id          INT             NOT NULL AUTO_INCREMENT,
    fullname    VARCHAR(128)    NOT NULL,
    age         VARCHAR(128)    NOT NULL,
    alamat      VARCHAR(128)    NOT NULL,
    users_id    INT             NOT NULL,
    createt_at  TIMESTAMP       NULL,
    updated_at  TIMESTAMP       NULL,
    deleted_at  TIMESTAMP       NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (users_id) REFERENCES users(id)
)ENGINE = InnoDB;