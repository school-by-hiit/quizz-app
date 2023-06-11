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

package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/michaelcoll/quiz-app/internal/back/domain"
	"github.com/michaelcoll/quiz-app/internal/back/infrastructure/sqlc"
)

type UserDBRepository struct {
	domain.UserRepository

	q  *sqlc.Queries
	uc *cache.Cache
}

func NewUserRepository(c *sql.DB) *UserDBRepository {
	userCache := cache.New(30*time.Minute, 10*time.Minute)
	return &UserDBRepository{q: sqlc.New(c), uc: userCache}
}

func (r *UserDBRepository) FindUserById(ctx context.Context, id string) (*domain.User, error) {

	if user, found := r.uc.Get(id); found {
		return user.(*domain.User), nil
	}

	entity, err := r.q.FindUserById(ctx, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	user := r.toUser(entity)

	r.uc.Set(id, user, cache.DefaultExpiration)

	return user, nil
}

func (r *UserDBRepository) CreateOrReplaceUser(ctx context.Context, user *domain.User) error {
	err := r.q.CreateOrReplaceUser(ctx, sqlc.CreateOrReplaceUserParams{
		ID:        user.Id,
		Email:     user.Email,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		RoleID:    int64(user.Role),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *UserDBRepository) UpdateUserRole(ctx context.Context, userId string, role domain.Role) error {
	err := r.q.UpdateUserRole(ctx, sqlc.UpdateUserRoleParams{
		ID:     userId,
		RoleID: int64(role),
	})
	if err != nil {
		return err
	}

	r.uc.Delete(userId)

	return nil
}

func (r *UserDBRepository) FindAllUser(ctx context.Context) ([]*domain.User, error) {
	entities, err := r.q.FindAllUser(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(entities))

	for i, entity := range entities {
		users[i] = r.toUser(entity)
	}

	return users, nil
}

func (r *UserDBRepository) UpdateUserActive(ctx context.Context, id string, active bool) error {
	return r.q.UpdateUserActive(ctx, sqlc.UpdateUserActiveParams{
		Active: active,
		ID:     id,
	})
}

func (r *UserDBRepository) toUser(entity sqlc.User) *domain.User {
	return &domain.User{
		Id:        entity.ID,
		Email:     entity.Email,
		Firstname: entity.Firstname,
		Lastname:  entity.Lastname,
		Active:    entity.Active,
		Role:      r.toRole(entity.RoleID),
	}
}

func (r *UserDBRepository) toRole(entity int64) domain.Role {
	switch entity {
	case int64(domain.Admin):
		return domain.Admin
	case int64(domain.Teacher):
		return domain.Teacher
	case int64(domain.Student):
		return domain.Student
	}

	return 0
}
