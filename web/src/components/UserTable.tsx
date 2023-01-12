import { Box, Paper, Typography } from '@mui/material';
import React from 'react';
import { DataGrid, GridColDef, GridValueGetterParams } from '@mui/x-data-grid';
import { useQuery } from '@tanstack/react-query';
import { FindUsers } from '../delivery/Admin';
import { Role } from '../types/Role';

const columns: GridColDef[] = [
  { field: 'id', headerName: 'ID', width: 290 },
  {
    field: 'firstName',
    headerName: 'ชื่อ',
    width: 150,
  },
  {
    field: 'lastName',
    headerName: 'นามสกุล',
    width: 150,
  },
  {
    field: 'email',
    headerName: 'อีเมล',
    width: 290,
  },
  {
    field: 'roles',
    headerName: 'สิทธิ',
    description: 'This column has a value getter and is not sortable.',
    sortable: false,
    width: 200,
    valueGetter: (params: GridValueGetterParams) => `${params.row.roles.map((v: Role) => v.displayName)}`,
  },
];

type Props = {};

const UserTable = (props: Props) => {
  const { data } = useQuery(['users'], FindUsers);

  return (
    <div>
      <Paper elevation={0} sx={{ padding: 1, marginBottom: '1rem' }}>
        <Box>
          <Typography variant='h6' component={'h6'}>
            ผู้ใช้งานทั้งหมด
          </Typography>
        </Box>
      </Paper>
      <Paper elevation={0}>
        <Box paddingX={1} paddingY={1} sx={{ height: 400, width: '100%' }}>
          {data && (
            <DataGrid
              rows={data}
              columns={columns}
              pageSize={5}
              rowsPerPageOptions={[5]}
              checkboxSelection
              disableSelectionOnClick
              experimentalFeatures={{ newEditingApi: true }}
            />
          )}
        </Box>
      </Paper>
    </div>
  );
};

export default UserTable;
