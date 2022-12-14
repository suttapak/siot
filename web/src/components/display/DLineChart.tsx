import React, { useRef } from 'react';
import { CartesianGrid, Line, LineChart, Tooltip, XAxis, YAxis } from 'recharts';
import { DisplayType } from '../../types/Display';
import { DataDisplay } from '../../types/Data';
import { useSocketIO } from '../../hooks/useSocketIO';
import { useQuery } from '@tanstack/react-query';
import { GetDisplayData } from '../../delivery/DisplayData';
import { useParams } from 'react-router-dom';

type Props = {
  canSub: string;
  widget?: DisplayType;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};

const mockData: DataDisplay[] = [
  {
    id: 1,
    createdAt: new Date(),
    updatedAt: new Date(),
    data: 99,
    label: '22 aug 2000',
    displayId: 1,
  },
  {
    id: 1,
    createdAt: new Date(),
    updatedAt: new Date(),
    data: 700,
    label: '22 aug 2000',
    displayId: 1,
  },
  {
    id: 1,
    createdAt: new Date(),
    updatedAt: new Date(),
    data: 320,
    label: '22 aug 2000',
    displayId: 1,
  },
  {
    id: 1,
    createdAt: new Date(),
    updatedAt: new Date(),
    data: 1023,
    label: '22 aug 2000',
    displayId: 1,
  },
];

const DLineChart = (props: Props) => {
  const { canSub, widget } = props;
  const [state, setState] = React.useState<DataDisplay[] | undefined>([]);
  const { boxId } = useParams();

  const { isLoading } = useQuery(
    [widget?.key ? widget?.key : 'displayData'],
    async () => {
      if (!boxId || !widget?.id) {
        return;
      }
      return await GetDisplayData({ boxId: boxId, displayId: widget?.id });
    },
    {
      onSuccess(data) {
        setState(data);
      },
    }
  );
  const { client } = useSocketIO();

  React.useEffect(() => {
    setState(widget?.displayData.length! > 0 ? widget?.displayData! : mockData);
    // eslint-disable-next-line
  }, [widget]);

  React.useEffect(() => {
    if (widget) {
      client.emit('subscript', { boxId: widget?.BoxId, key: canSub + '/' + widget?.key });
    }
    // eslint-disable-next-line
  }, [canSub, widget]);

  React.useEffect(() => {
    if (widget) {
      client.emit('subscript', { boxId: widget?.BoxId, key: canSub + '/' + widget?.key });
    }
    // eslint-disable-next-line
  }, [client.connected]);

  client.on(canSub + '/' + widget?.key, (data: { displayData: DataDisplay[] }) => {
    setState(data.displayData);
  });
  const refWidth = useRef<HTMLDivElement>(null);

  const [width, setWidth] = React.useState(300);

  React.useEffect(() => {
    setWidth(refWidth.current?.clientWidth ? refWidth.current?.clientWidth : 300);
  }, [refWidth]);

  if (isLoading || !state) {
    return (
      <div role='status' className={`${props.widgetMode && 'cursor-move'} w-full h-24 shadow rounded-lg flex justify-center items-center `}>
        <div className='w-full dark:bg-gray-700 rounded-lg mb-4'></div>
      </div>
    );
  }

  return (
    <div
      ref={refWidth}
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 1) : null)}
      className={`${props.widgetMode && 'cursor-move  w-44 h-24'} border w-full h-52 shadow rounded-lg flex justify-center items-center relative `}
      draggable={props.widgetMode}
    >
      <LineChart
        width={props.widgetMode ? 176 : width}
        height={props.widgetMode ? 96 : 208}
        data={state.sort((a, b) => a.id - b.id)}
        margin={{ top: 5, right: 10, left: 0, bottom: 0 }}
      >
        <CartesianGrid strokeDasharray='3 3' />
        <XAxis dataKey={'label'} />
        <YAxis />
        <Tooltip />
        <Line type='monotone' dataKey={'data'} stroke='#8884d8' />
      </LineChart>
    </div>
  );
};

export default DLineChart;
