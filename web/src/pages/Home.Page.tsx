import React from 'react';
import { useQuery } from '@tanstack/react-query';
import { Link } from 'react-router-dom';
import { useToast } from '../hooks/useToast';
import { Box } from '../types/Box';
import { FindBoxes } from '../delivery/Box';
import { BoxCard } from '../components/BoxCard';
import { BsPlusLg } from 'react-icons/bs';
import { Modal } from '../components/Modal';
import { CreateBoxComponent } from '../form/Box';

export function HomePage() {
  const { error, data } = useQuery<Box[], { message: string }>(['boxes'], FindBoxes);

  const toast = useToast();

  if (error) toast.addMessage(error.message);

  const [open, setOpen] = React.useState(false);

  return (
    <>
      <div className='container mx-auto bg-white rounded-lg px-6 '>
        {/* title */}
        <div className='py-2 my-2 border-b-teal-200 border-b flex justify-between '>
          <h3 className='text-2xl font-bold '>Your Boxes</h3>
          <button type='button' onClick={() => setOpen(true)}>
            <div className='rounded-full p-2 overflow-hidden w-10 h-10 flex justify-center items-center hover:bg-gray-200'>
              <BsPlusLg />
            </div>
          </button>
        </div>
        <div className='grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-3 '>
          {data?.map((v) => {
            return (
              <Link key={v.id} to={`/boxes/${v.id}`}>
                <BoxCard box={v} />
              </Link>
            );
          })}
        </div>
      </div>
      {/* modal */}
      <Modal open={open} setOpen={setOpen}>
        <CreateBoxComponent setOpen={setOpen} />
      </Modal>
    </>
  );
}
