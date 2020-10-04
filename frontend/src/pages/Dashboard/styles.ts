import styled from 'styled-components';
import { shade } from 'polished';

export const Container = styled.div`
  width: 100%;
  max-width: 1120px;
  margin: 0 auto;
  padding: 20px 20px;

  .MuiTableCell-root {
    padding: 10px;
  }
`;

export const ButtonDownload = styled.button`
  background: #f39d87;
  color: #fff;
  border-radius: 5px;
  padding: 3px 10px;
  border: 0;
  transition: background-color 0.2s;

  &:hover {
    background: ${shade(0.2, '#f39d87')};
  }
`;

export const Title = styled.h1`
  font-size: 48px;
  color: #3a3a3a;
`;
