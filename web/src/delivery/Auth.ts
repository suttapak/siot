import { apiClient } from './Api';
import { AccessToken } from '../types/AccessToken';
import { getUserToken } from '../libs';

export interface LoginDto {
  email: string;
  password: string;
}

export interface RegisterDto {
  firstName?: string | null;
  lastName?: string | null;
  email: string;
  password: string;
}

export const login = async (body: LoginDto): Promise<AccessToken> => {
  const res = await apiClient.post<AccessToken>(
    '/auth/login',
    {
      email: body.email,
      password: body.password,
    },
    {
      headers: {
        Authorization: `Bearer ${getUserToken()}`,
      },
    }
  );
  return res.data;
};

export const register = async (body: RegisterDto): Promise<AccessToken> => {
  const res = await apiClient.post<AccessToken>(
    '/auth/register',
    {
      firstName: body.firstName,
      lastName: body.lastName,
      email: body.email,
      password: body.password,
    },
    {
      headers: {
        Authorization: `Bearer ${getUserToken()}`,
      },
    }
  );
  return res.data;
};
