import React from 'react'
import Home from '../components/Home/Home';
import Footer from '../components/Footer/Footer';

import {
  Container,
  Background,
  ShadeBackground,
  GlobalStyle
} from '../components/Style';

const IndexPage = () => (
  <Container>
    <GlobalStyle/>
    <Background />
    <ShadeBackground/>
    <Footer />
    <Home />
  </Container>
)

export default IndexPage
