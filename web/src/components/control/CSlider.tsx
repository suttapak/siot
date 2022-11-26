import React from 'react';
import { Control } from '../../types/Control';

type Props = {
  widget?: Control;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};
export const CSlider = (props: Props) => {
  const [state, setState] = React.useState<number>(0);

  return (
    <div
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 3) : null)}
      className={`${props.widgetMode && 'cursor-move w-44'} border bg-white w-full h-24 shadow rounded-lg flex justify-center items-center relative `}
      draggable={props.widgetMode}
    >
      <input type='range' name='slider' id='slider' value={state} onChange={(e) => setState(e.target.valueAsNumber)} />
      <p className='absolute bottom-0 text-gray-500 right-0 p-2'>{state}</p>
    </div>
  );
};
