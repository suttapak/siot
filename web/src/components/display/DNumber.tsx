import React from 'react';
import { DisplayType } from '../../types/Display';

type Props = {
  widget?: DisplayType;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};

const DNumber = (props: Props) => {
  return (
    <div
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 2) : null)}
      className={`${props.widgetMode && 'cursor-move'} w-24 h-24 shadow rounded-lg flex justify-center items-center `}
      draggable={props.widgetMode}
    >
      <div className={`${props.widgetMode && 'cursor-move'} w-20 h-10 transition-all duration-150 rounded-lg`}>
        <p className='text-2xl text-center'>{props.widget?.displayData[props.widget.displayData.length - 1].XNumber || '0'}</p>{' '}
      </div>
    </div>
  );
};

export default DNumber;
