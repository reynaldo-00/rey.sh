<template>
    <div class="container">
        <Name/>
        <Bio/>
        <Tech/>
    </div>
</template>

<script>
import Name from './components/Name'
import Bio from './components/Bio'
import Tech from './components/Tech'

export default {
    name: 'App',
    components: {
        Name,
        Bio,
        Tech
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
        padding: 0px 50px;
    }
</style>
