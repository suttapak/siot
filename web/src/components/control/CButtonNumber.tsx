import React from 'react';
import { AiOutlineMinus, AiOutlinePlus } from 'react-icons/ai';
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

export const CButtonNumber = (props: Props) => {
  const { canSub, widget } = props;
  const [state, setState] = React.useState(widget?.controlData.length! > 0 ? widget?.controlData[widget?.controlData.length - 1].data! : 0);
  const [open, setOpen] = React.useState(false);

  const { client } = useSocketIO();

  React.useEffect(() => {
    client.emit('subscript', { boxId: widget?.BoxId, key: canSub + '/' + widget?.key });
    // eslint-disable-next-line
  }, [canSub, widget]);

  const onPubMsg = (ctx: CtxPubType | null) => {
    client.emit('publish', ctx);
  };

  client.on(canSub + '/' + widget?.key, (data: { controlData: DataControl[] }) => {
    setState(data ? data.controlData[0].data : 0);
  });

  return (
    <div
      onDrag={() => (props.setWidgetId ? props.setWidgetId(props.widget?.id ? props.widget?.id : 2) : null)}
      className={`${
        props.widgetMode && 'cursor-move w-24'
      } relative bg-white border space-y-2 w-full h-52 shadow rounded-lg flex justify-center items-center flex-col `}
      draggable={props.widgetMode}
    >
      <NameKeyWidget open={open} setOpen={setOpen} widget={widget} />
      <div className=' flex items-center flex-col'>
        <button
          onClick={() => onPubMsg(!props.widgetMode ? { boxId: widget?.BoxId!, data: state + 1, key: widget?.key! } : null)}
          onDoubleClick={() => onPubMsg(!props.widgetMode ? { boxId: widget?.BoxId!, data: state + 10, key: widget?.key! } : null)}
          className={`${props.widgetMode && 'cursor-move'} ease-in transition-all w-10 h-10 border duration-150 rounded flex justify-center items-center`}
        >
          <AiOutlinePlus size={22} />
        </button>
        <div className='w-24 h-12 flex items-center'>
          <input
            onChange={(e) => {
              if (/^\d+(\.\d+)?$/.test(e.target.value) || e.target.value === '') {
                let data: number = 0;
                e.target.value === '' ? (data = 0) : (data = Number(e.target.value));

                setState(data);
                onPubMsg(!props.widgetMode ? { boxId: widget?.BoxId!, data: data, key: widget?.key! } : null);
              }
            }}
            type={'text'}
            className={`${props.widgetMode ? 'cursor-move' : 'cursor-pointer'} text-2xl w-full text-center `}
            value={state}
          />
        </div>
        <button
          onClick={() => onPubMsg(!props.widgetMode ? { boxId: widget?.BoxId!, data: state - 1, key: widget?.key! } : null)}
          onDoubleClick={() => onPubMsg(!props.widgetMode ? { boxId: widget?.BoxId!, data: state - 10, key: widget?.key! } : null)}
          className={`${props.widgetMode && 'cursor-move'} w-10 h-10 border transition-all duration-150 rounded flex justify-center items-center`}
        >
          <AiOutlineMinus size={22} />
        </button>
      </div>
    </div>
  );
};
