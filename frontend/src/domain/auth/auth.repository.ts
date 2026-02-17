import type { User } from "../user/user.model";

export interface IAuthRepository {
  login(userId: string): Promise<User>;
  logout(): Promise<void>;
  fetchCurrentSession(): Promise<User | null>;
}