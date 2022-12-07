import React from 'react';
import { useSocketIO } from '../../hooks/useSocketIO';
import { DataDisplay } from '../../types/Data';
import { DisplayType } from '../../types/Display';

type Props = {
  canSub: string;
  widget?: DisplayType;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};

const DNumber = (props: Props) => {
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

  const { canSub, widget } = props;
  const [state, setState] = React.useState(widget?.displayData.length! > 0 ? widget?.displayData.sort((a, b) => a.id - b.id)! : mock);

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
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 2) : null)}
      className={`${props.widgetMode && 'cursor-move'} w-full h-24 shadow rounded-lg flex justify-center items-center `}
      draggable={props.widgetMode}
    >
      <div className={`${props.widgetMode && 'cursor-move'} w-20 h-10 transition-all duration-150 rounded-lg`}>
        <p className='text-2xl text-center'>{state[state.length - 1].data}</p>{' '}
      </div>
    </div>
  );
};

export default DNumber;
