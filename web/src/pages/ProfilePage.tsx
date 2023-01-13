import React from 'react';
import Menu from '../components/Menu';
import { Modal } from '../components/Modal';
import { UpdateAvatarForm } from '../form/Avatar';
import { useAuth } from '../hooks';
import { Avatar, Box, Paper } from '@mui/material';
import Accordion from '@mui/material/Accordion';
import AccordionSummary from '@mui/material/AccordionSummary';
import AccordionDetails from '@mui/material/AccordionDetails';
import Typography from '@mui/material/Typography';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import Container from '@mui/material/Container';
import { ChangePasswordForm } from '../form/ChangePasswordForm';

type Props = {};

const ProfilePage = (props: Props) => {
  const { user } = useAuth();

  const [openMenu, setOpenMenu] = React.useState(false);
  const [openModal, setOpenModal] = React.useState(false);

  return (
    <React.Fragment>
      <Container maxWidth='xs'>
        <Paper elevation={0}>
          <Box paddingX={'1.5rem'} paddingY={'1.25rem'} marginY={'1rem'}>
            <div className='py-5'>
              <div className='w-20 h-20 relative rounded-full overflow-hidden'>
                <div
                  onClick={() => setOpenMenu(!openMenu)}
                  className='absolute cursor-pointer top-0 left-0 w-full h-full z-20 hover:opacity-20 opacity-0 bg-gray-300'
                ></div>
                <Avatar
                  sx={{
                    width: '100%',
                    height: '100%',
                  }}
                  src={process.env.REACT_APP_SERVER_URL + '' + user?.avatar.url}
                  alt={'' + user?.avatar.title}
                />
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
          </Box>
        </Paper>
      </Container>
      {/* setting zone */}
      <Container maxWidth='xs'>
        <Paper elevation={0}>
          <Box paddingX={'.4rem'} paddingY={'.4rem'} marginY={'.4rem'}>
            <Accordion elevation={0}>
              <AccordionSummary expandIcon={<ExpandMoreIcon />} aria-controls='panel1a-content' id='panel1a-header'>
                <Typography>เปลี่ยนรหัสผ่าน</Typography>
              </AccordionSummary>
              <AccordionDetails>
                {/* form */}
                {user && <ChangePasswordForm uId={user.id} />}
              </AccordionDetails>
            </Accordion>
          </Box>
        </Paper>
      </Container>
      <Modal open={openModal} setOpen={setOpenModal}>
        <UpdateAvatarForm setOpen={setOpenModal} />
      </Modal>
    </React.Fragment>
  );
};

export default ProfilePage;
