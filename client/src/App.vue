<template>
    <div class="container">
        <Spotify/>
        <div class="section first">
            <Header/>
            <NameBio/>
            <BuildWith/>
        </div>

        <!-- <div class="section">
            <Tech/>
        </div>
        <div class="section">
            <Projects/>
        </div> -->
    </div>
</template>

<script>
import Header from './components/Header'
import NameBio from './components/NameBio'
import Tech from './components/Tech'
import Projects from './components/Projects'
import Spotify from './components/Spotify'
import BuildWith from './components/BuildWith'

export default {
    name: 'App',
    components: {
        NameBio,
        Tech,
        Projects,
        Spotify,
        Header,
        BuildWith,
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
        min-width: 435px;
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
        min-height: 600px;
        position: relative;
    }

    .first {
        text-align: left;
    }
</style>
