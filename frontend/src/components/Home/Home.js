import React, { Component } from 'react';
import { HomeWrapper, TextBox, AboutMe } from './Style';

class Home extends Component {
    render() {
        // const bio = `Full Stack Software Engineer based in NYC. Most of my projects consist of writing server side code in Node.js while using SQLite3 or PostgreSQL to store data. Usually use React for the client with Redux to manage the application state, Axios or Apollo to get data from REST or GraphQL APIs, with either styled components or Sass for styling. Hosting most of my projects on Heroku and ZEIT, or using Google Cloud Platform. `;
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
                <AboutMe>
                    <span>{'test'}</span>
                    {/* <span>
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud.
                    </span> */}
                </AboutMe>
            </HomeWrapper>
        );
    }
}

export default Home;
