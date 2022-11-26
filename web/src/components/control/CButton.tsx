import React from 'react';
import { Control } from '../../types/Control';

type Props = {
  widget?: Control;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};

export const CButton = (props: Props) => {
  const [state, setState] = React.useState<number>(0);

  return (
    <div
      onDragStart={(e) => e.dataTransfer.setData('text/plain', '')}
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 1) : null)}
      className={`${props.widgetMode && 'cursor-move w-24'} border bg-white w-full h-24 shadow rounded-lg flex justify-center items-center `}
      draggable={props.widgetMode}
    >
      <button
        onClick={() => setState(!props.widgetMode ? (state + 1) % 2 : state)}
        className={`${props.widgetMode && 'cursor-move'} w-20 h-10 ${state === 1 ? 'bg-yellow-200' : 'bg-gray-200'} transition-all duration-150 rounded-lg`}
      >
        {state === 1 ? 'On' : 'Off'}
      </button>
    </div>
  );
};
