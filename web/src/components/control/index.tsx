import { useQuery } from '@tanstack/react-query';
import React from 'react';
import { getWidgetControl } from '../../delivery/WidgetControl';
import { CButton } from './CButton';
import { CButtonNumber } from './CButtonNumber';
import { CSlider } from './CSlider';
import { CSwitch } from './CSwitch';

export const controlComponent = [CButton, CSlider, CSwitch, CButtonNumber];

interface Props {
  widgetMode: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId: React.Dispatch<React.SetStateAction<number>>;
}

export function Control(props: Props) {
  const { data } = useQuery(['widgetDisplays'], getWidgetControl);

  return (
    <>
      <div className='w-full py-4 flex justify-center items-center flex-col space-y-4 px-2'>
        {data?.map((value) => (
          <React.Fragment>
            {controlComponent.map((V) => (
              <React.Fragment key={V.toString()}>
                {value.name.toLowerCase() === V.name.toLowerCase() && <V setWidgetId={props.setWidgetId} widgetMode />}
              </React.Fragment>
            ))}
          </React.Fragment>
        ))}
      </div>
    </>
  );
}
