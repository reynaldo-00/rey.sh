<template>
    <div class="container">
        <Spotify/>
        <AboveTheFold/>
        <TechStack/>
        <Projects/>
        <Contact/>
    </div>
</template>

<script>
import AboveTheFold from './components/AboveTheFold/index';
import TechStack from './components/TechStack/index';
import Projects from './components/Projects/index';
import Contact from './components/Contact/index';
import Spotify from './components/Spotify';
import { page } from 'vue-analytics'

export default {
    name: 'App',
    components: {
        AboveTheFold,
        TechStack,
        Projects,
        Contact,
        Spotify,
    },
    methods: {
        track () {
            page('/')
        }
    },
    beforeMount: function() {
        this.track();
    }
}
</script>

<style lang="scss">
    * {
        box-sizing: border-box;
    }
    body {
        width: 100%;
        height: 100vh;
        margin: 0;
        padding: 0;
        background-color: black;
        color: white;
        font-family: 'Open Sans';
        overflow-x: hidden; 
    }
    .container {
        width: 100%;
        max-width: 1000px;
        min-width: 330px;
        margin: 0 auto;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: center;
    }
    .section {
        width: 100%;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: flex-start;
        position: relative;
        min-height: 650px;
        height: 100vh;
        max-height: 1080px;
        @media (max-width: 620px) {
            height: auto;
            min-width: 330px;
        }
    }
</style>
