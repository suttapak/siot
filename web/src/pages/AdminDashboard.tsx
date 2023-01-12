import { Container } from '@mui/material';
import React from 'react';
import UserTable from '../components/UserTable';

type Props = {};

const AdminDashboard = (props: Props) => {
  return (
    <Container>
      <UserTable />
    </Container>
  );
};

export default AdminDashboard;
