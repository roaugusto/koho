import React from 'react';
import { Link } from 'react-router-dom';
import { FiHome, FiLayers } from 'react-icons/fi';

import { Container, ListItems } from './styles';

interface RightNavProps {
  open: boolean;
}

const RightNav: React.FC<RightNavProps> = ({ open }) => {
  return (
    <Container open={open}>
      <ListItems>
        <div>
          <FiHome />
          <Link to="/">Dashboard</Link>
        </div>
        <div>
          <FiLayers />
          <Link to="/import">Import File</Link>
        </div>
      </ListItems>
    </Container>
  );
};

export default RightNav;
