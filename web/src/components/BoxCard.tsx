import { useQuery } from '@tanstack/react-query';
import { getUserById } from '../delivery/User';
import { Box } from '../types/Box';
import { HiDotsVertical } from 'react-icons/hi';

export const BoxCard = ({ box }: { box: Box }) => {
  const { data } = useQuery(['user'], async () => await getUserById(box.ownerId));

  return (
    <div className='w-full  h-32 rounded-lg overflow-hidden relative  col-span-1'>
      <div className='w-full px-4 h-24 bg-blue-400 relative'>
        <h2 className='text-xl font-bold pt-2'>{box.name}</h2>
        <h5 className='text-base'>
          {data?.firstName} {data?.lastName}
        </h5>
        <p className='pt-4 text-xs'>Member {box.members.length}</p>
      </div>
      {/* logo  */}
      <div className='absolute top-2 right-0'>
        <HiDotsVertical className='text-2xl' />
      </div>
      {/* avatar */}
      <div className='w-20 h-20 bg-white p-1 rounded-full overflow-hidden absolute bottom-0 right-10 flex justify-center items-center'>
        <div className='w-full h-full rounded-full overflow-hidden'>
          <img src='./mock-avatar.png' alt='avatar' className='w-full ' />
        </div>
      </div>
    </div>
  );
};
