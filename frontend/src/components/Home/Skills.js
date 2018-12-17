import React from 'react';
import { AboutMe } from './Style';

const Skills = () => {
    const bio = `Full Stack Software Engineer based in NYC. Most of my projects consist of writing server side code in Node.js while using SQLite3 or PostgreSQL to store data. Usually use React for the client with Redux to manage the application state, Axios or Apollo to get data from REST or GraphQL APIs, with either styled components or Sass for styling. Hosting most of my projects on Heroku and ZEIT, or using Google Cloud Platform. `;
    return (
        <AboutMe>
            <span>{bio}</span>
        </AboutMe>
    );
}

export default Skills;
