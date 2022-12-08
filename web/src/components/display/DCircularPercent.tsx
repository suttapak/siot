import React from 'react';
import { buildStyles, CircularProgressbar } from 'react-circular-progressbar';
import { DisplayType } from '../../types/Display';
import 'react-circular-progressbar/dist/styles.css';
import { DataControl, DataDisplay } from '../../types/Data';
import { useSocketIO } from '../../hooks/useSocketIO';
import { useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { GetDisplayData } from '../../delivery/DisplayData';

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
  const [state, setState] = React.useState<DataDisplay[]>([]);

  const { isLoading } = useQuery([widget?.key], async () => await GetDisplayData({ boxId: String(boxId), displayId: widget?.id! }), {
    onSuccess(data) {
      setState(data);
    },
  });
  const { client } = useSocketIO();

  React.useEffect(() => {
    setState(widget?.displayData.length! > 0 ? widget?.displayData! : mock);
  }, [widget]);

  React.useEffect(() => {
    if (widget) {
      client.emit('subscript', { boxId: widget?.BoxId, key: canSub + '/' + widget?.key });
    }
  }, [canSub, widget]);

  React.useEffect(() => {
    if (widget) {
      client.emit('subscript', { boxId: widget?.BoxId, key: canSub + '/' + widget?.key });
    }
  }, [client.connected]);

  client.on(canSub + '/' + widget?.key, (data: { displayData: DataDisplay[] }) => {
    setState(data.displayData);
  });
  return (
    <div
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 3) : null)}
      className={`${props.widgetMode && 'cursor-move w-44'} w-full h-52 shadow rounded-lg p-3 flex flex-col justify-center items-center `}
      draggable={props.widgetMode}
    >
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
