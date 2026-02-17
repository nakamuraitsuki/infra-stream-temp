import { apiClient } from "../../api/client";
import type { IUserRepository } from "../../domain/user/user.repository";

export class UserRepositoryImpl implements IUserRepository {
  async updateProfile(name: string | null, bio: string | null): Promise<void> {
    await apiClient.patch("/users/profile", {
      name,
      bio,
    });
  }

  async updateIcon(icon: File | null): Promise<void> {
    if (!icon) {
      throw new Error("Icon file is required");
    }

    const formData = new FormData();
    formData.append("file", icon);

    await apiClient.post("/users/icon", formData);
  }
}