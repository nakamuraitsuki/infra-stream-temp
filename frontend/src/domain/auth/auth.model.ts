import type { User } from "../user/user.model";

export interface AuthSession {
  readonly isAuthenticated: boolean;
  readonly user: User | null;
}