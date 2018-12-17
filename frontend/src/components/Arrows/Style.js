import styled from 'styled-components';

export const Container = styled.div`
    position: absolute;
    top: 50%;
    height: 60px;
    margin-top: -30px;
    z-index: 5;
    width: 100%;
    .fa-chevron-right {
        position: absolute;
        right: 80px;
        top: 50%;
        margin-top: -25px;
        font-size: 50px;
        cursor: pointer;
        color: white;
    }
    .fa-chevron-left {
        position: absolute;
        left: 80px;
        top: 50%;
        margin-top: -25px;
        font-size: 50px;
        cursor: pointer;
        color: white;
    }
`;