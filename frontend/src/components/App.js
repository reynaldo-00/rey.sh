import React, { Component } from 'react';
import styled from 'styled-components';
import bg from './assets/splash2.jpg'
import Home from './Home/Home';
import Header from './Header/Header';
import Footer from './Footer/Footer';
import { Route } from 'react-router-dom';
import Projects from './Projects/Projects';
import Arrows from './Arrows/Arrows';

class App extends Component {
  render() {
    return (
      <Container>
        <ShadeBackground/>
        <Header/>
        <Route path="/" render={props => <Arrows {...props}/>}/>
        <Route path="/" exact render={props => <Home {...props}/>}/>
        <Route path="/projects" exact render={props => <Projects {...props}/>}/>
        <Footer/>
      </Container>
    );
  }
}

export default App;

export const Container = styled.div`
  width: 100vw;
  height: 100vh;
  background: url(${bg});
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  position: relative;
  filter: grayscale(30%);
  display: flex;
  justify-content: center;
  align-items: center;
`;

export const ShadeBackground = styled.div`
  width: 100vw;
  height: 100vh;
  background-color: rgba(0,0,0,0.45);
`;