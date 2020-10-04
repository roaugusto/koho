import styled from 'styled-components';

interface RighNavProps {
  open: boolean;
}

export const Container = styled.ul<RighNavProps>`
  list-style: none;
  display: flex;
  flex: 1;
  align-items: center;
  margin-left: 50px;
  justify-content: flex-end;

  button {
    margin-left: auto;
    background: transparent;
    border: 0;
  }

  svg {
    color: #999591;
    width: 26px;
    height: 26px;
  }

  @media (max-width: 768px) {
    z-index: 9998;
    flex-flow: column nowrap;
    background-color: #395ba5;
    opacity: 0.95;
    position: fixed;
    top: 0;
    right: 0;
    height: 100vh;
    width: 300px;
    padding-top: 30px;
    padding-left: 20px;
    align-items: flex-start;
    justify-content: flex-start;

    transform: ${props => (props.open ? 'translateX(0)' : 'translateX(100%)')};
    transition: transform 0.3s ease-in-out;

    transform-origin: 0px;

    li {
      padding-bottom: 10px;
    }

    svg {
      color: #f4ede8;
    }
  }
`;

export const ListItems = styled.div`
  display: flex;
  margin-bottom: 10px;

  svg {
    margin-right: 5px;
    color: #373e4d;
  }

  div {
    display: flex;
    align-items: center;
    margin-right: 30px;
  }

  a {
    text-decoration: none;
    color: #373e4d;

    &:hover {
      opacity: 0.8;
    }
  }

  @media (max-width: 768px) {
    flex-direction: column;
    margin-bottom: 30px;

    div {
      margin-bottom: 10px;
    }

    @media (max-width: 768px) {
      a {
        color: #f4ede8;
      }
      svg {
        margin-right: 20px;
        color: #f4ede8;
      }
    }
  }
`;
