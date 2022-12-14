import React from 'react';
import { UpdateBoxPopupDetail } from '../form/BoxPopupDetail';
import { Box } from '../types/Box';

type Props = {
  box: Box;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
};

const BoxPopupDetail = ({ box, setOpen }: Props) => {
  return (
    <div className='w-0 h-0 relative'>
      {/* title  */}
      <div className='w-full py-2'>
        <h2 className='text-center text-lg'>{box.name}</h2>
      </div>
      {/* body */}
      <div className='w-full'>
        <UpdateBoxPopupDetail setOpen={setOpen} />
      </div>
    </div>
  );
};

export default BoxPopupDetail;
