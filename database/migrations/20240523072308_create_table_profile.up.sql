CREATE TABLE profiles
(
    id          INT             NOT NULL AUTO_INCREMENT,
    Fullname    VARCHAR(128)    NOT NULL,
    Age         VARCHAR(128)    NOT NULL,
    Alamat      VARCHAR(128)    NOT NULL,
    -- jenis_kelamin       VARCHAR(128)    NOT NULL,
    -- tanggal_lahir       VARCHAR(128)    NOT NULL,
    -- tempat_lahir        VARCHAR(128)    NOT NULL,
    -- umur                VARCHAR(128)    NOT NULL,
    -- status_pernikahan   VARCHAR(128)    NULL,
    -- domisili            VARCHAR(128)    NULL,
    -- alamat_rumah        VARCHAR(128)    NULL,
    -- nomor_handphone     VARCHAR(128)    NULL,
    -- nomor_telpon        VARCHAR(128)    NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB;