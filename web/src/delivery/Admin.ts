import { getUserToken } from '../libs';
import { User } from '../types/User';
import { apiClient } from './Api';

export const FindUsers = async () => {
  const res = await apiClient.get<User[]>('/admin/users', {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};

export type AddRoleDto = {
  userId: string;
  role: number;
};

export const AddRoleDeliver = async (body: AddRoleDto) => {
  const res = await apiClient.post<User[]>('/admin/users/roles', JSON.stringify(body), {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};
