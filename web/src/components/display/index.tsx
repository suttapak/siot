import { useQuery } from '@tanstack/react-query';
import React from 'react';
import { getWidgetDisplays } from '../../delivery/WidgetDisplay';
import DCircularPercent from './DCircularPercent';
import DLineChart from './DLineChart';
import DNumber from './DNumber';

interface Props {
  widgetMode?: boolean;
  // optional setWidget ID when drag to new widget
  setWidgetId: React.Dispatch<React.SetStateAction<number>>;
}

export const displayComponent = [DNumber, DLineChart, DCircularPercent];

const Display = (props: Props) => {
  const { data } = useQuery(['widgetDisplays'], getWidgetDisplays);

  return (
    <>
      <div className='w-full py-4 flex justify-center items-center flex-col space-y-4 px-2'>
        {data?.map((value) => (
          <>
            {displayComponent.map((V, i) => (
              <React.Fragment key={i}>
                {V.name.toLowerCase() === value.name.toLocaleLowerCase() && <V setWidgetId={props.setWidgetId} widgetMode />}
              </React.Fragment>
            ))}
          </>
        ))}
      </div>
    </>
  );
};

export default Display;
