import { useQuery } from '@tanstack/react-query';
import moment from 'moment';
import React from 'react';
import { useParams } from 'react-router-dom';
import BoxContainer from '../components/BoxContainer';
import { FindBox } from '../delivery/Box';
import 'moment/locale/th';
import { GetBoxMembers } from '../delivery/BoxMember';
import { AiOutlineCheck, AiOutlineClose, AiOutlineUsergroupAdd } from 'react-icons/ai';
import { Modal } from '../components/Modal';
import { CreateMemberComponent } from '../form/Member';
moment.locale('th');
type Props = {};

const BoxMemberPage = (props: Props) => {
  const { boxId } = useParams();

  const { data: box } = useQuery(['box'], async () => await FindBox(String(boxId)));
  const { data: member } = useQuery(['members'], async () => await GetBoxMembers(String(boxId)));

  const [open, setOpen] = React.useState(false);

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
            <div className='w-full flex'>
              <div className='h-10 flex-1 flex items-center justify-center'>
                <h2 className='align-middle text-lg text-gray-600 px-6  flex-1'>Members.</h2>
              </div>
              <button
                onClick={() => setOpen(true)}
                className='mb-4 w-10 h-10 p-1 rounded-full flex justify-center items-center hover:bg-gray-200 bg-gray-100 overflow-hidden'
              >
                <AiOutlineUsergroupAdd size={26} />
              </button>
            </div>
            <hr className='mb-4' />
            <table className='border-collapse table-auto w-full text-sm'>
              <thead>
                <tr>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>id.</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Email</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Frist Name</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>lastname</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Can Read</th>
                  <th className='border-b dark:border-slate-600 font-medium p-4 pl-8 pt-0 pb-3 text-gray-700 dark:text-gray-200 text-left'>Can Write</th>
                </tr>
              </thead>
              <tbody className='bg-white dark:bg-slate-800'>
                {/* TODO : add box member */}
                {member?.map((value, index) => (
                  <React.Fragment>
                    <tr>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{index + 1}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{value.user.email}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{value.user.firstName}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>{value.user.lastName}</td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>
                        {value.boxMemberPermission.canRead ? <AiOutlineCheck /> : <AiOutlineClose />}
                      </td>
                      <td className='border-b border-slate-100 dark:border-slate-700 p-4 pl-8 text-gray-500 dark:text-gray-400'>
                        {value.boxMemberPermission.canWrite ? <AiOutlineCheck /> : <AiOutlineClose />}
                      </td>
                    </tr>
                  </React.Fragment>
                ))}
              </tbody>
            </table>
          </div>
        </div>
        <Modal open={open} setOpen={setOpen}>
          <CreateMemberComponent setOpen={setOpen} boxId={String(boxId)} />
        </Modal>
      </BoxContainer>
    </React.Fragment>
  );
};

export default BoxMemberPage;
