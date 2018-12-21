import styled from 'styled-components';
import { SVG } from '../Home/Style';

export const Container = styled.footer`
    margin-top: 20px;
    width: 100vw;
    height: 84px;
    z-index: 5;
    display: flex;
    justify-content: center;
    align-items: center;
    color: white;
    background: transparent;
    a {
        text-decoration: none;
        opacity: 0.74;
        :hover {
            opacity: 1;
        }
    }
    ${SVG} {
        margin: 0 30px;
        cursor: pointer;
        @media (max-width: 440px) {
            margin: 0 20px;
        }
    }
`;

export const Resume = styled.span`
    ::after {
        content: 'RESUME'
    }
    width: 120px;
    height: 36px;
    background: rgba(0,0,0,0.8);
    color: rgba(255,255,255,0.8);
    font-size: 22px;
    font-weight: 800;
    border-radius: 25px;
    text-align: center;
    line-height: 34px;
    cursor: pointer;
`;