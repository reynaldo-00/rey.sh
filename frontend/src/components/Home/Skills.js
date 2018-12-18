import React from 'react';
import { AboutMe, SVG, Skill, SkillWrapper, Title } from './Style';

const Skills = () => {
    // const bio = `Full Stack Software Engineer based in NYC. Most of my projects consist of writing server side code in Node.js while using SQLite3 or PostgreSQL to store data. Usually use React for the client with Redux to manage the application state, Axios or Apollo to get data from REST or GraphQL APIs, with either styled components or Sass for styling. Hosting most of my projects on Heroku and ZEIT, or using Google Cloud Platform. `;
    return (
        <AboutMe>
            <Title>build things using</Title>
            <SkillWrapper>
                <Skill>
                    <SVG apollo/>
                    <span>Apollo</span>
                </Skill>
                <Skill>
                    <SVG styled/>
                    <span>Styled C</span>
                </Skill>
                <Skill>
                    <SVG redux/>
                    <span>Redux</span>
                </Skill>
                <Skill>
                    <SVG react/>
                    <span>React</span>
                </Skill>
                <Skill>
                    <SVG node/>
                    <span>Node</span>
                </Skill>
                <Skill>
                    <SVG knex/>
                    <span>Knex</span>
                </Skill>
                <Skill>
                    <SVG graphql/>
                    <span>GraphQL</span>
                </Skill>
                <Skill>
                    <SVG jwt/>
                    <span>JWT</span>
                </Skill>
                <Skill>
                    <SVG oauth/>
                    <span>OAuth</span>
                </Skill>
            </SkillWrapper>
        </AboutMe>
    );
}

export default Skills;
