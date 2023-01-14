import React from 'react';
import Box from '@mui/material/Box';
import { NavbarComponent } from './Navbar';
import { styled } from '@mui/material/styles';
import Toolbar from '@mui/material/Toolbar';

const DrawerHeader = styled(Toolbar)(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'flex-end',
  padding: theme.spacing(0, 1),

  // necessary for content to be below app bar
}));

interface Props {
  children: JSX.Element;
}

export function Container({ children }: Props): JSX.Element {
  const [open, setOpen] = React.useState<boolean>(false);
  return (
    <>
      <NavbarComponent open={open} setOpen={setOpen}>
        <Box component='main' sx={{ flexGrow: 1, p: 3, backgroundColor: '#D3D3D3', minHeight: '100vh' }}>
          <DrawerHeader variant='dense' />
          {children}
        </Box>
      </NavbarComponent>
    </>
  );
}
