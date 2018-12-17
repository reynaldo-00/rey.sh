import React, { Component } from 'react';
import { Container, Resume } from './Style';

class Footer extends Component {
    render() {
        return (
            <Container>
                <i className="fab fa-codepen"></i>
                <i className="fab fa-github"></i>
                <Resume/>
                <i className="fab fa-linkedin-in"></i>
                <i className="fab fa-twitter"></i>
            </Container>
        );
    }
}

export default Footer;
