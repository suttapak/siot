import { useQuery } from '@tanstack/react-query';
import { getUserById } from '../delivery/User';
import { Box } from '../types/Box';
import { HiDotsVertical } from 'react-icons/hi';
import { Link } from 'react-router-dom';
import { Modal } from './Modal';
import BoxPopupDetail from './BoxPopupDetail';
import React from 'react';
import { UpdateBoxPopupDetail } from '../form/BoxPopupDetail';

export const BoxCard = ({ box }: { box: Box }) => {
  const { data } = useQuery(['user'], async () => await getUserById(box.ownerId));

  const [open, setOpen] = React.useState(false);
  const [openEdit, setOpenEdit] = React.useState(false);

  return (
    <React.Fragment>
      <div onClick={() => setOpen(false)} className={`${open ? '' : 'hidden'} z-10 bg-opacity-0 top-0 left-0 w-full h-full  absolute py-5 bg-gray-200 `}></div>

      <div className='w-full  h-32 rounded-lg  relative  col-span-1'>
        <div className='w-full px-4 h-24 bg-neutral-700 relative rounded-t-lg'>
          <Link to={`/boxes/${box.id}`} className='hover:underline decoration-gray-50'>
            <h2 className='text-xl pt-2 text-gray-50'>{box.name}</h2>
          </Link>
          <p className='pt-1 text-xs text-gray-100'>สมาชิก {box.members.length}</p>
          <h5 className='text-xs pt-5 text-gray-100 '>
            {data?.firstName} {data?.lastName}
          </h5>
        </div>
        {/* logo  */}
        <div className='absolute z-20 top-1 right-1 cursor-pointer'>
          <div
            onClick={() => setOpen(!open)}
            className={`${
              open ? 'bg-gray-500 bg-opacity-70' : ''
            } w-10 h-10 rounded-full hover:bg-gray-500 hover:bg-opacity-70 flex justify-center items-center`}
          >
            <HiDotsVertical className='text-2xl text-gray-50' />
          </div>
          <div className='w-0 h-0 relative'>
            <div className={`${open ? '' : 'hidden'} top-0 z-30 absolute py-1.5 min-w-[128px] rounded-lg bg-white border shadow`}>
              <div
                onClick={() => {
                  setOpenEdit(true);
                  setOpen(false);
                }}
                className='px-3 w-full py-2  text-base hover:bg-gray-200'
              >
                <span>Edit</span>
              </div>
              <div className='px-3 w-full py-2  text-base hover:bg-gray-200'>
                <span>Details</span>
              </div>
            </div>
          </div>
        </div>

        {/* avatar */}
        <div className='w-16 h-16 bg-white p-1 rounded-full overflow-hidden absolute bottom-0 right-5 flex justify-center items-center'>
          <div className='w-full h-full rounded-full overflow-hidden'>
            <img src={process.env.REACT_APP_SERVER_URL + String(data?.avatar.url.substring(1))} alt={data?.avatar.title} className='w-full ' />
          </div>
        </div>
      </div>
      <Modal open={openEdit} setOpen={setOpenEdit}>
        <UpdateBoxPopupDetail setOpen={setOpenEdit} />
      </Modal>
    </React.Fragment>
  );
};
