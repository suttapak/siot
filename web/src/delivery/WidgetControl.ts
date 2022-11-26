import { getUserToken } from '../libs';
import { Widget } from '../types/Widget';
import { apiClient } from './Api';

export const getWidgetControl = async () => {
  const res = await apiClient.get<Widget[]>('widgets/controls', { headers: { Authorization: `Bearer ${getUserToken()}` } });
  return res.data;
};
