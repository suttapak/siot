import React from 'react';
import { Control } from './control';
import Display from './display';

type Props = {
  canSub: string;
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  setWidgetId: React.Dispatch<React.SetStateAction<number>>;
  modeControl: boolean;
  setModeControl: React.Dispatch<React.SetStateAction<boolean>>;
};

export const WedgitBar = (props: Props) => {
  const { modeControl, setModeControl } = props;

  return (
    <div className='max-w-[220px] w-full h-screen overflow-x-hidden'>
      {/* body */}
      <div className='w-full h-full pb-2'>
        <div className='w-full bg-whiteh-full relative'>
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
            {modeControl ? (
              <Control canSub={props.canSub} setWidgetId={props.setWidgetId} widgetMode={true} />
            ) : (
              <Display canSub={props.canSub} setWidgetId={props.setWidgetId} widgetMode />
            )}
          </div>
        </div>
      </div>
    </div>
  );
};
