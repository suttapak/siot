import React from 'react';

type Props = {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  children: React.ReactNode;
  position?: 'center' | 'l' | 'r';
};

const Menu = ({ open, children, setOpen, position }: Props) => {
  return (
    <React.Fragment>
      <div onClick={() => setOpen(false)} className={`${open ? 'fixed' : 'hidden'} z-40 bg-opacity-0 top-0 left-0 w-full h-full  py-5 bg-gray-200 `}></div>

      <div className='w-0 h-0 relative z-50'>
        <div
          className={`${open ? '' : 'hidden'} ${
            position ? (position === 'center' ? 'right-0 translate-x-1/2' : position === 'r' ? 'right-0' : 'left-0') : 'right-0'
          } top-0 -right-10 z-30 absolute  min-w-[128px] whitespace-nowrap w-auto rounded-md bg-white border shadow `}
        >
          <ul>{children}</ul>
        </div>
      </div>
    </React.Fragment>
  );
};

export default Menu;
