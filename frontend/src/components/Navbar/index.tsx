import React from 'react';

import logoImg from '../../assets/koho-logo2.png';

import Burger from './Burger';

import { Container, Logo } from './styles';

const Navbar: React.FC = () => {
  return (
    <Container>
      <div>
        <Logo>
          <img src={logoImg} alt="Gft" />
        </Logo>
        <Burger />
      </div>
    </Container>
  );
};

export default Navbar;
