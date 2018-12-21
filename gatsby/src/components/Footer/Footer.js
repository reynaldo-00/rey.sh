import React from 'react';
import { 
    Container, 
    // Resume,
} from './Style';

import { SVG } from '../Home/Style';

const Footer = () => {
    return (
        <Container>
            <a 
                href="https://codepen.io/reynld/" 
                rel="noopener noreferrer"
                target="_blank"
            >
                <SVG codepen/>
            </a>

            <a 
                href="https://github.com/reynld" 
                rel="noopener noreferrer"
                target="_blank"
            >
                <SVG github/>
            </a>
            {/* <Resume/> */}
            <a 
                href="https://www.linkedin.com/in/reynld/" 
                rel="noopener noreferrer"
                target="_blank"
            >
                <SVG linkedin/>
            </a>
            <a 
                href="https://twitter.com/creynl" 
                rel="noopener noreferrer"
                target="_blank"
            >
                <SVG twitter/>
            </a>
        </Container>
    );
}

export default Footer;
