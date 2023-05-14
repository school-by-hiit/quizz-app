// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc

import (
	"time"
)

type Quiz struct {
	Sha1      string    `db:"sha1"`
	Name      string    `db:"name"`
	Filename  string    `db:"filename"`
	Version   int64     `db:"version"`
	Active    int64     `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

type QuizAnswer struct {
	Sha1    string `db:"sha1"`
	Valid   int64  `db:"valid"`
	Content string `db:"content"`
}

type QuizQuestion struct {
	Sha1    string `db:"sha1"`
	Content string `db:"content"`
}

type QuizQuestionAnswer struct {
	QuestionSha1 string `db:"question_sha1"`
	AnswerSha1   string `db:"answer_sha1"`
}

type QuizQuestionQuiz struct {
	QuizSha1     string `db:"quiz_sha1"`
	QuestionSha1 string `db:"question_sha1"`
}
