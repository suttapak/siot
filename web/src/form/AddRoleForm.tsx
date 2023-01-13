import React from 'react';

import { BsStack } from 'react-icons/bs';
import { useForm } from 'react-hook-form';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { Alert, Box, Button, CircularProgress, FormControl, InputLabel, MenuItem, Select, TextField } from '@mui/material';
import { AddRoleDeliver, AddRoleDto } from '../delivery/Admin';

type Props = {};

const AddRoleForm = (props: Props) => {
  const { register, handleSubmit, reset } = useForm<AddRoleDto>();

  const queryClient = useQueryClient();

  const { mutate, isLoading, error } = useMutation(async (body: AddRoleDto) => await AddRoleDeliver(body), {
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['users'] });
    },
  });

  return (
    <>
      <form
        onSubmit={handleSubmit((v) => {
          mutate(v, {
            onSuccess: () => {
              reset();
            },
          });
        })}
      >
        {/* body */}
        <Box>
          <Box paddingY={1} flexGrow={1}>
            <TextField
              type={'text'}
              label='รหัสผู้ใช้งาน'
              placeholder='30a5f73d-c40e-43ee-b428-f4ac8e4377b2'
              variant='standard'
              {...register('userId', { required: true })}
              fullWidth
            />
          </Box>
          <Box paddingY={1} flexGrow={1}>
            <FormControl fullWidth>
              <InputLabel id='demo-simple-select-label'>สิทธิ</InputLabel>
              <Select variant='standard' labelId='demo-simple-select-label' {...register('role', { required: true })} id='demo-simple-select' label='Age'>
                <MenuItem value={2}>Admin</MenuItem>
                <MenuItem value={3}>SuperAdmin</MenuItem>
              </Select>
            </FormControl>
          </Box>
          <Box paddingY={1} flexGrow={1}>
            <Alert severity='info'>เมื่อเพิ่มแล้วจะไม่สามารถปรับเปลี่ยนสิทธิผู้ใช้งานได้!</Alert>
          </Box>

          <Box paddingY={1} flexGrow={1}>
            <Button type='submit'>
              <BsStack /> <span className='pl-2'>ยืนยัน</span>{' '}
              {isLoading && (
                <div role='status' className='pl-2'>
                  <CircularProgress size={22} />
                  <span className='sr-only'>Loading...</span>
                </div>
              )}
            </Button>
            {Boolean(error) && <Alert severity='error'>{'มีบางอย่างไม่ถูกต้อง'}</Alert>}
          </Box>
        </Box>
        {/* error handling */}
      </form>
    </>
  );
};

export default AddRoleForm;
