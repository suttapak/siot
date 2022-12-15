import { getUserToken } from '../libs';
import { Box } from '../types/Box';
import { apiClient } from './Api';

export interface CreateBoxDto {
  name: string;
  description: string;
}

export interface UpdateBoxDto {
  name?: string;
  description?: string;
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

export const UpdateBoxDeliver = async (box: UpdateBoxDto, boxId?: string) => {
  if (!boxId) {
    return;
  }
  const res = await apiClient.put<Box>(`/boxes/${boxId}`, JSON.stringify(box), {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};

export const DeleteBoxDeliver = async (boxId?: string) => {
  if (!boxId) {
    return;
  }
  const res = await apiClient.delete<Box>(`/boxes/${boxId}`, {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};
