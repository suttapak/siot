import axios from 'axios';
import { getUserToken } from '../libs';
import { AvatarType } from '../types/Avatar';

export interface UpdateAvatarParams {
  file: FileList;
}

export const UpdateAvatarDeliver = async ({ file }: UpdateAvatarParams) => {
  let formdata = new FormData();
  formdata.append('file', file[0]);

  const res = await axios.put<AvatarType>('http://localhost:4000/avatar', formdata, {
    headers: {
      Authorization: `Bearer ${getUserToken()}`,
      'Content-Type': 'multipart/form-data',
    },
  });
  return res.data;
};
