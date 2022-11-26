import React, { useRef } from 'react';
import { CartesianGrid, Line, LineChart, Tooltip, XAxis, YAxis } from 'recharts';
import { DisplayType } from '../../types/Display';
import { Data } from '../../types/Data';

type Props = {
  widget?: DisplayType;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};

const DLineChart = (props: Props) => {
  const mockData: Data[] = [
    {
      id: 1,
      createdAt: 1,
      updatedAt: 1,
      XNumber: 100,
      YNumber: 6,
      XString: 'x',
      YString: 'y',
    },
    {
      id: 1,
      createdAt: 1,
      updatedAt: 1,
      XNumber: 200,
      YNumber: 100,
      XString: 'x',
      YString: 'y',
    },
    {
      id: 1,
      createdAt: 1,
      updatedAt: 1,
      XNumber: 300,
      YNumber: 300,
      XString: 'x',
      YString: 'y',
    },
    {
      id: 1,
      createdAt: 1,
      updatedAt: 1,
      XNumber: 400,
      YNumber: 90,
      XString: 'x',
      YString: 'y',
    },
  ];
  const [data] = React.useState(props.widget?.displayData ? props.widget?.displayData : mockData);

  const refWidth = useRef<HTMLDivElement>(null);

  return (
    <div
      ref={refWidth}
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 1) : null)}
      className={`${props.widgetMode && 'cursor-move  w-44 h-24'} border w-full h-52 shadow rounded-lg flex justify-center items-center relative `}
      draggable={props.widgetMode}
    >
      <LineChart
        width={props.widgetMode ? 176 : refWidth.current?.clientWidth}
        height={props.widgetMode ? 96 : 208}
        data={data}
        margin={{ top: 5, right: 10, left: 0, bottom: 0 }}
      >
        <CartesianGrid strokeDasharray='3 3' />
        <XAxis dataKey={'XNumber' || 'String'} />
        <YAxis />
        <Tooltip />
        <Line type='monotone' dataKey={'YNumber' || 'YString'} stroke='#8884d8' />
      </LineChart>
    </div>
  );
};

export default DLineChart;
