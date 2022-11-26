import { Notification } from './Notification';

export interface Setting {
  id: number;
  createdAt: Date;
  updatedAt: Date;
  userId: string;
  notification: Notification;
}
