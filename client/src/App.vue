<template>
    <div class="container">

        <div class="section first">
            <NameBio/>
        </div>

        <div class="section">
            <Tech/>
        </div>

        <div class="section">
            <Projects/>
        </div>
        <Spotify/>
    </div>
</template>

<script>
import NameBio from './components/NameBio'
import Tech from './components/Tech'
import Projects from './components/Projects'
import Spotify from './components/Spotify'

export default {
    name: 'App',
    components: {
        NameBio,
        Tech,
        Projects,
        Spotify
    }
}
</script>

<style>
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
        max-width: 1050px;
        min-width: 835px;
        padding: 0px 50px;
        margin: 0 auto;
    }
    .section {
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        text-align: center;
        height: 100vh;
    }

    .first {
        text-align: left;
    }
</style>
