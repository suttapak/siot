import * as React from 'react';

import { Link } from 'react-router-dom';
import { useAuth } from '../hooks';

import { MdSpaceDashboard } from 'react-icons/md';
import { CgMenuGridO } from 'react-icons/cg';
import { AiOutlineLogin, AiOutlineLogout } from 'react-icons/ai';
import { GrClose } from 'react-icons/gr';

interface ItemProps {
  path: string;
  children: React.ReactNode;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

function Items({ path, children, setOpen }: ItemProps): JSX.Element {
  return (
    <>
      <li>
        <Link
          onClick={() => setOpen(false)}
          to={path}
          className='flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700'
        >
          {children}
        </Link>
      </li>
    </>
  );
}

interface NavbarComponentProps {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

export function NavbarComponent({ open, setOpen }: NavbarComponentProps) {
  const auth = useAuth();

  return (
    <React.Fragment>
      {/* <!-- drawer init and show --> */}
      <div className='fixed top-0 left-0 z-40 w-full h-10 bg-white  flex justify-center items-center'>
        <div className='w-full px-3  mx-auto flex justify-center items-center'>
          <button onClick={() => setOpen(!open)} className='text-gray-800 font-bold text-lg ' type='button'>
            <CgMenuGridO size={30} color='gray' />
          </button>
          <div className='w-full mx-auto flex justify-start'>
            <Link to={'/'} className=''>
              <h1 className='text-gray-800 pl-4 font-bold font-sans text-lg md:text-2xl'>SIOT</h1>
            </Link>
          </div>
          {auth.user && (
            <div className='w-9 h-9 rounded-full pointer-events-none overflow-hidden flex justify-center items-center bg-blue-300'>
              <img src={process.env.REACT_APP_SERVER_URL + auth.user.avatar.url.substring(1)} alt={auth.user.avatar.title} />
            </div>
          )}
        </div>
      </div>

      {/* <!-- drawer component --> */}
      <div
        className={`fixed border-r z-40 h-screen p-4 overflow-y-auto bg-white w-64 dark:bg-gray-800 transition-transform ${
          open ? 'translate-x-0' : '-translate-x-full'
        } `}
        aria-labelledby='drawer-navigation-label'
      >
        <h5 className='text-base font-semibold text-gray-500 uppercase dark:text-gray-400'>Menu</h5>
        <button
          onClick={() => setOpen(false)}
          type='button'
          data-drawer-dismiss='drawer-navigation'
          aria-controls='drawer-navigation'
          className='text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 absolute top-2.5 right-2.5 inline-flex items-center dark:hover:bg-gray-600 dark:hover:text-white'
        >
          <GrClose />
          <span className='sr-only'>Close menu</span>
        </button>
        <div className='py-4 overflow-y-auto'>
          <ul className='space-y-2 flex-col'>
            <Items path='/' setOpen={setOpen}>
              <MdSpaceDashboard />
              <span className='ml-3'>Dashboard</span>
            </Items>
            {!auth.user ? (
              <>
                <Items path='/signin' setOpen={setOpen}>
                  <AiOutlineLogin />
                  <span className='ml-3'>Signin</span>
                </Items>
              </>
            ) : (
              <button
                onClick={() => {
                  auth.signout();
                  setOpen(false);
                }}
                className='flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700'
              >
                <AiOutlineLogout />
                <span className='ml-3'>Signout</span>
              </button>
            )}
          </ul>
        </div>
      </div>
      <div onClick={() => setOpen(false)} className={`${open ? 'block' : 'hidden'} bg-gray-900 bg-opacity-50 dark:bg-opacity-80 fixed inset-0 z-30`}></div>
    </React.Fragment>
  );
}
