/*
 * Copyright (c) 2023 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package presentation

import (
	"net/http"
	"regexp"

	"github.com/google/uuid"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
)

type Quiz struct {
	Sha1      string         `json:"sha1"`
	Filename  string         `json:"filename"`
	Name      string         `json:"name"`
	Version   int            `json:"version"`
	CreatedAt string         `json:"createdAt"`
	Duration  int            `json:"duration"`
	Active    bool           `json:"active"`
	Questions []QuizQuestion `json:"questions,omitempty"`
}

type QuizQuestion struct {
	Sha1    string               `json:"sha1"`
	Content string               `json:"content"`
	Answers []QuizQuestionAnswer `json:"answers,omitempty"`
}

type QuizQuestionAnswer struct {
	Sha1    string `json:"sha1"`
	Content string `json:"content"`
}

func (q *Quiz) fromDomain(domain *domain.Quiz) *Quiz {
	q.Sha1 = domain.Sha1
	q.Filename = domain.Filename
	q.Name = domain.Name
	q.Version = domain.Version
	q.Duration = domain.Duration
	q.CreatedAt = domain.CreatedAt
	q.Active = domain.Active
	q.Questions = make([]QuizQuestion, len(domain.Questions))

	i := 0
	for _, question := range domain.Questions {

		j := 0
		answers := make([]QuizQuestionAnswer, len(question.Answers))
		for _, a := range question.Answers {
			answers[j] = QuizQuestionAnswer{
				Sha1:    a.Sha1,
				Content: a.Content,
			}
			j++
		}

		q.Questions[i] = QuizQuestion{
			Sha1:    question.Sha1,
			Content: question.Content,
			Answers: answers,
		}
		i++
	}

	return q
}

func toQuizDtos(domains []*domain.Quiz) []*Quiz {
	dtos := make([]*Quiz, len(domains))

	for i, d := range domains {
		dto := &Quiz{}
		dtos[i] = dto.fromDomain(d)
	}

	return dtos
}

type endPointDef struct {
	regex  *regexp.Regexp
	method string
}

func (e *endPointDef) match(request *http.Request) bool {
	path := request.URL.Path
	method := request.Method

	return e.regex.MatchString(path) && e.method == method
}

type RegisterRequestBody struct {
	Id        string `json:"id" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Active    bool   `json:"active"`
	Role      string `json:"role,omitempty"`
}

func (u *User) fromDomain(d *domain.User) *User {
	u.Id = d.Id
	u.Email = d.Email
	u.Firstname = d.Firstname
	u.Lastname = d.Lastname
	u.Active = d.Active
	switch d.Role {
	case domain.Admin:
		u.Role = "ADMIN"
	case domain.Teacher:
		u.Role = "TEACHER"
	case domain.Student:
		u.Role = "STUDENT"
	}

	return u
}

type SessionResult struct {
	GoodAnswer  int `json:"goodAnswer,omitempty"`
	TotalAnswer int `json:"totalAnswer,omitempty"`
}

type Session struct {
	Id           uuid.UUID      `json:"id"`
	QuizSha1     string         `json:"quizSha1,omitempty"`
	QuizName     string         `json:"quizName,omitempty"`
	QuizActive   bool           `json:"quizActive,omitempty"`
	UserId       string         `json:"userId,omitempty"`
	UserName     string         `json:"userName,omitempty"`
	RemainingSec int            `json:"remainingSec,omitempty"`
	Result       *SessionResult `json:"result,omitempty"`
}

func (s *Session) fromDomain(d *domain.Session) *Session {
	s.Id = d.Id
	s.QuizSha1 = d.QuizSha1
	s.UserName = d.UserName
	s.QuizActive = d.QuizActive
	s.UserId = d.UserId
	s.UserName = d.UserName
	s.RemainingSec = d.RemainingSec
	if d.Result != nil {
		s.Result = &SessionResult{
			GoodAnswer:  d.Result.GoodAnswer,
			TotalAnswer: d.Result.TotalAnswer,
		}
	}

	return s
}

func toSessionDtos(domains []*domain.Session) []*Session {
	dtos := make([]*Session, len(domains))

	for i, d := range domains {
		dto := &Session{}
		dtos[i] = dto.fromDomain(d)
	}

	return dtos
}

type SessionAnswerRequestBody struct {
	QuestionSha1 string `json:"questionSha1" binding:"required"`
	AnswerSha1   string `json:"answerSha1" binding:"required"`
	Checked      bool   `json:"checked" binding:"required"`
}
