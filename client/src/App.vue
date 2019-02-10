<template>
    <div class="container">
        <Spotify/>
        <AboveTheFold/>
    </div>
</template>

<script>
import Spotify from './components/Spotify/index'
import AboveTheFold from './components/AboveTheFold/index';

export default {
    name: 'App',
    components: {
        Spotify,
        AboveTheFold
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
        padding: 24px;
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
    }
    .first {
        height: 100vh;
    }
</style>
