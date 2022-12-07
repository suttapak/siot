import React from 'react';
import { buildStyles, CircularProgressbar } from 'react-circular-progressbar';
import { DisplayType } from '../../types/Display';
import 'react-circular-progressbar/dist/styles.css';
import { DataControl, DataDisplay } from '../../types/Data';
import { useSocketIO } from '../../hooks/useSocketIO';

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
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 3) : null)}
      className={`${props.widgetMode && 'cursor-move w-44'} w-full h-52 shadow rounded-lg p-3 flex flex-col justify-center items-center `}
      draggable={props.widgetMode}
    >
      <CircularProgressbar
        value={state[state.length - 1].data % 100}
        text={`${(state[state.length - 1].data % 100).toFixed(2)}%`}
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
