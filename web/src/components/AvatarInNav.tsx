import React from 'react';
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
import { useAuth } from '../hooks';
import { Link } from 'react-router-dom';

type Props = {};

const AvatarInNav = (props: Props) => {
  const auth = useAuth();

  const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(null);

  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  return (
    <>
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
    </>
  );
};

export default AvatarInNav;
