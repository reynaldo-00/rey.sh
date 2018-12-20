import React from 'react';
import styled from 'styled-components';
import bg from './assets/splash2.jpg'
// import Header from './Header/Header';
import Home from './Home/Home';
// import About from './About/About'
// import Projects from './Projects/Projects';
import Footer from './Footer/Footer';

const App = () => {
  return (
    <Container>
      <Background />
      <ShadeBackground/>
      <Footer />
      {/* <Header/> */}
      <Home />
      {/* <About/>
      <Projects /> */}
    </Container>
  );
}

export default App;

const Background = styled.div`
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