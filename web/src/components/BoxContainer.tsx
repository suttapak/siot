import * as React from 'react';
import Box from '@mui/material/Box';
import BottomNavigation from '@mui/material/BottomNavigation';
import BottomNavigationAction from '@mui/material/BottomNavigationAction';
import Paper from '@mui/material/Paper';
import { Link, useLocation, useParams } from 'react-router-dom';

import TokenIcon from '@mui/icons-material/Token';
import DashboardIcon from '@mui/icons-material/Dashboard';
import KeyIcon from '@mui/icons-material/Key';
import PeopleIcon from '@mui/icons-material/People';

type Props = {
  children: React.ReactNode;
};

const BoxContainer = (props: Props) => {
  const { children } = props;
  const { boxId } = useParams();
  const path = useLocation();

  React.useEffect(() => {
    if (path.pathname === `/boxes/${boxId}`) {
      setValue(0);
    } else if (path.pathname === `/boxes/${boxId}/dashboard`) {
      setValue(1);
    } else if (path.pathname === `/boxes/${boxId}/secret`) {
      setValue(2);
    } else if (path.pathname === `/boxes/${boxId}/members`) {
      setValue(3);
    }
    // eslint-disable-next-line
  }, []);

  const [value, setValue] = React.useState(0);
  return (
    <React.Fragment>
      <Box sx={{ pb: 7 }}>
        <div>{children}</div>
        <Box sx={{ position: 'fixed', pl: '64px', bottom: 0, left: 0, right: 0, display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
          <Paper
            sx={{
              minWidth: 300,
              width: '90%',
              borderTopLeftRadius: 10,
              borderTopRightRadius: 10,
              overflow: 'hidden',
            }}
            elevation={3}
          >
            <BottomNavigation
              showLabels
              value={value}
              onChange={(event, newValue) => {
                setValue(newValue);
              }}
            >
              <BottomNavigationAction component={Link} to={`/boxes/${boxId}`} label='Box' icon={<TokenIcon />} />
              <BottomNavigationAction label='Dashboard ' component={Link} to={`/boxes/${boxId}/dashboard`} icon={<DashboardIcon />} />
              <BottomNavigationAction label='Secret ' component={Link} to={`/boxes/${boxId}/secret`} icon={<KeyIcon />} />
              <BottomNavigationAction label='Member ' component={Link} to={`/boxes/${boxId}/members`} icon={<PeopleIcon />} />
            </BottomNavigation>
          </Paper>
        </Box>
      </Box>
    </React.Fragment>
  );
};

export default BoxContainer;
