import { useQuery } from '@tanstack/react-query';
import React from 'react';
import { getWidgetDisplays } from '../../delivery/WidgetDisplay';
import DCircularPercent from './DCircularPercent';
import DLineChart from './DLineChart';
import DNumber from './DNumber';
import DOnOff from './DOnOff';
import DOnOffSwitch from './DOnOffSwitch';

interface Props {
  canSub: string;
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId: React.Dispatch<React.SetStateAction<number>>;
}

export const displayComponent = [DNumber, DLineChart, DCircularPercent];

const Display = (props: Props) => {
  const { data } = useQuery(['widgetDisplays'], getWidgetDisplays);
  const { canSub, setWidgetId } = props;

  return (
    <>
      <div className='w-full py-4 flex justify-center items-center flex-col space-y-4 px-2'>
        {data?.map((value) => (
          <React.Fragment key={value.id}>
            {value.name === 'DLineChart' ? (
              <DLineChart canSub={canSub} setWidgetId={setWidgetId} widgetMode />
            ) : value.name === 'DNumber' ? (
              <DNumber canSub={canSub} setWidgetId={setWidgetId} widgetMode />
            ) : value.name === 'DCircularPercent' ? (
              <DCircularPercent canSub={canSub} setWidgetId={setWidgetId} widgetMode />
            ) : value.name === 'DOnOff' ? (
              <DOnOff canSub={canSub} setWidgetId={setWidgetId} widgetMode />
            ) : value.name === 'DOnOffSwitch' ? (
              <DOnOffSwitch canSub={canSub} setWidgetId={setWidgetId} widgetMode />
            ) : (
              <div className='hidden'></div>
            )}
          </React.Fragment>
        ))}
      </div>
    </>
  );
};

export default Display;
