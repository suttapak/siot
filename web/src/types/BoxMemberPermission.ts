export interface BoxMemberPermission {
  id: number;
  createdAt: Date;
  updatedAt: Date;
  canRead: boolean;
  canWrite: boolean;
  boxMemberId: number;
}
