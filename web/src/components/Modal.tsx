import React from 'react';

interface ModalProps {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  children: React.ReactNode;
}

export function Modal({ children, open, setOpen }: ModalProps) {
  return (
    <>
      <div
        aria-hidden='true'
        className={`${
          open ? 'fixed' : 'hidden'
        } overflow-y-auto overflow-x-hidden max-w-md mx-auto fixed top-0 right-0 left-0 z-50 w-auto md:inset-0 h-modal md:h-full`}
      >
        <div className='relative top-1/2 -translate-y-1/2 p-4 w-auto mx-auto max-w-md h-full md:h-auto'>
          <div className='relative bg-white rounded-lg shadow dark:bg-gray-700'>
            <button
              onClick={() => setOpen(false)}
              type='button'
              className='absolute top-3 right-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-800 dark:hover:text-white'
              data-modal-toggle='authentication-modal'
            >
              <svg aria-hidden='true' className='w-5 h-5' fill='currentColor' viewBox='0 0 20 20' xmlns='http://www.w3.org/2000/svg'>
                <path
                  fillRule='evenodd'
                  d='M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z'
                  clipRule='evenodd'
                ></path>
              </svg>
              <span className='sr-only'>Close modal</span>
            </button>
            <div className='py-6 px-6 lg:px-8'>{children}</div>
          </div>
        </div>
      </div>
      <div onClick={() => setOpen(false)} className={`${open ? 'block' : 'hidden'} bg-gray-900 bg-opacity-20 dark:bg-opacity-80 fixed inset-0 z-30`}></div>
    </>
  );
}
