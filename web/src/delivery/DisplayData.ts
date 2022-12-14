import { getUserToken } from '../libs';
import { DataDisplay } from '../types/Data';
import { apiClient } from './Api';

export interface GetDisplayDataParams {
  boxId: string;
  displayId: number;
}
export const GetDisplayData = async (body: GetDisplayDataParams) => {
  const res = await apiClient.get<DataDisplay[]>(`boxes/${body.boxId}/displays/${body.displayId}/data`, {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};
