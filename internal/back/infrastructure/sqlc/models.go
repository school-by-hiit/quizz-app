// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc

import (
	"database/sql"
)

type Quizz struct {
	Filename string `db:"filename"`
}

type QuizzAnswer struct {
	Sha1    string        `db:"sha1"`
	Valid   sql.NullInt64 `db:"valid"`
	Content string        `db:"content"`
}

type QuizzQuestion struct {
	Sha1    string `db:"sha1"`
	Content string `db:"content"`
}

type QuizzQuestionAnswer struct {
	QuestionSha1 sql.NullString `db:"question_sha1"`
	AnswerSha1   sql.NullString `db:"answer_sha1"`
}

type QuizzQuestionVersion struct {
	VersionSha1  sql.NullString `db:"version_sha1"`
	QuestionSha1 sql.NullString `db:"question_sha1"`
}

type QuizzVersion struct {
	Sha1     string         `db:"sha1"`
	Filename sql.NullString `db:"filename"`
	Version  int64          `db:"version"`
	Active   sql.NullInt64  `db:"active"`
}
