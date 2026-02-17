import axios from "axios";
import { apiClient } from "../../api/client";
import type { AuthSession } from "../../domain/auth/auth.model";
import type { IAuthRepository } from "../../domain/auth/auth.repository";
import type { User } from "../../domain/user/user.model";

export class AuthRepositoryImpl implements IAuthRepository {
  async login(_name: string, _password: string): Promise<User> {
    // NOTE: 現在はBackendにDummyLoginがあるので引数は使わない
    const { data } = await apiClient.post<User>("/users/login");
    return data;
  }

  async logout(): Promise<void> {
    await apiClient.post("/users/logout");
  }

  async fetchCurrentSession(): Promise<AuthSession> {
    try {
      const { data } = await apiClient.get<User>("/users/me");
      return {
        status: "authenticated",
        user: data,
      };
    } catch (error) {
      if (axios.isAxiosError(error)) {
        if (error.response?.status === 401) {
          return { status: "unauthenticated", user: null };
        }
      }

      console.error("Unexpected error during fetching session:", error);
      throw error;
    }
  }
}
