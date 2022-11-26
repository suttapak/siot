import React from 'react';

type Props = {};

const Footer = (props: Props) => {
  return (
    <footer className='absolute bottom-0 left-0 w-full pt-2 m-0 bg-gray-100 flex justify-center items-center'>
      <div className='py-6 bg-white w-full text-center flex justify-center items-center'>
        <p className='text-gray-500'>
          Development by <span className='font-bold text-gray-500 '>matee suttapak © 2022</span>
        </p>
      </div>
    </footer>
  );
};

export default Footer;
