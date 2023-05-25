// Code generated by sqlc-addon. DO NOT EDIT.
// versions:
//   sqlc-addon v1.6.1

package db

import (
	"database/sql"
	"log"
)

const testVersionTableExists = `
SELECT COUNT(name) FROM sqlite_master WHERE type='table' AND name='db_version';
`

const initSql = `
CREATE TABLE db_version
(
    version_number INTEGER NOT NULL
);
INSERT INTO db_version (version_number) VALUES (0);
`

const selectVersionStmt = `
SELECT version_number FROM db_version;
`

const updateVersionStmt = `
UPDATE db_version
SET version_number = ?
WHERE 1;
`

const v1Init = `
CREATE TABLE quiz
(
    sha1       TEXT PRIMARY KEY,
    name       TEXT    NOT NULL,
    filename   TEXT    NOT NULL,
    version    INTEGER NOT NULL DEFAULT 1,
    active     INTEGER NOT NULL DEFAULT 1,
    created_at TEXT    NOT NULL,
    duration   INTEGER NOT NULL,

    CONSTRAINT filename_fk FOREIGN KEY (filename) REFERENCES quiz (filename),
    CONSTRAINT quiz_version_unique UNIQUE (filename, version)
);

CREATE TABLE quiz_question
(
    sha1    TEXT PRIMARY KEY,
    content TEXT NOT NULL
);

CREATE TABLE quiz_question_quiz
(
    quiz_sha1     TEXT NOT NULL,
    question_sha1 TEXT NOT NULL,

    CONSTRAINT pk PRIMARY KEY (quiz_sha1, question_sha1),
    CONSTRAINT quiz_fk FOREIGN KEY (quiz_sha1) REFERENCES quiz (sha1),
    CONSTRAINT question_fk FOREIGN KEY (question_sha1) REFERENCES quiz_question (sha1)
);

CREATE TABLE quiz_answer
(
    sha1    TEXT PRIMARY KEY,
    valid   INTEGER NOT NULL,
    content TEXT    NOT NULL
);

CREATE TABLE quiz_question_answer
(
    question_sha1 TEXT NOT NULL,
    answer_sha1   TEXT NOT NULL,

    CONSTRAINT question_fk FOREIGN KEY (question_sha1) REFERENCES quiz_question (sha1),
    CONSTRAINT answer_fk FOREIGN KEY (answer_sha1) REFERENCES quiz_answer (sha1)
);
`

const v2Auth = `
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
    active    INTEGER NOT NULL DEFAULT 1,
    role_id   INTEGER NOT NULL,

    CONSTRAINT role_fk FOREIGN KEY (role_id) REFERENCES role (id)
);
`

const v3Session = `
CREATE TABLE session
(
    uuid       uuid PRIMARY KEY,
    quiz_sha1  TEXT      NOT NULL,
    user_id    TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL,

    CONSTRAINT quiz_user_unique UNIQUE (quiz_sha1, user_id),
    CONSTRAINT quiz_fk FOREIGN KEY (quiz_sha1) REFERENCES quiz (sha1),
    CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES user (id)
);

CREATE TABLE session_answer
(
    session_uuid  TEXT    NOT NULL,
    question_sha1 TEXT    NOT NULL,
    answer_sha1   TEXT    NOT NULL,
    checked       INTEGER NOT NULL,

    CONSTRAINT pk PRIMARY KEY (session_uuid, question_sha1, answer_sha1),
    CONSTRAINT session_fk FOREIGN KEY (session_uuid) REFERENCES session (uuid),
    CONSTRAINT question_fk FOREIGN KEY (question_sha1) REFERENCES quiz_question (sha1),
    CONSTRAINT answer_fk FOREIGN KEY (answer_sha1) REFERENCES quiz_answer (sha1)
);

CREATE VIEW session_view
AS
SELECT s.uuid                                                                       AS uuid,
       q.sha1                                                                       AS quiz_sha1,
       q.name                                                                       AS quiz_name,
       q.active                                                                     AS quiz_active,
       u.id                                                                         AS user_id,
       CAST(u.firstname || ' ' || u.lastname AS TEXT)                               AS user_name,
       CAST(MAX(q.duration - (STRFTIME('%s', 'now') - s.created_at), 0) AS INTEGER) AS remaining_sec
FROM session s
         JOIN quiz q ON q.sha1 = s.quiz_sha1
         JOIN user u ON u.id = s.user_id;

CREATE VIEW quiz_answer_count_view
AS
SELECT q.sha1          AS quiz_sha1,
       COUNT(qa.valid) AS checked_answers
FROM quiz q
         JOIN quiz_question_quiz qqq ON q.sha1 = qqq.quiz_sha1
         JOIN quiz_question_answer qqa ON qqq.question_sha1 = qqa.question_sha1
         JOIN quiz_answer qa ON qa.sha1 = qqa.answer_sha1
GROUP BY q.sha1;

CREATE VIEW session_response_view
AS
SELECT qqq.quiz_sha1,
       qqq.question_sha1,
       qqa.answer_sha1,
       s.uuid  AS session_uuid,
       s.user_id,
       sa.checked,
       CASE
           WHEN checked IS NOT NULL
               THEN CASE
                        WHEN qa.valid == sa.checked
                            THEN 1
                        ELSE 0
               END
           END AS result
FROM quiz_question_quiz qqq
         JOIN quiz_question_answer qqa ON qqq.question_sha1 = qqa.question_sha1
         JOIN quiz_answer qa ON qa.sha1 = qqa.answer_sha1
         LEFT JOIN session s ON qqq.quiz_sha1 = s.quiz_sha1
         LEFT JOIN session_answer sa ON qa.sha1 = sa.answer_sha1 AND sa.question_sha1 = qqq.question_sha1 AND
                                        sa.answer_sha1 = qqa.answer_sha1;
`

var migrations = map[int]string{
	1: v1Init,
	2: v2Auth,
	3: v3Session,
}

type DB interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Prepare(string) (*sql.Stmt, error)
}

// New creates a new instance of Migrations struct
func New(db DB) *Migrations {
	return &Migrations{db: db}
}

type Migrations struct {
	db DB
}

// Migrate migrates the database using the migration scripts provided
func (m *Migrations) Migrate() {
	initialized, err := m.isInitialized()
	if err != nil {
		log.Fatalf("Can't detect if database is initialized %v", err)
	}
	if initialized {
		version, err := m.getVersion()
		if err != nil {
			log.Fatalf("Can't read database version %v", err)
		}
		m.applyMigration(version)
	} else {
		m.createDBVersionTable()
		m.applyMigration(0)
	}
}

// isInitialized checks if the table db_version is present in the current database
func (m *Migrations) isInitialized() (bool, error) {
	stmt, err := m.db.Prepare(testVersionTableExists)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var tablePresent int
	err = stmt.QueryRow().Scan(&tablePresent)
	if err != nil {
		return false, err
	}

	return tablePresent == 1, nil
}

// getVersion returns the current version of the schema
func (m *Migrations) getVersion() (int, error) {
	stmt, err := m.db.Prepare(selectVersionStmt)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var version int
	err = stmt.QueryRow().Scan(&version)
	if err != nil {
		return 0, err
	}

	return version, nil
}

// applyMigration a migration
func (m *Migrations) createDBVersionTable() {
	_, err := m.db.Exec(initSql)
	if err != nil {
		log.Fatalf("Could not create db_version table %v", err)
	}
}

// applyMigration a migration
func (m *Migrations) applyMigration(fromVersion int) {
	updStmt, err := m.db.Prepare(updateVersionStmt)
	if err != nil {
		log.Fatalf("Could not prepare Stmt : %v", err)
	}
	defer updStmt.Close()

	for version, script := range migrations {
		if version > fromVersion {
			_, err := m.db.Exec(script)
			if err != nil {
				log.Fatalf("Could not apply migration : %s, %v", script, err)
			}

			_, err = updStmt.Exec(version)
			if err != nil {
				log.Fatalf("Could not update version : %v", err)
			}
		}
	}
}
