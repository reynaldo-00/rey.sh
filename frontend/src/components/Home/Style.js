import styled from 'styled-components';

export const HomeWrapper = styled.div`
    position: absolute;
    z-index: 3;
    width: 100%;
    max-width: 1250px;
    height: 100%;
    color: white;
    display: flex;
    justify-content: flex-start;
    align-items: center;
`;

export const TextBox = styled.div`
    padding-left: 80px;
    margin-top: -20%;
    text-shadow: 1px 1px rgba(0,0,0,0.3);
    display: flex;
    justify-content: space-around;
    align-items: flex-start;
    flex-direction: column;
    h1 {
        font-size: 75px;
        font-weight: 800;
        letter-spacing: 1.2px;
        line-height: 55px;
        margin: 20px 0;
    }
    section {
        display: flex;
        align-items: center;
        justify-content: center;
        position: relative;
        height: 45px;
        h2 {
            font-size: 28px;
            margin-left: 10px;
        }
        i {
            font-size: 28px;
        }
    }
`;

export const AboutMe = styled.div`
    width: 600px;
    /* height: 170px; */
    background-color: rgba(0,0,0,0.6);
    border-radius: 12px;
    padding: 20px;
    position: absolute;
    bottom: 100px;
    left: 50%;
    margin-left: -300px;
    color: white;
`;