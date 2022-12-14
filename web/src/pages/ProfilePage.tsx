import React from 'react';
import Menu from '../components/Menu';
import { Modal } from '../components/Modal';
import { UpdateAvatarForm } from '../form/Avatar';
import { useAuth } from '../hooks';

type Props = {};

const ProfilePage = (props: Props) => {
  const { user } = useAuth();

  const [openMenu, setOpenMenu] = React.useState(false);
  const [openModal, setOpenModal] = React.useState(false);

  return (
    <React.Fragment>
      <div className='px-6 mx-auto max-w-md bg-white rounded-lg'>
        <div className='py-5'>
          <div className='w-20 h-20 relative rounded-full overflow-hidden'>
            <div
              onClick={() => setOpenMenu(!openMenu)}
              className='absolute cursor-pointer top-0 left-0 w-full h-full z-20 hover:opacity-20 opacity-0 bg-gray-300'
            ></div>
            <img className='relative z-10' src={process.env.REACT_APP_SERVER_URL + '' + user?.avatar.url.substring(1)} alt={'' + user?.avatar.title} />
          </div>
          <Menu position='center' open={openMenu} setOpen={setOpenMenu}>
            <li className='px-3 w-full py-2 cursor-pointer  hover:bg-gray-200'>
              <span
                onClick={() => {
                  setOpenModal(!openModal);
                  setOpenMenu(!openMenu);
                }}
              >
                Change profile
              </span>
            </li>
          </Menu>
          <h1 className='text-lg tracking-widest'>
            {user?.firstName} {user?.lastName}
          </h1>
          <p className='text-gray-700'>{user?.email}</p>
        </div>
      </div>
      <Modal open={openModal} setOpen={setOpenModal}>
        <UpdateAvatarForm setOpen={setOpenModal} />
      </Modal>
    </React.Fragment>
  );
};

export default ProfilePage;
