// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: quiz.sql

package sqlc

import (
	"context"
)

const activateOnlyVersion = `-- name: ActivateOnlyVersion :exec
UPDATE quiz
SET active = 0
WHERE filename = ?
AND version <> ?
`

type ActivateOnlyVersionParams struct {
	Filename string `db:"filename"`
	Version  int64  `db:"version"`
}

func (q *Queries) ActivateOnlyVersion(ctx context.Context, arg ActivateOnlyVersionParams) error {
	_, err := q.db.ExecContext(ctx, activateOnlyVersion, arg.Filename, arg.Version)
	return err
}

const countAllActiveQuiz = `-- name: CountAllActiveQuiz :one
SELECT count(1)
FROM quiz
WHERE active = 1
`

func (q *Queries) CountAllActiveQuiz(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAllActiveQuiz)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createOrReplaceAnswer = `-- name: CreateOrReplaceAnswer :exec
REPLACE INTO quiz_answer (sha1, content, valid)
VALUES (?, ?, ?)
`

type CreateOrReplaceAnswerParams struct {
	Sha1    string `db:"sha1"`
	Content string `db:"content"`
	Valid   int64  `db:"valid"`
}

func (q *Queries) CreateOrReplaceAnswer(ctx context.Context, arg CreateOrReplaceAnswerParams) error {
	_, err := q.db.ExecContext(ctx, createOrReplaceAnswer, arg.Sha1, arg.Content, arg.Valid)
	return err
}

const createOrReplaceQuestion = `-- name: CreateOrReplaceQuestion :exec
REPLACE INTO quiz_question (sha1, content)
VALUES (?, ?)
`

type CreateOrReplaceQuestionParams struct {
	Sha1    string `db:"sha1"`
	Content string `db:"content"`
}

func (q *Queries) CreateOrReplaceQuestion(ctx context.Context, arg CreateOrReplaceQuestionParams) error {
	_, err := q.db.ExecContext(ctx, createOrReplaceQuestion, arg.Sha1, arg.Content)
	return err
}

const createOrReplaceQuiz = `-- name: CreateOrReplaceQuiz :exec
REPLACE INTO quiz (sha1, name, filename, version, created_at)
VALUES (?, ?, ?, ?, ?)
`

type CreateOrReplaceQuizParams struct {
	Sha1      string `db:"sha1"`
	Name      string `db:"name"`
	Filename  string `db:"filename"`
	Version   int64  `db:"version"`
	CreatedAt string `db:"created_at"`
}

func (q *Queries) CreateOrReplaceQuiz(ctx context.Context, arg CreateOrReplaceQuizParams) error {
	_, err := q.db.ExecContext(ctx, createOrReplaceQuiz,
		arg.Sha1,
		arg.Name,
		arg.Filename,
		arg.Version,
		arg.CreatedAt,
	)
	return err
}

const findAllActiveQuiz = `-- name: FindAllActiveQuiz :many
SELECT sha1, name, filename, version, active, created_at
FROM quiz
WHERE active = 1
LIMIT ? OFFSET ?
`

type FindAllActiveQuizParams struct {
	Limit  int64 `db:"limit"`
	Offset int64 `db:"offset"`
}

func (q *Queries) FindAllActiveQuiz(ctx context.Context, arg FindAllActiveQuizParams) ([]Quiz, error) {
	rows, err := q.db.QueryContext(ctx, findAllActiveQuiz, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Quiz{}
	for rows.Next() {
		var i Quiz
		if err := rows.Scan(
			&i.Sha1,
			&i.Name,
			&i.Filename,
			&i.Version,
			&i.Active,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findQuizByFilenameAndLatestVersion = `-- name: FindQuizByFilenameAndLatestVersion :one
SELECT sha1, name, filename, version, active, created_at
FROM quiz
WHERE filename = ?
ORDER BY version DESC
LIMIT 1
`

func (q *Queries) FindQuizByFilenameAndLatestVersion(ctx context.Context, filename string) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, findQuizByFilenameAndLatestVersion, filename)
	var i Quiz
	err := row.Scan(
		&i.Sha1,
		&i.Name,
		&i.Filename,
		&i.Version,
		&i.Active,
		&i.CreatedAt,
	)
	return i, err
}

const findQuizBySha1 = `-- name: FindQuizBySha1 :one
SELECT sha1, name, filename, version, active, created_at
FROM quiz
WHERE sha1 = ?
`

func (q *Queries) FindQuizBySha1(ctx context.Context, sha1 string) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, findQuizBySha1, sha1)
	var i Quiz
	err := row.Scan(
		&i.Sha1,
		&i.Name,
		&i.Filename,
		&i.Version,
		&i.Active,
		&i.CreatedAt,
	)
	return i, err
}

const findQuizFullBySha1 = `-- name: FindQuizFullBySha1 :many
SELECT
    q.sha1 as quiz_sha1,
    q.filename as quiz_filename,
    q.name as quiz_name,
    q.version as quiz_version,
    q.created_at as quiz_created_at,
    q.active as quiz_active,
    qq.sha1 as question_sha1,
    qq.content as question_content,
    qa.sha1 as answer_sha1,
    qa.content as answer_content,
    qa.valid as answer_valid
FROM quiz q
         JOIN quiz_question_quiz qqq ON q.sha1 = qqq.quiz_sha1
         JOIN quiz_question qq ON qq.sha1 = qqq.question_sha1
         JOIN quiz_question_answer qqa ON qq.sha1 = qqa.question_sha1
         JOIN quiz_answer qa ON qa.sha1 = qqa.answer_sha1
WHERE q.sha1 = ?
`

type FindQuizFullBySha1Row struct {
	QuizSha1        string `db:"quiz_sha1"`
	QuizFilename    string `db:"quiz_filename"`
	QuizName        string `db:"quiz_name"`
	QuizVersion     int64  `db:"quiz_version"`
	QuizCreatedAt   string `db:"quiz_created_at"`
	QuizActive      int64  `db:"quiz_active"`
	QuestionSha1    string `db:"question_sha1"`
	QuestionContent string `db:"question_content"`
	AnswerSha1      string `db:"answer_sha1"`
	AnswerContent   string `db:"answer_content"`
	AnswerValid     int64  `db:"answer_valid"`
}

func (q *Queries) FindQuizFullBySha1(ctx context.Context, sha1 string) ([]FindQuizFullBySha1Row, error) {
	rows, err := q.db.QueryContext(ctx, findQuizFullBySha1, sha1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindQuizFullBySha1Row{}
	for rows.Next() {
		var i FindQuizFullBySha1Row
		if err := rows.Scan(
			&i.QuizSha1,
			&i.QuizFilename,
			&i.QuizName,
			&i.QuizVersion,
			&i.QuizCreatedAt,
			&i.QuizActive,
			&i.QuestionSha1,
			&i.QuestionContent,
			&i.AnswerSha1,
			&i.AnswerContent,
			&i.AnswerValid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const linkAnswer = `-- name: LinkAnswer :exec
REPLACE INTO quiz_question_answer (question_sha1, answer_sha1)
VALUES (?, ?)
`

type LinkAnswerParams struct {
	QuestionSha1 string `db:"question_sha1"`
	AnswerSha1   string `db:"answer_sha1"`
}

func (q *Queries) LinkAnswer(ctx context.Context, arg LinkAnswerParams) error {
	_, err := q.db.ExecContext(ctx, linkAnswer, arg.QuestionSha1, arg.AnswerSha1)
	return err
}

const linkQuestion = `-- name: LinkQuestion :exec
REPLACE INTO quiz_question_quiz (quiz_sha1, question_sha1)
VALUES (?, ?)
`

type LinkQuestionParams struct {
	QuizSha1     string `db:"quiz_sha1"`
	QuestionSha1 string `db:"question_sha1"`
}

func (q *Queries) LinkQuestion(ctx context.Context, arg LinkQuestionParams) error {
	_, err := q.db.ExecContext(ctx, linkQuestion, arg.QuizSha1, arg.QuestionSha1)
	return err
}
