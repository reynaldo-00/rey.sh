<template>
    <div class="container">
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

export default {
    name: 'App',
    components: {
        AboveTheFold,
        TechStack,
        Projects,
        Contact
    }
}
</script>

<style lang="scss">
    * {
        box-sizing: border-box;
    }
    body {
        width: 100vw;
        height: 100vh;
        margin: 0;
        padding: 0;
        background-color: black;
        color: white;
        font-family: 'Open Sans';
    }
    .container {
        width: 100%;
        max-width: 1000px;
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
    }
</style>
