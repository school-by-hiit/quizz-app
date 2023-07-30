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

import { defineStore } from "pinia";
import { googleLogout } from "vue3-google-login";

import { getApi } from "~/api/common-api";
import { User, UserRoleEnum } from "~/api/model";

export type AuthState = {
  userId?: string;
  userName?: string;
  userEmail?: string;
  exp: number;
  userRole?: UserRoleEnum;
  picture?: string;
  jwtToken?: string;
  logged: boolean;
};

export const useAuthStore = defineStore("auth", {
  state: () =>
    ({
      userId: undefined,
      userName: undefined,
      userEmail: undefined,
      exp: 0,
      userRole: undefined,
      picture: undefined,
      jwtToken: undefined,
      logged: false,
    }) as AuthState,
  getters: {
    isLogged({ logged, exp }: AuthState): boolean {
      return logged && !(new Date().getTime() > exp);
    },
    getPicture({ picture }: AuthState): string | undefined {
      return picture;
    },
    getUsername({ userName }: AuthState): string | undefined {
      return userName;
    },
    getUserEmail({ userEmail }: AuthState): string | undefined {
      return userEmail;
    },
    getUserRole({ userRole }: AuthState): string | undefined {
      return userRole;
    },
    hasExpired({ exp }: AuthState): boolean {
      return new Date().getTime() > exp;
    },
  },
  actions: {
    login(token: string, picture: string, exp: number) {
      this.jwtToken = token;
      this.picture = picture;
      this.exp = exp * 1000;
      getApi()
        .post<User>(`/api/v1/login`)
        .then(({ data }) => {
          this.logged = true;
          this.userId = data.id;
          this.userName = data.firstname + " " + data.lastname;
          this.userEmail = data.email;
          this.userRole = data.role;
        });
    },
    logout() {
      googleLogout();
      this.logged = false;
    },
  },
  persist: {
    storage: sessionStorage,
  },
});
