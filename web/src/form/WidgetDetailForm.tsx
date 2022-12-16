import React from 'react';

import { BsStack } from 'react-icons/bs';
import { useForm } from 'react-hook-form';
import { useMutation } from '@tanstack/react-query';
import { ImSpinner9 } from 'react-icons/im';
import { useAuth } from '../hooks';
import { UpdateWidgetDeliverApi, UpdateWidgetDeliverType } from '../delivery/UpdateWidgetDeliver';
import { DisplayType } from '../types/Display';
import { Control } from '../types/Control';

type Props = {
  widget?: DisplayType | Control;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
};

export const UpdateWidgetForm: React.FC<Props> = (props: Props) => {
  const { widget } = props;
  const { register, handleSubmit, reset } = useForm<UpdateWidgetDeliverType>();

  const { mutate, isLoading } = useMutation(UpdateWidgetDeliverApi);

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
        <h3 className='text-xl font-bold text-center'>Update Widget</h3>
      </div>
      <hr />
      {/* body */}

      <div className='py-2'>
        <label className='block text-gray-700 text-xs font-bold mb-1'>Name</label>
        <input
          defaultValue={widget?.name}
          className='form-control block w-full px-2 py-1 text-md font-normal text-gray-700 bg-white  border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none'
          type='text'
          placeholder='Smart farm at Ubon'
          id='name'
          {...(register('name'), { required: true })}
        />
      </div>
      <div className='py-2'>
        <label className='block text-gray-700 text-xs font-bold mb-1'>Name</label>
        <input
          defaultValue={widget?.key}
          className='form-control block w-full px-2 py-1 text-md font-normal text-gray-700 bg-white  border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none'
          type='text'
          placeholder='Smart farm at Ubon'
          id='name'
          {...(register('key'), { required: true })}
        />
      </div>
      <div className='py-2'>
        <label className='block text-gray-700 text-xs font-bold mb-1'>Name</label>
        <input
          defaultValue={widget?.description}
          className='form-control block w-full px-2 py-1 text-md font-normal text-gray-700 bg-white  border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none'
          type='text'
          placeholder='Smart farm at Ubon'
          id='name'
          {...(register('description'), { required: true })}
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
