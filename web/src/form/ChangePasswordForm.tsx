import React from 'react';

import { BsStack } from 'react-icons/bs';
import { useForm } from 'react-hook-form';
import { ChangePasswordDeliver, ChangePasswordDto } from '../delivery/User';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { Alert, Box, Button, CircularProgress, TextField } from '@mui/material';
import { useAuth } from '../hooks';

type Props = {
  uId: string;
};

export const ChangePasswordForm: React.FC<Props> = (props: Props) => {
  const { register, handleSubmit, reset } = useForm<ChangePasswordDto>();
  const { uId } = props;
  const { signout } = useAuth();

  const queryClient = useQueryClient();

  const { mutate, isLoading, error } = useMutation(async (body: ChangePasswordDto) => await ChangePasswordDeliver(uId, body), {
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['boxes'] });
    },
  });

  return (
    <form
      onSubmit={handleSubmit((v) => {
        mutate(v, {
          onSuccess: () => {
            reset();
            signout();
          },
        });
      })}
    >
      {/* body */}
      <Box paddingY={0.1}>
        <TextField type={'text'} label='รหัสผ่านเดิม' variant='standard' {...register('password', { required: true })} fullWidth />
      </Box>
      <Box paddingY={0.1}>
        <TextField type={'text'} label='รหัสผ่านใหม่' variant='standard' {...register('newPassword', { required: true })} fullWidth />
      </Box>
      <Box paddingY={1}>
        <Alert severity='warning'>ตอนนี้ยังไม่มีระบบการกู้รหัสผ่าน โปรดจดจำรหัสผ่าน!</Alert>
      </Box>
      <div className='py-4'>
        <Button type='submit'>
          <BsStack /> <span className='pl-2'>ยืนยันการเปลี่ยนรหัสผ่าน</span>{' '}
          {isLoading && (
            <div role='status' className='pl-2'>
              <CircularProgress size={22} />
              <span className='sr-only'>Loading...</span>
            </div>
          )}
        </Button>
        {Boolean(error) && <Alert severity='error'>{'มีบางอย่างไม่ถูกต้อง'}</Alert>}
      </div>
      {/* error handling */}
    </form>
  );
};
