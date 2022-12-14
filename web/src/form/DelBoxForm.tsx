import React from 'react';

import { BsStack } from 'react-icons/bs';
import { DeleteBoxDeliver } from '../delivery/Box';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { ImSpinner9 } from 'react-icons/im';
import { Box } from '../types/Box';

type Props = {
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  box?: Box;
  boxId?: string;
};

export const DelBoxDeliver: React.FC<Props> = (props: Props) => {
  const { boxId, box } = props;

  const queryClient = useQueryClient();

  const { mutate, isLoading } = useMutation(async () => await DeleteBoxDeliver(boxId), {
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['boxes'] });
    },
  });

  return (
    <div>
      <div className='w-full pb-2'>
        <h1 className='text-4xl font-bold text-center py-10'>SIOT</h1>
        <h3 className='text-xl font-bold text-center'>Delete Box `{box?.name}`</h3>
      </div>

      <div className='py-4'>
        <button
          type='button'
          onClick={() => mutate()}
          className='text-lg font-medium flex items-center px-3 py-1 rounded-lg hover:bg-gray-100 transition-colors'
        >
          <BsStack /> <span className='pl-2'>Delete Box</span>{' '}
          {isLoading && (
            <div role='status' className='pl-2'>
              <ImSpinner9 className='mr-2 w-6 h-6 text-gray-200 animate-spin fill-blue-600 dark:to-gray-600' />

              <span className='sr-only'>Loading...</span>
            </div>
          )}
        </button>
      </div>
      {/* error handling */}
    </div>
  );
};
