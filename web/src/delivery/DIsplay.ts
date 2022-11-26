import { getUserToken } from '../libs';
import { DisplayType } from '../types/Display';
import { apiClient } from './Api';

export const GetDisplays = async (boxId: string) => {
  const res = await apiClient.get<DisplayType[]>(`boxes/${boxId}/displays`, {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};

export type CreateDisplayDto = {
  name: string;
  key: string;
  description?: string;
  layout: {
    i: string;
    x: number;
    y: number;
    h: number;
    w: number;
  };
  widget: {
    id: number;
  };
};
export const CreateDisplays = async (boxId: string, body: CreateDisplayDto) => {
  const res = await apiClient.post(`boxes/${boxId}/displays`, JSON.stringify(body), {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};
