// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc

import (
	"time"

	"github.com/google/uuid"
)

type Quiz struct {
	Sha1      string `db:"sha1"`
	Name      string `db:"name"`
	Filename  string `db:"filename"`
	Version   int    `db:"version"`
	Active    bool   `db:"active"`
	CreatedAt string `db:"created_at"`
	Duration  int    `db:"duration"`
}

type QuizAnswer struct {
	Sha1    string `db:"sha1"`
	Valid   bool   `db:"valid"`
	Content string `db:"content"`
}

type QuizAnswerCountView struct {
	QuizSha1       string `db:"quiz_sha1"`
	CheckedAnswers int    `db:"checked_answers"`
}

type QuizClassVisibility struct {
	ClassUuid uuid.UUID `db:"class_uuid"`
	QuizSha1  string    `db:"quiz_sha1"`
}

type QuizQuestion struct {
	Sha1     string `db:"sha1"`
	Position int64  `db:"position"`
	Content  string `db:"content"`
}

type QuizQuestionAnswer struct {
	QuestionSha1 string `db:"question_sha1"`
	AnswerSha1   string `db:"answer_sha1"`
}

type QuizQuestionQuiz struct {
	QuizSha1     string `db:"quiz_sha1"`
	QuestionSha1 string `db:"question_sha1"`
}

type QuizSessionDetailView struct {
	SessionUuid      uuid.UUID `db:"session_uuid"`
	UserID           string    `db:"user_id"`
	RemainingSec     int       `db:"remaining_sec"`
	QuizSha1         string    `db:"quiz_sha1"`
	QuizName         string    `db:"quiz_name"`
	CheckedAnswers   int       `db:"checked_answers"`
	Results          int       `db:"results"`
	QuestionSha1     string    `db:"question_sha1"`
	QuestionContent  string    `db:"question_content"`
	QuestionPosition int       `db:"question_position"`
	AnswerSha1       string    `db:"answer_sha1"`
	AnswerContent    string    `db:"answer_content"`
	AnswerChecked    bool      `db:"answer_checked"`
	AnswerValid      bool      `db:"answer_valid"`
}

type QuizSessionView struct {
	QuizSha1       string    `db:"quiz_sha1"`
	QuizName       string    `db:"quiz_name"`
	QuizFilename   string    `db:"quiz_filename"`
	QuizVersion    int       `db:"quiz_version"`
	QuizDuration   int       `db:"quiz_duration"`
	QuizCreatedAt  string    `db:"quiz_created_at"`
	SessionUuid    uuid.UUID `db:"session_uuid"`
	UserID         string    `db:"user_id"`
	UserName       string    `db:"user_name"`
	RemainingSec   int       `db:"remaining_sec"`
	CheckedAnswers int       `db:"checked_answers"`
	Results        int       `db:"results"`
}

type Role struct {
	ID   int8   `db:"id"`
	Name string `db:"name"`
}

type Session struct {
	Uuid      uuid.UUID `db:"uuid"`
	QuizSha1  string    `db:"quiz_sha1"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

type SessionAnswer struct {
	SessionUuid  uuid.UUID `db:"session_uuid"`
	UserID       string    `db:"user_id"`
	QuestionSha1 string    `db:"question_sha1"`
	AnswerSha1   string    `db:"answer_sha1"`
	Checked      bool      `db:"checked"`
}

type SessionResponseView struct {
	QuizSha1     string      `db:"quiz_sha1"`
	QuestionSha1 string      `db:"question_sha1"`
	AnswerSha1   string      `db:"answer_sha1"`
	SessionUuid  uuid.UUID   `db:"session_uuid"`
	UserID       string      `db:"user_id"`
	Checked      bool        `db:"checked"`
	Result       interface{} `db:"result"`
}

type SessionView struct {
	Uuid           uuid.UUID `db:"uuid"`
	QuizSha1       string    `db:"quiz_sha1"`
	QuizName       string    `db:"quiz_name"`
	QuizActive     bool      `db:"quiz_active"`
	UserID         string    `db:"user_id"`
	UserName       string    `db:"user_name"`
	RemainingSec   int       `db:"remaining_sec"`
	CheckedAnswers int       `db:"checked_answers"`
	Results        int       `db:"results"`
}

type StudentClass struct {
	Uuid uuid.UUID `db:"uuid"`
	Name string    `db:"name"`
}

type User struct {
	ID        string    `db:"id"`
	Email     string    `db:"email"`
	Firstname string    `db:"firstname"`
	Lastname  string    `db:"lastname"`
	Active    bool      `db:"active"`
	RoleID    int8      `db:"role_id"`
	ClassUuid uuid.UUID `db:"class_uuid"`
}
