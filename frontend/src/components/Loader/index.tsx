import React from 'react';

import CircularProgress from '@material-ui/core/CircularProgress';
import { Container, Popup } from './styles';

interface LoaderProps {
  showPopup: boolean;
}

const Loader: React.FC<LoaderProps> = ({ showPopup }: LoaderProps) => {
  return (
    <>
      {showPopup && (
        <Container>
          <Popup>
            <CircularProgress />
          </Popup>
        </Container>
      )}
    </>
  );
};

export default Loader;
