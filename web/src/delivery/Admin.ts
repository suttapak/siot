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
