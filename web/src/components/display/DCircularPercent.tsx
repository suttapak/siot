import React from 'react';
import { buildStyles, CircularProgressbar } from 'react-circular-progressbar';
import { DisplayType } from '../../types/Display';
import 'react-circular-progressbar/dist/styles.css';
import { DataDisplay } from '../../types/Data';
import { useSocketIO } from '../../hooks/useSocketIO';
import { useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { GetDisplayData } from '../../delivery/DisplayData';
import NameKeyWidget from '../NameKeyWidget';

type Props = {
  canSub: string;
  widget?: DisplayType;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};
const DCircularPercent = (props: Props) => {
  const mock: DataDisplay[] = [
    {
      id: 1,
      createdAt: new Date(),
      updatedAt: new Date(),
      data: 0,
      label: 'None value',
      displayId: 1,
    },
  ];

  const { canSub, widget } = props;
  const { boxId } = useParams();
  const [state, setState] = React.useState<DataDisplay[] | undefined>([]);
  const [open, setOpen] = React.useState(false);

  const { isLoading } = useQuery(
    [widget?.key ? widget?.key : 'displayData'],
    async () => {
      if (!widget?.id) {
        return;
      }
      return await GetDisplayData({ boxId: String(boxId), displayId: widget?.id });
    },
    {
      onSuccess(data) {
        setState(data);
      },
    }
  );
  const { client } = useSocketIO();

  React.useEffect(() => {
    setState(widget?.displayData.length! > 0 ? widget?.displayData! : mock);
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

  if (isLoading || !state) {
    return (
      <div role='status' className={`${props.widgetMode && 'cursor-move'} w-full h-24 shadow rounded-lg flex justify-center items-center `}>
        <div className='w-full dark:bg-gray-700 rounded-lg mb-4'></div>
      </div>
    );
  }
  return (
    <div
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 3) : null)}
      className={`${props.widgetMode && 'cursor-move w-44'} relative w-full h-52 shadow rounded-lg p-3 flex flex-col justify-center items-center `}
      draggable={props.widgetMode}
    >
      <NameKeyWidget open={open} setOpen={setOpen} widget={widget} />

      <CircularProgressbar
        value={state.length > 0 ? state[state.length - 1].data % 100 : 0}
        text={state.length > 0 ? `${(state[state.length - 1].data % 100).toFixed(2)}%` : 'NULL'}
        circleRatio={0.75}
        styles={buildStyles({
          rotation: 1 / 2 + 1 / 8,
          strokeLinecap: 'butt',
          trailColor: '#eee',
        })}
      />
      <p className='text-base'>Percent of {widget?.name}</p>
    </div>
  );
};

export default DCircularPercent;
