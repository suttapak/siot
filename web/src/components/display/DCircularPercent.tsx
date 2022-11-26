import React from 'react';
import { buildStyles, CircularProgressbar } from 'react-circular-progressbar';
import { DisplayType } from '../../types/Display';
import 'react-circular-progressbar/dist/styles.css';
import { Data } from '../../types/Data';

type Props = {
  widget?: DisplayType;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};
const DCircularPercent = (props: Props) => {
  const mock: Data[] = [
    {
      id: 1,
      createdAt: 1,
      updatedAt: 1,
      XNumber: 1,
      YNumber: 59,
      XString: 'one',
      YString: 'five',
    },
  ];
  const [data] = React.useState(props.widget?.displayData || mock);

  if (!data) {
    return null;
  }

  return (
    <div
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 3) : null)}
      className={`${props.widgetMode && 'cursor-move w-44'} w-full h-44 shadow rounded-lg p-3 flex flex-col justify-center items-center `}
      draggable={props.widgetMode}
    >
      <CircularProgressbar
        value={data[data?.length - 1].YNumber % 100}
        text={`${data[data?.length - 1].YNumber % 100}%`}
        circleRatio={0.75}
        styles={buildStyles({
          rotation: 1 / 2 + 1 / 8,
          strokeLinecap: 'butt',
          trailColor: '#eee',
        })}
      />
      <p className='text-base'>Percent of {data[data.length - 1].XString}</p>
    </div>
  );
};

export default DCircularPercent;
