import React from 'react';

import { BsStack } from 'react-icons/bs';
import { useForm } from 'react-hook-form';
import { useMutation } from '@tanstack/react-query';
import { ImSpinner9 } from 'react-icons/im';
import { UpdateAvatarDeliver, UpdateAvatarParams } from '../delivery/Avatar';
import { useAuth } from '../hooks';

type Props = {
  setOpen: React.Dispatch<React.SetStateAction<any>>;
};

export const UpdateAvatarForm: React.FC<Props> = (props: Props) => {
  const { register, handleSubmit, reset } = useForm<UpdateAvatarParams>();

  const { mutate, isLoading } = useMutation(UpdateAvatarDeliver);

  const { reUser } = useAuth();

  return (
    <form
      onSubmit={handleSubmit((v) => {
        mutate(v, {
          onSuccess: () => {
            reset();
            reUser();
            props.setOpen(false);
          },
        });
      })}
    >
      {/* title */}
      <div className='w-full pb-2'>
        <h1 className='text-4xl font-bold text-center py-10'>SIOT</h1>
        <h3 className='text-xl font-bold text-center'>Update Avatar</h3>
      </div>
      <hr />
      {/* body */}

      <div className='py-2'>
        <label className='block text-gray-700 text-xs font-bold mb-1'>Select image</label>
        <input
          className='form-control block w-full px-2 py-1 text-md font-normal text-gray-700 bg-white  border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none'
          type='file'
          {...register('file')}
        />
      </div>
      <div className='py-4'>
        <button type='submit' className='text-lg font-bold flex items-center px-3 py-1 rounded-lg hover:bg-gray-100 transition-colors'>
          <BsStack /> <span className='pl-2'>Create Box</span>{' '}
          {isLoading && (
            <div role='status' className='pl-2'>
              <ImSpinner9 className='mr-2 w-6 h-6 text-gray-200 animate-spin fill-blue-600 dark:to-gray-600' />

              <span className='sr-only'>Loading...</span>
            </div>
          )}
        </button>
      </div>
      {/* error handling */}
    </form>
  );
};
