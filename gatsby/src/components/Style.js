import styled, { createGlobalStyle } from 'styled-components';
import bg from './assets/splash2.jpg'

export const GlobalStyle = createGlobalStyle`
  body {
    @import url('https://fonts.googleapis.com/css?family=Open+Sans:400,600,700,800');
    font-family: 'Open Sans', sans-serif;
  }
`

export const Background = styled.div`
  width: 100vw;
  height: 100vh;
  width: 100vw;
  height: 100vh;
  background: url(${bg});
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  position: fixed;
  left: 0; 
  top: 0;
  z-index: -2;
  filter: grayscale(30%);
`;

export const Container = styled.div`
  width: 100vw;
  min-width: 320px;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  position: relative;
`;

export const ShadeBackground = styled.div`
  width: 100vw;
  height: 100vh;
  background-color: rgba(0,0,0,0.45);
  position: fixed;
  left: 0;
  top: 0;
  z-index: -1;
`;