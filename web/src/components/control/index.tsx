import { useQuery } from '@tanstack/react-query';
import React from 'react';
import { getWidgetControl } from '../../delivery/WidgetControl';
import { CButton } from './CButton';
import { CButtonNumber } from './CButtonNumber';
import { CSlider } from './CSlider';
import { CSwitch } from './CSwitch';

interface Props {
  canSub: string;
  widgetMode: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId: React.Dispatch<React.SetStateAction<number>>;
}

export function Control(props: Props) {
  const { data } = useQuery(['widgetDisplays'], getWidgetControl);
  const { canSub, setWidgetId } = props;

  return (
    <>
      <div className='w-full py-4 flex justify-center items-center flex-col space-y-4 px-2'>
        {data?.map((value) => (
          <React.Fragment>
            {value.name === 'CButton' ? (
              <CButton canSub={canSub} setWidgetId={setWidgetId} widgetMode />
            ) : value.name === 'CButtonNumber' ? (
              <CButtonNumber canSub={canSub} setWidgetId={setWidgetId} widgetMode />
            ) : value.name === 'CSlider' ? (
              <CSlider canSub={canSub} setWidgetId={setWidgetId} widgetMode />
            ) : value.name === 'CSwitch' ? (
              <CSwitch canSub={canSub} setWidgetId={setWidgetId} widgetMode />
            ) : (
              <div className='hidden'></div>
            )}
          </React.Fragment>
        ))}
      </div>
    </>
  );
}
