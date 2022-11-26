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
}

export interface UserAccessBox {
  id: string;
  createdAt: Date;
  updatedAt: Date;
  firstName: string;
  lastName: string;
  email: string;
}
