import React from 'react';
import { HomeWrapper, TextBox} from './Style';
import Footer from '../Footer/Footer';
import Skills from './Skills';
const Home = () => {
    return (
        <HomeWrapper>
            <TextBox>
                <h1>CHRIS REYNALDO</h1>
                <section>
                    <i className="fas fa-code"></i>
                    <h2>Software Engineer</h2>
                </section>
                <section>
                    <i className="far fa-map"></i>
                    <h2>Queens, New York</h2>
                </section>
            </TextBox>
            <Skills/>
            <Footer/>
        </HomeWrapper>
    );
}

export default Home;
