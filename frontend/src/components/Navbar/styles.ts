import styled from 'styled-components';

export const Container = styled.nav`
  display: flex;
  /* flex-wrap: wrap; */
  /* width: 1200px; */
  /* margin: auto; */
  justify-content: space-around;
  background: #ede3dd;

  > div {
    padding: 23px 0;
    display: flex;
    width: 1200px;
  }
`;

export const Logo = styled.div`
  display: flex;
  align-items: center;
  margin-left: 50px;

  img {
    height: 30px;
    margin-left: 20px;
  }

  @media (max-width: 768px) {
    margin-left: 0;
  }
`;
