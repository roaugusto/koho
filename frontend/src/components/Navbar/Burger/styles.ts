import styled from 'styled-components';

interface ButtonProps {
  open: boolean;
}

export const Container = styled.div<ButtonProps>`
  z-index: 9999;
  width: 2rem;
  height: 2rem;
  position: fixed;
  top: 15px;
  right: 20px;
  display: none;

  div {
    width: 2rem;
    height: 0.25rem;
    background-color: ${props => (props.open ? '#ccc' : '#395ba5')};
    border-radius: 10px;
    transform-origin: 1px;
    transition: all 0.3s;

    &:nth-child(1) {
      /* width: ${props => (props.open ? '1.5rem' : '2rem')}; */
      transform: ${props => (props.open ? 'rotate(45deg)' : 'rotate(0)')};
    }

    &:nth-child(2) {
      transform: ${props =>
        props.open ? 'translateX(250%)' : 'translateX(0)'};
    }

    &:nth-child(3) {
      /* width: ${props => (props.open ? '1.5rem' : '2rem')}; */
      transform: ${props => (props.open ? 'rotate(-45deg)' : 'rotate(0)')};
    }
  }

  @media (max-width: 768px) {
    display: flex;
    justify-content: space-around;
    flex-flow: column nowrap;
  }
`;
