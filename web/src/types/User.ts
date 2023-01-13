import { AvatarType } from './Avatar';
import { Role } from './Role';
import { Setting } from './Setting';

export interface User {
  id: string;
  createdAt: Date;
  updatedAt: Date;
  firstName: string;
  lastName: string;
  email: string;
  settingId: number;
  roles: Role[];
  box: any[];
  setting: Setting;
  avatar: AvatarType;
}

export interface UserAccessBox {
  id: string;
  createdAt: Date;
  updatedAt: Date;
  firstName: string;
  lastName: string;
  email: string;
}
