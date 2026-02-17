export interface IUserRepository {
  updateProfile: (name: string | null, bio: string | null) => Promise<void>;
  updateIcon: (icon: File | null) => Promise<void>;
}