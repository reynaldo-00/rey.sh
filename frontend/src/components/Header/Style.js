import styled from 'styled-components';

export const Container = styled.header`
    width: 100vw;
    /* max-width: 900px; */
    background-color: transparent;
    height: 84px;
    /* position: fixed; */
    position: absolute;
    top: 0;
    left: 0;
    z-index: 99999;
`;

export const Navigation = styled.nav`
    width: 400px;
    position: absolute;
    right: 80px;
    height: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-direction: row;
    a {
        text-decoration: none;
        font-size: 16px;
        font-weight: 800;
        color: white;
        user-select: none;
    }
`;
