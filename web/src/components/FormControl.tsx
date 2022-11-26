import React from 'react';

interface FromControlProps {
  placeholder: string;
  type: React.HTMLInputTypeAttribute | undefined;
  inputRef: React.RefObject<HTMLInputElement>;
}

export function FormControl(props: FromControlProps) {
  const { placeholder, type, inputRef } = props;
  return (
    <React.Fragment>
      <div className='mb-4'>
        <input
          ref={inputRef}
          type={type}
          className='form-control block w-full px-2 py-1 text-md font-normal text-gray-700 bg-white  border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none'
          placeholder={placeholder}
        />
      </div>
    </React.Fragment>
  );
}
