import { AccessToken } from '../types/AccessToken';

export const getUserToken = () => {
  const token: AccessToken | null = JSON.parse(String(window.localStorage.getItem('accessToken')));

  return token?.accessToken;
};
