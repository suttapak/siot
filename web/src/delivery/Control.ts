import { getUserToken } from '../libs';
import { Control } from '../types/Control';
import { apiClient } from './Api';

export const getControls = async (boxId: string) => {
  const res = await apiClient.get<Control[]>(`boxes/${boxId}/controls`, {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};

export type CreateControlDto = {
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
export const CreateControl = async (boxId: string, body: CreateControlDto) => {
  const res = await apiClient.post(`boxes/${boxId}/controls`, JSON.stringify(body), {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};
