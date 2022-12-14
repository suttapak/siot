import { getUserToken } from '../libs';
import { Member } from '../types/Member';
import { apiClient } from './Api';

export const GetBoxMembers = async (boxId: string) => {
  const res = await apiClient.get<Member[]>(`boxes/${boxId}/members`, {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};

export const AddMember = async (boxId: string, body: AddMemberDto) => {
  const res = await apiClient.post<Member>(`boxes/${boxId}/members`, JSON.stringify(body), {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};

export type AddMemberDto = {
  userEmail: string;
  canRead: boolean;
  canWrite: boolean;
};
