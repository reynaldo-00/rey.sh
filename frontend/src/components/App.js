import React from 'react';
import styled from 'styled-components';
import bg from './assets/splash2.jpg'
import Home from './Home/Home';
import Header from './Header/Header';
import { Route } from 'react-router-dom';
import Projects from './Projects/Projects';
// import Arrows from './Arrows/Arrows';

const App = () => {
  return (
    <Container>
      <Background />
      <ShadeBackground/>
      <Header/>
      {/* <Route path="/" render={props => <Arrows {...props}/>}/> */}
      <Route path="/" exact render={props => <Home {...props}/>}/>
      <Route path="/projects" exact render={props => <Projects {...props}/>}/>
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
  z-index: 1;
  filter: grayscale(30%);
`;

export const Container = styled.div`
  width: 100vw;
  min-height: 100vh;
  display: flex;
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
  z-index: 2;
`;