import { Container } from '@mui/material';
import React from 'react';
import UserTable from '../components/UserTable';
import AddRole from '../components/AddRole';
import Grid from '@mui/material/Grid';
import { useAuth } from '../hooks';

type Props = {};

const AdminDashboard = (props: Props) => {
  const { user } = useAuth();

  return (
    <Container
      sx={{
        paddingBottom: '2.2rem',
      }}
    >
      <UserTable />
      <Grid container spacing={1}>
        {user && (
          <>
            {user.roles.find((v) => v.name.toLowerCase() == 'superadmin') && (
              <Grid item xs={6}>
                <AddRole />
              </Grid>
            )}
          </>
        )}
      </Grid>
    </Container>
  );
};

export default AdminDashboard;
