import React from 'react';
import { AboutContainer } from './Style';

const About = () => {
    const techBio = 'Most of my projects consist of server side written in Node while using Knex as an ORM with SQLite3 during development and PostgreSQL in production to store data. Usually use React for the client with Redux to manage the application state, Axios or Apollo to get data from a REST or GraphQL API and styled components for styling. Hosting most of my projects on Heroku and ZEIT, or Google Cloud Platform.';
    
    return (
        <AboutContainer>
            {techBio}
        </AboutContainer>
    );
}

export default About;
