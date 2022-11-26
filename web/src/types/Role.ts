export interface Role {
  id: number;
  createdAt: Date;
  updatedAt: Date;
  name: string;
  permissionState: number;
  displayName: string;
  description: string;
}
