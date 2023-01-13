import { useQuery } from '@tanstack/react-query';
import React from 'react';
import { FaEye } from 'react-icons/fa';
import { HiOutlineWrenchScrewdriver } from 'react-icons/hi2';
import { useParams } from 'react-router';
import { WedgitBar } from '../components/WedgitBar';
import { FindBox } from '../delivery/Box';
import { getControls } from '../delivery/Control';

import ReactGridLayout, { Responsive, WidthProvider } from 'react-grid-layout';
import { GetDisplays } from '../delivery/DIsplay';
import { Modal } from '../components/Modal';
import { CreateControlComponent } from '../form/Control';
import { CreateDisplayComponent } from '../form/Display';
import BoxContainer from '../components/BoxContainer';
import DCircularPercent from '../components/display/DCircularPercent';
import DNumber from '../components/display/DNumber';
import DLineChart from '../components/display/DLineChart';
import { CButton } from '../components/control/CButton';
import { CButtonNumber } from '../components/control/CButtonNumber';
import { CSlider } from '../components/control/CSlider';
import { CSwitch } from '../components/control/CSwitch';

const ResponsiveGridLayout = WidthProvider(Responsive);

export function BoxDashBoard() {
  const { boxId } = useParams();

  const { data: box } = useQuery(['box'], async () => await FindBox(String(boxId)));

  const { data: controls } = useQuery(['controls'], async () => await getControls(String(boxId)));
  const { data: displays } = useQuery(['displays'], async () => await GetDisplays(String(boxId)));

  // mode - edit and power
  const [modeEdit, setModeEdit] = React.useState(false);
  const [controlMode, setControlMode] = React.useState(false);

  const [widgetId, setWidgetId] = React.useState<number>(0);
  const [layout, setLayout] = React.useState<ReactGridLayout.Layout>();

  const [open, setOpen] = React.useState(false);

  if (!controls || !displays) {
    return null;
  }

  const onDrop = (layout: ReactGridLayout.Layout[], item: ReactGridLayout.Layout, e: Event) => {
    setLayout(item);
    setOpen(true);
  };

  const layoutData = controls.map((value) => {
    return {
      i: value.layout.id.toString(),
      x: value.layout.x,
      w: value.layout.w,
      h: value.layout.h,
      y: value.layout.y,
    };
  });
  layoutData.push(
    ...displays.map((value) => {
      return {
        i: value.layout.id.toString(),
        x: value.layout.x,
        w: value.layout.w,
        h: value.layout.h,
        y: value.layout.y,
      };
    })
  );

  return (
    <React.Fragment>
      <BoxContainer>
        {modeEdit && (
          <WedgitBar
            canSub={'' + box?.canSub.canSubscribe}
            setWidgetId={setWidgetId}
            modeControl={controlMode}
            setModeControl={setControlMode}
            open={modeEdit}
            setOpen={setModeEdit}
          />
        )}
        <div className='px-16 lg:px-56 pr-3  w-full '>
          {/* title  */}
          <div className='w-full h-8 px-6 flex items-center bg-white mt-2 py-5 rounded-lg mb-4 '>
            <h3 className='w-full text-base font-medium flex-1'>{box?.name}</h3>
            <div className='flex items-center space-x-1'>
              <button
                type='button'
                onClick={() => setModeEdit(false)}
                className={`p-2 rounded-lg hover:bg-gray-200 transition-all ${!modeEdit && 'bg-gray-200'}`}
              >
                <FaEye />
              </button>
              <button
                type='button'
                onClick={() => setModeEdit(true)}
                className={`p-2 rounded-lg hover:bg-gray-200 transition-all ${modeEdit && 'bg-gray-200'}`}
              >
                <HiOutlineWrenchScrewdriver />
              </button>
            </div>
          </div>

          <ResponsiveGridLayout
            rowHeight={100}
            onDrop={onDrop}
            isDroppable={modeEdit}
            isDraggable={modeEdit}
            isResizable={false}
            cols={{ lg: 10, md: 10, sm: 10, xs: 2, xxs: 2 }}
            breakpoints={{ lg: 1024, md: 768, sm: 640, xs: 480, xxs: 0 }}
            compactType={'vertical'}
            className='layout bg-white rounded-lg   min-h-[480px]'
            layouts={{
              lg: layoutData,
            }}
          >
            {controls?.map((value) => {
              return (
                <div key={value.layout.id.toString()} className='w-full'>
                  <>
                    {value.widget.name === 'CButton' ? (
                      <CButton canSub={'' + box?.canSub.canSubscribe} widget={value} />
                    ) : value.widget.name === 'CButtonNumber' ? (
                      <CButtonNumber canSub={'' + box?.canSub.canSubscribe} widget={value} />
                    ) : value.widget.name === 'CSlider' ? (
                      <CSlider canSub={'' + box?.canSub.canSubscribe} widget={value} />
                    ) : value.widget.name === 'CSwitch' ? (
                      <CSwitch canSub={'' + box?.canSub.canSubscribe} widget={value} />
                    ) : (
                      <div className='hidden'></div>
                    )}
                  </>
                </div>
              );
            })}
            {displays?.map((value) => {
              return (
                <div key={value.layout.id.toString()}>
                  <>
                    {value.widget.name === 'DCircularPercent' ? (
                      <DCircularPercent canSub={'' + box?.canSub.canSubscribe} widget={value} />
                    ) : value.widget.name === 'DNumber' ? (
                      <DNumber canSub={'' + box?.canSub.canSubscribe} widget={value} />
                    ) : value.widget.name === 'DLineChart' ? (
                      <DLineChart canSub={'' + box?.canSub.canSubscribe} widget={value} />
                    ) : (
                      <div className='hidden'></div>
                    )}
                  </>
                </div>
              );
            })}
          </ResponsiveGridLayout>
          <Modal open={open} setOpen={setOpen}>
            {controlMode ? (
              <CreateControlComponent boxId={String(boxId)} widgetId={widgetId} layout={layout || { i: '', w: 0, h: 0, x: 0, y: 0 }} setOpen={setOpen} />
            ) : (
              <CreateDisplayComponent boxId={String(boxId)} widgetId={widgetId} layout={layout || { i: '', w: 0, h: 0, x: 0, y: 0 }} setOpen={setOpen} />
            )}
          </Modal>
        </div>
      </BoxContainer>
    </React.Fragment>
  );
}
