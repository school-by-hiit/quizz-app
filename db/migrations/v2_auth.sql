CREATE TABLE role
(
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO role (id, name)
VALUES (1, 'Admin');
INSERT INTO role (id, name)
VALUES (2, 'Teacher');
INSERT INTO role (id, name)
VALUES (3, 'Student');

CREATE TABLE user
(
    id        TEXT PRIMARY KEY,
    email     TEXT    NOT NULL,
    firstname TEXT    NOT NULL,
    lastname  TEXT    NOT NULL,
    active    INTEGER NOT NULL DEFAULT 1
);

CREATE TABLE user_role
(
    user_id TEXT,
    role_id INTEGER,

    CONSTRAINT pk PRIMARY KEY (user_id, role_id),
    CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES user (id),
    CONSTRAINT role_fk FOREIGN KEY (role_id) REFERENCES role (id)
);

CREATE TABLE token
(
    opaque_token TEXT PRIMARY KEY,
    user_id      TEXT      NOT NULL,
    expires      TIMESTAMP NOT NULL,
    aud          TEXT      NOT NULL,

    CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES user (id)
);