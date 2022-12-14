import React from 'react';
import { Link, useLocation, useParams } from 'react-router-dom';
import { BsPersonLinesFill } from 'react-icons/bs';
import { MdDashboardCustomize } from 'react-icons/md';
import { FaUserSecret } from 'react-icons/fa';
import { AiOutlineCodeSandbox } from 'react-icons/ai';

type Props = {};

interface ItemProps {
  path: string;
  children: React.ReactNode;
  active: string;
  setActive: React.Dispatch<React.SetStateAction<string>>;
}

function LinkItem({ path, children, active, setActive }: ItemProps): JSX.Element {
  return (
    <Link
      onClick={() => {
        setActive(path);
      }}
      to={path}
      className={`flex items-center px-2 py-1 h-8 text-base font-normal ${
        active === path ? 'bg-gray-200' : ''
      }  text-gray-900 rounded-lg dark:text-white hover:bg-gray-200 dark:hover:bg-gray-700`}
    >
      {children}
    </Link>
  );
}
export const MenuBar = (props: Props) => {
  const { boxId } = useParams();
  const path = useLocation();

  React.useEffect(() => {
    setActive(path.pathname);
    // eslint-disable-next-line
  }, []);

  const [active, setActive] = React.useState('');

  return (
    <div className='max-w-[51px] lg:max-w-[220px] z-10  w-full p-2 space-y-2 fixed pt-16 bottom-0 left-0 h-screen'>
      <LinkItem setActive={setActive} path={`/boxes/${boxId}`} active={active}>
        <div className='mr-4'>
          <AiOutlineCodeSandbox size={18} />
        </div>
        <span className='lg:block hidden'>Box</span>
      </LinkItem>
      <LinkItem setActive={setActive} path={`/boxes/${boxId}/dashboard`} active={active}>
        <div className='mr-4'>
          <MdDashboardCustomize size={18} />
        </div>
        <span className='lg:block hidden'>Dashboard</span>
      </LinkItem>
      <LinkItem setActive={setActive} path={`/boxes/${boxId}/secret`} active={active}>
        <div className='mr-4'>
          <FaUserSecret size={18} />
        </div>
        <span className='lg:block hidden'>Box Secret</span>
      </LinkItem>
      <LinkItem setActive={setActive} path={`/boxes/${boxId}/members`} active={active}>
        <div className='mr-4'>
          <BsPersonLinesFill size={18} />
        </div>
        <span className='lg:block hidden'>Member of Box</span>
      </LinkItem>
    </div>
  );
};
