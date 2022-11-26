import { getUserToken } from '../libs';
import { Widget } from '../types/Widget';
import { apiClient } from './Api';

export const getWidgetDisplays = async () => {
  const res = await apiClient.get<Widget[]>('widgets/displays', {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
    },
  });
  return res.data;
};
