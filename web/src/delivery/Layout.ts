import { Layout } from 'react-grid-layout';
import { getUserToken } from '../libs';
import { apiClient } from './Api';

export type UpdateLayoutDto = {
  id: number;
  i: string;
  x: number;
  y: number;
  w: number;
  h: number;
};

export const ParseLayoutDto = (layout: Layout[]) => {
  let l: UpdateLayoutDto[] = layout.map((v) => {
    let i: UpdateLayoutDto = {
      id: Number(v.i),
      i: v.i,
      x: v.x,
      y: v.y,
      w: v.w,
      h: v.h,
    };
    return i;
  });
  return l;
};

export const UpdateLayoutDeliver = async (boxId: string, body: UpdateLayoutDto[]) => {
  const res = await apiClient.put(`/boxes/${boxId}/layout`, JSON.stringify(body), {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
};
