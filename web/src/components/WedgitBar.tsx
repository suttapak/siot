import React from 'react';
import { Control } from './control';
import Display from './display';

type Props = {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  setWidgetId: React.Dispatch<React.SetStateAction<number>>;
  modeControl: boolean;
  setModeControl: React.Dispatch<React.SetStateAction<boolean>>;
};

export const WedgitBar = (props: Props) => {
  const { modeControl, setModeControl } = props;

  return (
    <div className='fixed bottom-0 right-0 max-w-[220px] w-full rounded-lg  h-screen overflow-x-hidden'>
      {/* body */}
      <div className='w-full h-full pt-16 pb-2'>
        <div className='w-full bg-white rounded-lg rounded-tr-none rounded-br-none h-full shadow relative'>
          {/* title  */}

          <div className='w-full h-auto flex items-center space-x-1 px-2 pt-2 border-b pb-1'>
            <button onClick={() => setModeControl(true)} className={`px-2 py-1 rounded-lg  flex-1 ${modeControl && 'bg-gray-200'}`}>
              Control
            </button>
            <button onClick={() => setModeControl(false)} className={`px-2 py-1 rounded-lg  flex-1 ${!modeControl && 'bg-gray-200'}`}>
              Display
            </button>
          </div>
          <div className='w-full flex justify-center'>
            {modeControl ? <Control setWidgetId={props.setWidgetId} widgetMode={true} /> : <Display setWidgetId={props.setWidgetId} widgetMode />}
          </div>
        </div>
      </div>
    </div>
  );
};
