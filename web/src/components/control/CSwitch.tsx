import React from 'react';
import { Control } from '../../types/Control';

type Props = {
  widgetMode?: boolean;
  widget?: Control;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};
export const CSwitch = (props: Props) => {
  const [click, setClick] = React.useState<number>(0);

  return (
    <div
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 4) : null)}
      className={`${props.widgetMode && 'cursor-move w-24'} border bg-white w-full h-24 shadow rounded-lg flex justify-center items-center `}
      draggable={props.widgetMode}
    >
      <button
        onClick={() => setClick(!props.widgetMode ? (click + 1) % 2 : click)}
        className={`${props.widgetMode && 'cursor-move'} w-20 h-10 bg-gray-200 rounded-full`}
      >
        <div className={`${click && 'translate-x-full'} w-10 h-10 p-0.5 transition-all`}>
          <div className={`w-full h-full  bg-white rounded-full`} />
        </div>
      </button>
    </div>
  );
};
