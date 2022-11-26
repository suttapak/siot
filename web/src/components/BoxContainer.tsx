import React from 'react';
import { MenuBar } from './MenuBar';

type Props = {
  children: React.ReactNode;
};

const BoxContainer = (props: Props) => {
  const { children } = props;
  return (
    <React.Fragment>
      <MenuBar />
      <div>{children}</div>
    </React.Fragment>
  );
};

export default BoxContainer;
