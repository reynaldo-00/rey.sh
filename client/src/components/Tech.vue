<template>
    <div>
        <p>
            {{ techBio }}
        </p>
    </div>
</template>

<script>
import StyledSVG from './SVG.js';

export default {
    name: 'Tech',
    components: {
        StyledSVG
    },
    data: () => ({
        techBio: 'Typically build apps using React for the frontend, Redux for state management, styled components for styling and axios to make calls to either a REST or GraphQL API. The backend usually consist of a REST API written in Node js using an Express server, Knex js as the ORM, SQLite3 for development, Postgres database for production, and an authentication system utilizing JWT and OAuth2.0. Hosting most of my projects on Heroku and ZEIT'
    })
}
</script>

<style>

</style>
