import { getUserToken } from '../libs';
import { Box } from '../types/Box';
import { apiClient } from './Api';

export interface CreateBoxDto {
  name: string;
  description: string;
}

export const FindBoxes = async () => {
  const res = await apiClient.get<Box[]>('/boxes', {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};

export const FindBox = async (boxId: string) => {
  const res = await apiClient.get<Box>(`/boxes/${boxId}`, {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};

export const CreateBox = async (box: CreateBoxDto) => {
  const res = await apiClient.post<Box>('/boxes', JSON.stringify(box), {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};
