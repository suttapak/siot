import { Box, Paper, Typography } from '@mui/material';
import React from 'react';
import AddRoleForm from '../form/AddRoleForm';

type Props = {};

const AddRole = (props: Props) => {
  return (
    <React.Fragment>
      <Paper
        elevation={0}
        sx={{
          height: 360,
          marginTop: 2,
        }}
      >
        <Box
          sx={{
            padding: 1,
          }}
        >
          <Typography variant='h6'>เพิ่มสิทธิให้ผู้ใช้งาน</Typography>
        </Box>
        <Box
          sx={{
            padding: 1,
          }}
        >
          <AddRoleForm />
        </Box>
      </Paper>
    </React.Fragment>
  );
};

export default AddRole;
