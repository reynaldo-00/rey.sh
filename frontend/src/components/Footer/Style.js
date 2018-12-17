import styled from 'styled-components';

export const Container = styled.footer`
    width: 100vw;
    height: 84px;
    position: absolute;
    bottom: 0;
    left: 0;
    z-index: 9999;
    display: flex;
    justify-content: center;
    align-items: center;
    color: white;
    background: transparent;
    i {
        font-size: 32px;
        margin: 0 30px;
        cursor: pointer;
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