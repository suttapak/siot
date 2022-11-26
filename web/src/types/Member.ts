import { BoxMemberPermission } from './BoxMemberPermission';
import { UserAccessBox } from './User';

export interface MemberOfBox {
  id: number;
  createdAt: Date;
  updatedAt: Date;
  userAccessToken: string;
  userId: string;
  BoxId: string;
  boxMemberPermission: BoxMemberPermission;
}

export interface Member {
  id: number;
  createdAt: Date;
  updatedAt: Date;
  userAccessToken: string;
  userId: string;
  boxId: string;
  boxMemberPermission: BoxMemberPermission;
  user: UserAccessBox;
}
