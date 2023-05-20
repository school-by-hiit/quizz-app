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

package domain

import (
	"context"
)

//go:generate mockery --name QuizRepository
type QuizRepository interface {
	FindBySha1(ctx context.Context, sha1 string) (*Quiz, error)
	FindFullBySha1(ctx context.Context, sha1 string) (*Quiz, error)
	FindLatestVersionByFilename(ctx context.Context, filename string) (*Quiz, error)
	FindAllActive(ctx context.Context, limit uint16, offset uint16) ([]*Quiz, error)
	CountAllActive(ctx context.Context) (uint32, error)
	Create(ctx context.Context, quiz *Quiz) error
	ActivateOnlyVersion(ctx context.Context, filename string, version int) error
}

//go:generate mockery --name AuthRepository
type AuthRepository interface {
	FindUserById(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	AddRoleToUser(ctx context.Context, userId string, role Role) error
	RemoveAllRoleFromUser(ctx context.Context, userId string) error
	CreateToken(ctx context.Context, token *AccessToken) error
	FindTokenByTokenStr(ctx context.Context, tokenStr string) (*AccessToken, error)
}