import React from 'react';
import { Control } from '../types/Control';
import { DisplayType } from '../types/Display';

type Props = {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  widget?: Control | DisplayType;
};

const NameKeyWidget = ({ open, setOpen, widget }: Props) => {
  return (
    <button
      onMouseEnter={() => setOpen(true)}
      onMouseLeave={() => setOpen(false)}
      type='button'
      className={`rounded-md ${open ? 'w-auto px-2' : 'w-7'} z-50 h-7 flex text-gray-500 justify-center items-center hover:bg-gray-300 absolute top-1 right-1`}
    >
      {open ? widget?.name : widget?.key.toUpperCase()}
    </button>
  );
};

export default NameKeyWidget;
