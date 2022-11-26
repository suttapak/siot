import React, { createRef, useState } from 'react';
import { AiOutlineLogin } from 'react-icons/ai';
import { ImSpinner2 } from 'react-icons/im';
import { Link, useNavigate } from 'react-router-dom';
import { Alert, FormControl } from '../components';
import { useAuth } from '../hooks';
import { validateEmail } from '../libs';

export function SignUp() {
  const emailRef = createRef<HTMLInputElement>();
  const passwordRef = createRef<HTMLInputElement>();
  const firstName = createRef<HTMLInputElement>();
  const lastName = createRef<HTMLInputElement>();

  const auth = useAuth();
  const navigate = useNavigate();
  const [error, setError] = useState<string>();
  const [loading, setLoading] = useState<boolean>(false);

  function handlerSignUP(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setLoading(true);
    if (!emailRef.current?.value || !passwordRef.current?.value) {
      setLoading(false);
      setError(!emailRef.current?.value ? 'Email must not empty!' : 'Password must not empty!');
      return;
    }

    if (!validateEmail(emailRef.current.value)) {
      setError('Email is invalide!');
      setLoading(false);
      return;
    }
    auth
      .signup({
        email: emailRef.current.value,
        password: passwordRef.current.value,
        firstName: firstName.current?.value,
        lastName: lastName.current?.value,
      })
      .then(() => {
        setLoading(false);
        navigate('/');
      })
      .catch((error) => {
        setLoading(false);
        setError(error);
      });
  }

  return (
    <>
      <div className='max-w-sm min-h-96 h-auto bg-white rounded  mx-auto py-4 px-3 shadow-sm drop-shadow'>
        {/* <-------- title --------> */}
        <div className='mt-4 mb-2 shadow-gray-500 border-b'>
          <h1 className='font-bold text-xl md:text-4xl text-center mb-8 p-4'>SIOT</h1>
          <h2 className='font-bold text-md md:text-lg text-center text-gray-700 mb-4'>Sign up for Free accout</h2>
        </div>
        <div className='w-full py-2 px-2'>
          <a href='https://github.com/suttapak/siot' className='text-blue-500 hover:border-b border-blue-500'>
            github repository
          </a>
        </div>
        <div className='mt-4 mb-4 px-2'>
          <form onSubmit={handlerSignUP} className='flex flex-col justify-center'>
            <FormControl type='text' placeholder='First Name.' inputRef={firstName} />
            <FormControl type='text' placeholder='Last Name.' inputRef={lastName} />
            <FormControl type='text' placeholder='Email.' inputRef={emailRef} />
            <FormControl type='password' placeholder='Password.' inputRef={passwordRef} />
            <button type='submit' className='flext justify-center items-center' disabled={loading}>
              <div className={`flex justify-center items-center text-lg md:text-2xl font-bold text-gray-800`}>
                {loading ? <ImSpinner2 className='animate-spin' /> : <AiOutlineLogin />}
                <span className='pl-2'>Sign up</span>
              </div>
            </button>
          </form>
        </div>
        {error && <Alert type={'error'} message={error.toString()} />}
        <div className='mb-2 border-t text-xs pt-2'>
          <p>
            Did you have a accout? go to{' '}
            <span className=''>
              <Link to={'/signin'} className='text-blue-500 hover:border border-blue-500'>
                sigh in
              </Link>
            </span>{' '}
          </p>
        </div>
      </div>
    </>
  );
}
