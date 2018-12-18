import React from 'react';
import { Container, Navigation } from './Style';
import { NavLink } from 'react-router-dom';

const Header = () => {
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

export default Header;
