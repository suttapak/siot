import React from 'react';
import { useToast } from '../hooks/useToast';
import Footer from './Footer';
import { NavbarComponent } from './Navbar';
import { Toast } from './Toast';

interface Props {
  children: JSX.Element;
}

export function Container({ children }: Props): JSX.Element {
  const [open, setOpen] = React.useState<boolean>(false);
  const toast = useToast();
  return (
    <>
      <NavbarComponent open={open} setOpen={setOpen} />
      <div
        className={`bg-gray-100 pb-12 min-h-screen h-full ${
          open ? 'pl-64' : 'pl-0 '
        } flex-col justify-center items-center transition-all text-xs md:text-sm pt-14`}
      >
        <div className='h-max'>{children}</div>
      </div>
      <Footer />
      <Toast message={toast.message} />
    </>
  );
}
