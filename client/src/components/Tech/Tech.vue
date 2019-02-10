<template>
    <div>
        <h2>BUILD THINGS WITH</h2>
        <div class="tech-used">
            <span
                v-for="(item, index) in tech"
                :key="index"
                class="tech-item"
            >
                <StyledSVG :logo="item"/>
                <span>{{ firstUpper(item) }}</span>
            </span>
        </div>
        <h3>LEARNING</h3>
        <div class="tech-used">
            <span
                v-for="(item, index) in learning"
                :key="index"
                class="tech-item"
            >
                <StyledSVG :logo="item"/>
                <span>{{ firstUpper(item) }}</span>
            </span>
        </div>
        <!-- <p>
            {{ techBio }}
        </p> -->
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
        tech: ['react', 'redux', 'styled', 'node', 'knex', 'jwt'],
        learning: ['vue', 'graphql', 'golang', 'oauth'],
        techBio: 'Typically build apps using React for the frontend, Redux for state management, styled components for styling and axios to make calls to either a REST or GraphQL API. The backend usually consist of a REST API written in Node js using an Express server, Knex js as the ORM, SQLite3 for development, Postgres database for production, and an authentication system utilizing JWT and OAuth2.0. Hosting most of my projects on Heroku and ZEIT'
    }),
    methods: {
        firstUpper(item) {
            const splitString = item.split("");
            splitString[0] = splitString[0].toUpperCase();
            return splitString.join("");
        }
    }
}
</script>

<style scoped>

h2 {
    font-family: 'Montserrat';
    font-weight: 800;
    font-size: 45px;
}

.tech-used {
    display: flex;
    justify-content: center;
    align-items: center;
}

.tech-item {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin: 0 20px;
    font-size: 16px;
    font-weight: 700;
}

</style>
