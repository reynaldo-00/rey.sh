import React from 'react';
import { 
    Container, 
    // Resume 
} from './Style';

const Footer = () => {
    return (
        <Container>
            <a 
                href="https://codepen.io/reynld/" 
                rel="noopener noreferrer"
                target="_blank"
            >
                <i className="fab fa-codepen"></i>
            </a>

            <a 
                href="https://github.com/reynld" 
                rel="noopener noreferrer"
                target="_blank"
            >
                <i className="fab fa-github"></i>
            </a>
            {/* <Resume/> */}
            <a 
                href="https://www.linkedin.com/in/reynld/" 
                rel="noopener noreferrer"
                target="_blank"
            >
                <i className="fab fa-linkedin-in"></i>
            </a>
            <a 
                href="https://twitter.com/creynl" 
                rel="noopener noreferrer"
                target="_blank"
            >
                <i className="fab fa-twitter"></i>
            </a>
        </Container>
    );
}

export default Footer;
