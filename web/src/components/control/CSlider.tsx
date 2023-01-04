import React from 'react';
import { useSocketIO } from '../../hooks/useSocketIO';
import { Control } from '../../types/Control';
import { CtxPubType } from '../../types/CtxPub.type';
import { DataControl } from '../../types/Data';
import NameKeyWidget from '../NameKeyWidget';

type Props = {
  canSub: string;
  widget?: Control;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId?: React.Dispatch<React.SetStateAction<number>>;
};
export const CSlider = (props: Props) => {
  const { canSub, widget } = props;
  const [state, setState] = React.useState(widget?.controlData.length! > 0 ? widget?.controlData[widget?.controlData.length - 1].data! : 0);
  const [open, setOpen] = React.useState(false);

  const { client } = useSocketIO();

  React.useEffect(() => {
    client.emit('subscript', { boxId: widget?.BoxId, key: canSub + '/' + widget?.key });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [canSub, widget]);

  const onPubMsg = (ctx: CtxPubType | null) => {
    client.emit('publish', ctx);
  };

  client.on(canSub + '/' + widget?.key, (data: { controlData: DataControl[] }) => {
    setState(data.controlData[0].data);
  });

  return (
    <div
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 3) : null)}
      className={`${props.widgetMode && 'cursor-move w-44'} border bg-white w-full h-24 shadow rounded-lg flex justify-center items-center relative `}
      draggable={props.widgetMode}
    >
      <NameKeyWidget open={open} setOpen={setOpen} widget={widget} />
      <input
        type='range'
        name='slider'
        id='slider'
        value={state}
        onChange={(e) => onPubMsg(!props.widgetMode ? { boxId: widget?.BoxId!, data: Number(e.target.valueAsNumber), key: widget?.key! } : null)}
      />
      <p className='absolute bottom-0 text-gray-500 right-0 p-2'>{state}</p>
    </div>
  );
};
