import { getUserToken } from '../libs';
import { ResponseOkType } from '../types/ResponseOk';
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

export type ChangePasswordDto = {
  password: string;
  newPassword: string;
};

export const ChangePasswordDeliver = async (uId: string, body: ChangePasswordDto) => {
  const res = await apiClient.put<ResponseOkType>(`/user/${uId}/password`, JSON.stringify(body), {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};
