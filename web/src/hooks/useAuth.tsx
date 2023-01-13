import React from 'react';
import { Navigate, useLocation, useNavigate } from 'react-router-dom';
import { login, LoginDto, register, RegisterDto } from '../delivery/Auth';
import { getUser } from '../delivery/User';
import { AccessToken } from '../types/AccessToken';
import { User } from '../types/User';
import { authContextInterface } from './index';
import { useLocalStorage } from './index';

const authContext = React.createContext<authContextInterface>({} as authContextInterface);

interface Props {
  children: React.ReactNode;
}
export const ProviderAuth: React.FC<Props> = ({ children }) => {
  const auth = useProviderAuth();
  return <authContext.Provider value={auth}>{children}</authContext.Provider>;
};

export const useAuth = () => {
  return React.useContext(authContext);
};

const useProviderAuth: () => authContextInterface = () => {
  const [user, setUser] = useLocalStorage<User>('user', null);
  const [token, setToken] = useLocalStorage<AccessToken>('accessToken', null);
  const [error] = React.useState<any>(null);

  const navigate = useNavigate();

  //load user form local storage.
  React.useEffect(() => {
    if (typeof window.localStorage.getItem('user') === 'string') {
      const user = JSON.parse(String(window.localStorage.getItem('user')));
      const accessToken = JSON.parse(String(window.localStorage.getItem('accessToken')));
      setUser(user);
      setToken(accessToken);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const signup = async (body: RegisterDto) => {
    return await register(body);
  };

  const reUser = async () => {
    const u = await getUser();
    setUser(u);
    return u;
  };

  const signin = async (body: LoginDto) => {
    const res = await login(body);
    setToken(res);
    const u = await getUser();
    setUser(u);
    return res;
  };

  const signout: () => void = () => {
    setUser(null);
    setToken(null);
    navigate('/signin', { replace: true });
  };
  const value = React.useMemo(() => {
    return { signup, signin, signout, user, token, error, reUser };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [user, token, error]);
  return value;
};

interface RequireAuthProps {
  children: JSX.Element;
}
export const RequireAuth: React.FC<RequireAuthProps> = ({ children }) => {
  const auth = useAuth();
  const location = useLocation();

  if (!auth.user) {
    return <Navigate to={'/signin'} state={{ form: location }} replace={true} />;
  }

  return children;
};

export const RequireAdminOrSuperAdmin: React.FC<RequireAuthProps> = ({ children }) => {
  const auth = useAuth();
  const location = useLocation();

  let adminOrSuperAdmin = false;
  if (!auth.user) {
    return <Navigate to={'/signin'} state={{ form: location }} replace={true} />;
  }

  for (let r of auth.user.roles) {
    if (r.name.toLowerCase() === 'admin' || r.name.toLowerCase() === 'superadmin') {
      adminOrSuperAdmin = true;
    }
  }

  if (!adminOrSuperAdmin) {
    return <Navigate to={'/signin'} state={{ form: location }} replace={true} />;
  }

  return children;
};

export const NotAuth: React.FC<RequireAuthProps> = ({ children }) => {
  const auth = useAuth();
  const location = useLocation();

  if (auth.user) {
    return <Navigate to={'/'} state={{ form: location }} replace={true} />;
  }

  return children;
};
