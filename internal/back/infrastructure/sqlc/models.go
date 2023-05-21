// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc

import ()

type Quiz struct {
	Sha1      string `db:"sha1"`
	Name      string `db:"name"`
	Filename  string `db:"filename"`
	Version   int64  `db:"version"`
	Active    int64  `db:"active"`
	CreatedAt string `db:"created_at"`
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

type Role struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type User struct {
	ID        string `db:"id"`
	Email     string `db:"email"`
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	Active    int64  `db:"active"`
	RoleID    int64  `db:"role_id"`
}
