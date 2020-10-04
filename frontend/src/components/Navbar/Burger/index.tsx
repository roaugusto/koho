import React, { useState } from 'react';
import RightNav from '../RightNav';

import { Container } from './styles';

const Burger: React.FC = () => {
  const [open, setOpen] = useState(false);

  const handleClick = (): void => {
    setOpen(!open);
  };

  return (
    <>
      <Container open={open} onClick={handleClick}>
        <div />
        <div />
        <div />
      </Container>
      <RightNav open={open} />
    </>
  );
};

export default Burger;
