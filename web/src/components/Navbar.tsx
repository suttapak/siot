import * as React from 'react';
import Box from '@mui/material/Box';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import Avatar from '@mui/material/Avatar';
import Tooltip from '@mui/material/Tooltip';
import MenuItem from '@mui/material/MenuItem';
import PersonAdd from '@mui/icons-material/PersonAdd';
import ListItemIcon from '@mui/material/ListItemIcon';
import Logout from '@mui/icons-material/Logout';

import { Link } from 'react-router-dom';
import { useAuth } from '../hooks';

import { CgMenuGridO } from 'react-icons/cg';
import { GrClose } from 'react-icons/gr';
import TokenIcon from '@mui/icons-material/Token';
import LibraryBooksIcon from '@mui/icons-material/LibraryBooks';
import ArticleIcon from '@mui/icons-material/Article';
import { AdminPanelSettings } from '@mui/icons-material';

interface ItemProps {
  path: string;
  children: React.ReactNode;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

function Items({ path, children, setOpen }: ItemProps): JSX.Element {
  return (
    <>
      <li>
        <Link
          onClick={() => setOpen(false)}
          to={path}
          className='flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700'
        >
          {children}
        </Link>
      </li>
    </>
  );
}

interface NavbarComponentProps {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

export function NavbarComponent({ open, setOpen }: NavbarComponentProps) {
  const auth = useAuth();

  const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(null);

  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  return (
    <React.Fragment>
      {/* <!-- drawer init and show --> */}
      <div className='fixed top-0 left-0 z-40 w-full h-10 bg-white  flex justify-center items-center'>
        <div className='w-full px-3  mx-auto flex justify-center items-center'>
          <button onClick={() => setOpen(!open)} className='text-gray-800 font-bold text-lg ' type='button'>
            <CgMenuGridO size={30} color='gray' />
          </button>
          <div className='w-full mx-auto flex justify-start'>
            <Link to={'/'} className=''>
              <h1 className='text-gray-800 pl-4 font-bold font-sans text-lg md:text-2xl'>SIOT</h1>
            </Link>
          </div>
          {auth.user && (
            <React.Fragment>
              <Menu
                sx={{ mt: '45px' }}
                id='menu-appbar'
                anchorEl={anchorElUser}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                open={Boolean(anchorElUser)}
                onClose={handleCloseUserMenu}
              >
                <Link to={`/profile/${auth.user.id}`}>
                  <MenuItem onClick={handleCloseUserMenu}>
                    <ListItemIcon>
                      <PersonAdd fontSize='small' />
                    </ListItemIcon>
                    <Typography textAlign='center'>Profile account</Typography>
                  </MenuItem>
                </Link>
                {auth.user && (
                  <MenuItem
                    onClick={() => {
                      handleCloseUserMenu();
                      auth.signout();
                    }}
                  >
                    <ListItemIcon>
                      <Logout fontSize='small' />
                    </ListItemIcon>
                    Logout
                  </MenuItem>
                )}
              </Menu>
              <Box>
                <Tooltip title='Open settings'>
                  <IconButton onClick={handleOpenUserMenu}>
                    <Avatar
                      sx={{ width: 35, height: 35 }}
                      src={process.env.REACT_APP_SERVER_URL_PATH + auth.user.avatar.url.substring(1)}
                      alt={auth.user.avatar.title}
                    />
                  </IconButton>
                </Tooltip>
              </Box>
            </React.Fragment>
          )}
        </div>
      </div>

      {/* <!-- drawer component --> */}
      <div
        className={`fixed border-r z-40 h-screen p-4 overflow-y-auto bg-white w-64 dark:bg-gray-800 transition-transform ${
          open ? 'translate-x-0' : '-translate-x-full'
        } `}
        aria-labelledby='drawer-navigation-label'
      >
        <h5 className='text-base font-semibold text-gray-500 uppercase dark:text-gray-400'>Menu</h5>
        <button
          onClick={() => setOpen(false)}
          type='button'
          data-drawer-dismiss='drawer-navigation'
          aria-controls='drawer-navigation'
          className='text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 absolute top-2.5 right-2.5 inline-flex items-center dark:hover:bg-gray-600 dark:hover:text-white'
        >
          <GrClose />
          <span className='sr-only'>Close menu</span>
        </button>
        <div className='py-4 overflow-y-auto'>
          <ul className='space-y-2 flex-col'>
            {auth.user && (
              <>
                {auth.user?.roles.filter((e) => e.name === 'admin' || e.name === 'superAdmin').length > 0 && (
                  <Items path='/admin/dashboard' setOpen={setOpen}>
                    <AdminPanelSettings fontSize={'small'} />
                    <span className='ml-3'>Admin dashboard</span>
                  </Items>
                )}
              </>
            )}
            <Items path='/' setOpen={setOpen}>
              <TokenIcon fontSize={'small'} />
              <span className='ml-3'>Home</span>
            </Items>
            <li>
              <a
                onClick={() => setOpen(false)}
                href={'https://document.rocket-translate.com/suttapak/siot.h'}
                className='flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700'
              >
                <ArticleIcon fontSize={'small'} />
                <span className='ml-3'>Document</span>
              </a>
            </li>
            <li>
              <a
                onClick={() => setOpen(false)}
                href={'https://github.com/suttapak/siot.h'}
                className='flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700'
              >
                <LibraryBooksIcon fontSize={'small'} />
                <span className='ml-3'>Library</span>
              </a>
            </li>
          </ul>
        </div>
      </div>
      <div onClick={() => setOpen(false)} className={`${open ? 'block' : 'hidden'} bg-gray-900 bg-opacity-50 dark:bg-opacity-80 fixed inset-0 z-30`}></div>
    </React.Fragment>
  );
}
