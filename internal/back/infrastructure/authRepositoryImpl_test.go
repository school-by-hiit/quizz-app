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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthDBRepository_FindTokenByTokenStr(t *testing.T) {

	connection := getDBConnection(t, true)
	defer connection.Close()

	r := NewAuthRepository()

	token, err := r.FindTokenByTokenStr("42")
	if err != nil {
		assert.Failf(t, "Fail to get token", "%v", err)
	}

	assert.Nil(t, token)
}
