import React from 'react';
import { AiOutlineMinus, AiOutlinePlus } from 'react-icons/ai';
import { Control } from '../../types/Control';

type Props = {
  widget?: Control;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};

export const CButtonNumber = (props: Props) => {
  const [state, setState] = React.useState<number>(0);

  const plusButton = React.useRef<HTMLButtonElement>() as React.MutableRefObject<HTMLButtonElement>;
  const minusButton = React.useRef<HTMLButtonElement>() as React.MutableRefObject<HTMLButtonElement>;

  return (
    <div
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 2) : null)}
      className={`${props.widgetMode && 'cursor-move w-24'} bg-white border space-y-2 w-full h-44 shadow rounded-lg flex justify-center items-center flex-col `}
      draggable={props.widgetMode}
    >
      <button
        ref={plusButton}
        onDoubleClick={() => setState(!props.widgetMode ? state + 10 : state)}
        onClick={() => {
          setState(!props.widgetMode ? state + 1 : state);
        }}
        className={`${props.widgetMode && 'cursor-move'} ease-in transition-all w-10 h-10 border duration-150 rounded flex justify-center items-center`}
      >
        <AiOutlinePlus size={22} />
      </button>
      <div className='w-24 h-12 flex items-center'>
        <p onDoubleClick={() => setState(0)} className={`${props.widgetMode ? 'cursor-move' : 'cursor-pointer'} text-2xl w-full text-center `}>
          {state}
        </p>
      </div>
      <button
        ref={minusButton}
        onClick={() => setState(!props.widgetMode ? state - 1 : state)}
        onDoubleClick={() => setState(!props.widgetMode ? state - 10 : state)}
        className={`${props.widgetMode && 'cursor-move'} w-10 h-10 border transition-all duration-150 rounded flex justify-center items-center`}
      >
        <AiOutlineMinus size={22} />
      </button>
    </div>
  );
};
