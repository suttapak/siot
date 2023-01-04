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

export const CButton = (props: Props) => {
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
    <React.Fragment>
      <div
        onDragStart={(e) => e.dataTransfer.setData('text/plain', '')}
        onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 1) : null)}
        className={`${props.widgetMode && 'cursor-move w-24'} relative border bg-white w-full h-24 shadow rounded-lg flex justify-center items-center `}
        draggable={props.widgetMode}
      >
        <NameKeyWidget open={open} setOpen={setOpen} widget={widget} />

        <button
          onClick={() => onPubMsg(!props.widgetMode ? { boxId: widget?.BoxId!, data: (state + 1) % 2, key: widget?.key! } : null)}
          className={`${props.widgetMode && 'cursor-move'} mt-6 w-20 h-10 ${
            state === 1 ? 'bg-yellow-200' : 'bg-gray-200'
          } transition-all duration-150 rounded-lg`}
        >
          {state === 1 ? 'On' : 'Off'}
        </button>
      </div>
    </React.Fragment>
  );
};
