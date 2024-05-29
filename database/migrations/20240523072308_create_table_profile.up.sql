CREATE TABLE profile
(
    id          INT             NOT NULL AUTO_INCREMENT,
    fullname    VARCHAR(128)    NOT NULL,
    age         VARCHAR(128)    NOT NULL,
    alamat      VARCHAR(128)    NOT NULL,
    createt_at  TIMESTAMP       NULL,
    updated_at  TIMESTAMP       NULL,
    deleted_at  TIMESTAMP       NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB;