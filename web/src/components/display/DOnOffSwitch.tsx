import { useQuery } from '@tanstack/react-query';
import React from 'react';
import { useParams } from 'react-router-dom';
import { GetDisplayData } from '../../delivery/DisplayData';
import { useSocketIO } from '../../hooks/useSocketIO';
import { DataDisplay } from '../../types/Data';
import { DisplayType } from '../../types/Display';
import NameKeyWidget from '../NameKeyWidget';

type Props = {
  canSub: string;
  widget?: DisplayType;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};

const DOnOffSwitch = (props: Props) => {
  const mock: DataDisplay[] = [
    {
      id: 1,
      createdAt: new Date(),
      updatedAt: new Date(),
      data: 900,
      label: 'None value',
      displayId: 1,
    },
  ];

  const { boxId } = useParams();

  const { canSub, widget } = props;
  const [state, setState] = React.useState<DataDisplay[] | undefined>([]);
  const [open, setOpen] = React.useState(false);

  const { isLoading } = useQuery(
    [widget?.id ? widget?.id : 'displayData'],
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
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 5) : null)}
      className={`${props.widgetMode ? 'cursor-move h-24' : 'h-full'} relative w-full shadow rounded-lg flex justify-center items-center `}
      draggable={props.widgetMode}
    >
      <NameKeyWidget open={open} setOpen={setOpen} widget={widget} />

      <div
        className={`${props.widgetMode && 'cursor-move'} mt-6 w-20 h-10 ${
          state.length > 0 ? (state[state.length - 1].data === 1 ? 'bg-yellow-200' : 'bg-gray-200') : 'bg-gray-200'
        } transition-all duration-150  rounded-full`}
      >
        <div className={`${state.length > 0 ? state[state.length - 1].data === 1 && 'translate-x-full' : ''} w-10 h-10 p-0.5 transition-all`}>
          <div className={`w-full h-full  bg-white rounded-full`} />
        </div>
      </div>
    </div>
  );
};

export default DOnOffSwitch;
