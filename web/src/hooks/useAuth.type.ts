import { LoginDto, RegisterDto } from '../delivery/Auth';
import { AccessToken } from '../types/AccessToken';
import { User } from '../types/User';

export interface authContextInterface {
  signup: ({ email, password }: RegisterDto) => Promise<AccessToken>;
  signin: ({ email, password }: LoginDto) => Promise<AccessToken>;
  signout: () => void;
  token: AccessToken | null;
  user: User | null;
  error: any;
}
