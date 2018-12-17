import React, { Component } from 'react';
import { Container, Navigation } from './Style';
import { NavLink } from 'react-router-dom';

class Header extends Component {
    render() {
        return (
            <Container>
                <Navigation>
                    <NavLink to="/">ABOUT</NavLink >
                    <NavLink to="/projects">PROJECTS</NavLink >
                    <NavLink to="/contact">CONTACT</NavLink >
                </Navigation>
            </Container>
        );
    }
}

export default Header;
