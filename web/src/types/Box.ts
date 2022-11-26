import { Secret } from './BoxSecret';
import { CanPub } from './CanPub';
import { CanSub } from './CanSub';
import { MemberOfBox } from './Member';

export interface Box {
  id: string;
  createdAt: Date;
  updatedAt: Date;
  name: string;
  description: string;
  ownerId: string;
  members: MemberOfBox[];
  boxSecret: Secret;
  canSub: CanSub;
  canPub: CanPub;
}
