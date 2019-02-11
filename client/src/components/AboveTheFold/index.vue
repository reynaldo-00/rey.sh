<template>
    <section class="section first">
        <div class="background-img"/>
        <Spotify/>
        <Header/>
        <Description/>
        <Scroll/>
    </section>
</template>

<script>
import Header from './Header';
import Spotify from './Spotify/index'
import Description from './Description';
import Scroll from './Scroll';

export default {
    name: 'AboveTheFold',
    components: {
        Header,
        Description,
        Scroll
    }
}
</script>

<style lang="scss">
    .first {
        height: 100vh;
        min-height: 0px;
        padding: 24px;
        position:relative;
        .background-img {
            position: fixed;
            width: 100vw;
            height: 100vh;
            left: 0;
            top: 0;
            background: linear-gradient(rgba(0, 0, 0, 0.6),rgba(0, 0, 0, 0.6)), url('../../assets/splash2.jpg');
            background-position: center center;
            background-repeat: no-repeat;
            background-size: cover;
        }
    }
</style>
