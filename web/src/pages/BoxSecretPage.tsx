import { useQuery } from '@tanstack/react-query';
import moment from 'moment';
import React from 'react';
import { useParams } from 'react-router-dom';
import BoxContainer from '../components/BoxContainer';
import { FindBox } from '../delivery/Box';
import 'moment/locale/th';
moment.locale('th');
type Props = {};

const BoxSecret = (props: Props) => {
  const { boxId } = useParams();

  const { data: box } = useQuery(['box'], async () => await FindBox(String(boxId)));

  return (
    <React.Fragment>
      <BoxContainer>
        <div className='px-16 lg:px-48 w-full relative -z-0 '>
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

        <div className='px-16 lg:px-48 w-full '>
          <div className='w-full px-8  bg-white mt-2 py-5 rounded-lg '>
            <h2 className='text-lg text-gray-600 px-6 mb-4'>Display data table.</h2>
            <hr className='mb-4' />
            <table className='border-collapse table-auto w-full text-sm'>
              <thead>
                <tr>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>id.</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>BoxId</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Secret</th>
                </tr>
              </thead>
              <tbody className='bg-white dark:bg-slate-800'>
                {box?.boxSecret && (
                  <tr>
                    <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>1</td>
                    <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{box.id}</td>
                    <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{box.boxSecret.secret}</td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </div>
      </BoxContainer>
    </React.Fragment>
  );
};

export default BoxSecret;
