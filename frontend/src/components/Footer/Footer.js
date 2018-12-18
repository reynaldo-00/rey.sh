import React, { Component } from 'react';
import { Container, Resume } from './Style';

class Footer extends Component {
    render() {
        return (
            <Container>
                <a 
                    href="https://codepen.io/reynaldo-00/" 
                    rel="noopener noreferrer"
                    target="_blank"
                >
                    <i className="fab fa-codepen"></i>
                </a>

                <a 
                    href="https://github.com/reynaldo-00" 
                    rel="noopener noreferrer"
                    target="_blank"
                >
                    <i className="fab fa-github"></i>
                </a>
                <Resume/>
                <a 
                    href="https://www.linkedin.com/in/reynaldo-00/" 
                    rel="noopener noreferrer"
                    target="_blank"
                >
                    <i className="fab fa-linkedin-in"></i>
                </a>
                <a 
                    href="https://twitter.com/chris_rey_" 
                    rel="noopener noreferrer"
                    target="_blank"
                >
                    <i className="fab fa-twitter"></i>
                </a>
            </Container>
        );
    }
}

export default Footer;
