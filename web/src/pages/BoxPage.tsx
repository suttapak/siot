import { useQuery } from '@tanstack/react-query';
import moment from 'moment';
import React from 'react';
import { useParams } from 'react-router-dom';
import BoxContainer from '../components/BoxContainer';
import { FindBox } from '../delivery/Box';
import { getControls } from '../delivery/Control';
import { GetDisplays } from '../delivery/DIsplay';
import 'moment/locale/th';
moment.locale('th');
type Props = {};

const BoxPage = (props: Props) => {
  const { boxId } = useParams();

  const { data: box } = useQuery(['box'], async () => await FindBox(String(boxId)));

  const { data: controls } = useQuery(['controls'], async () => await getControls(String(boxId)));
  const { data: displays } = useQuery(['displays'], async () => await GetDisplays(String(boxId)));

  return (
    <React.Fragment>
      <BoxContainer>
        <div className='px-16 lg:px-56 pr-3  w-full relative -z-0 '>
          <div className='relative w-full mt-2'>
            <div className='absolute -z-0 top-0 w-full left-0 rounded-lg h-full bg-gray-200 overflow-hidden opacity-90'>
              <img src='/iot.jpeg' alt='iot' className='w-full h-auto  block' />
            </div>
            <div className='absolute z-10 top-0 w-full left-0 rounded-lg h-full bg-gray-500 overflow-hidden opacity-20'></div>
            <div className='pt-72 w-full relative z-20 '>
              <div className='w-full  px-6  bg-white bg-opacity-95  py-10 rounded-b-lg mb-4 '>
                <div className='w-full relative'>
                  <h1 className='text-xl font-medium text-gray-600 mb-2 w-full py-1 px-6 '>{box?.name}.</h1>
                  {box?.description && (
                    <React.Fragment>
                      <div className='py-4 px-6 w-full bg-opacity-90 rounded-lg bg-slate-200'>
                        <p className='font-normal text-base pb-1'>
                          สร้างเมื่อ {moment(box.createdAt).format('LL')} มีสมาชิกทั่งหมด {box.members.length} accounts.
                        </p>
                        <p className='text-base '>{box?.description}</p>
                      </div>
                    </React.Fragment>
                  )}
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className='px-16 lg:px-56 pr-3  w-full '>
          <div className='w-full px-8  bg-white mt-2 py-5 rounded-lg mb-4 '>
            <h2 className='text-lg text-gray-600 px-6 mb-4'>Control data table.</h2>
            <hr className='mb-4' />
            <table className='border-collapse table-auto w-full text-sm'>
              <thead>
                <tr>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>id.</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Widget</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Name</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Key</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Description</th>
                </tr>
              </thead>
              <tbody className='bg-white dark:bg-slate-800'>
                {controls?.map((value, index) => {
                  return (
                    <tr key={index}>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{index + 1}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{value.widget.name}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{value.name}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{value.key}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{value.description}</td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>
        </div>
        <div className='px-16 lg:px-56 pr-3  w-full '>
          <div className='w-full px-8  bg-white mt-2 py-5 rounded-lg '>
            <h2 className='text-lg text-gray-600 px-6 mb-4'>Display data table.</h2>
            <hr className='mb-4' />
            <table className='border-collapse table-auto w-full text-sm'>
              <thead>
                <tr>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>id.</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Widget</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Name</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Key</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Description</th>
                </tr>
              </thead>
              <tbody className='bg-white dark:bg-slate-800'>
                {displays?.map((value, index) => {
                  return (
                    <tr key={index}>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{index + 1}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{value.widget.name}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{value.name}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{value.key}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{value.description}</td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>
        </div>
      </BoxContainer>
    </React.Fragment>
  );
};

export default BoxPage;
