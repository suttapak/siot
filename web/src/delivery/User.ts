import { getUserToken } from '../libs';
import { User } from '../types/User';
import { apiClient } from './Api';

export const getUser = async () => {
  return await (
    await apiClient.get<User>('/user', {
      headers: {
        Authorization: `Bearer ${getUserToken()}`,
      },
    })
  ).data;
};

export const getUserById = async (id: string) => {
  const res = await apiClient.get<User>('/user/' + id, {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};
