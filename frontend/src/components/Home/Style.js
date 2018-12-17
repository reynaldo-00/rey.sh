import styled from 'styled-components';

export const HomeWrapper = styled.div`
    position: absolute;
    z-index: 3;
    width: 100%;
    max-width: 1250px;
    /* height: 100%; */
    top: 0;
    left: 0;
    color: white;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
`;

export const TextBox = styled.div`
    padding-left: 80px;
    position: relative;
    margin-top: 120px;
    text-shadow: 1px 1px rgba(0,0,0,0.3);
    display: flex;
    justify-content: space-around;
    align-items: flex-start;
    flex-direction: column;
    h1 {
        height: 55px;
        font-size: 75px;
        font-weight: 800;
        letter-spacing: 1.2px;
        line-height: 55px;
        margin: 20px 0;
    }
    section {
        display: block;
        /* align-items: center;
        justify-content: center; */
        position: relative;
        height: 45px;
        h2 {
            display: inline-block;
            line-height: 45px;
            font-size: 28px;
            margin-left: 10px;
        }
        i {
            display: inline-block;
            font-size: 28px;
        }
    }
`;

export const AboutMe = styled.div`
    width: 600px;
    background-color: rgba(0,0,0,0.6);
    border-radius: 12px;
    padding: 20px;
    color: white;
    margin: 0 auto;
    margin-top: 80px;
`;