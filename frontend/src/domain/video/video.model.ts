import type { UserId } from "../user/user.model";

export type VideoId = string;

export type VideoStatus =
  | 'initial'
  | 'uploaded'
  | 'queued'
  | 'processing'
  | 'ready'
  | 'failed';

export type Tag =
  | 'competition_programing'
  | 'web_development'
  | 'machine_learning'
  | 'game_development'
  | 'infrastructure'
  | 'other';

export type Visability = 'public' | 'private';

export interface Video {
  readonly id: VideoId;
  readonly ownerId?: UserId;
  readonly status: VideoStatus;
  readonly title: string;
  readonly description: string;
  readonly tags: Tag[];
  readonly failureReason?: string;
  readonly visability?: Visability;
  readonly createdAt: Date;
}